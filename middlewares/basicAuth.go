package middleware

/*
Middleware para autenticação básica com username e password.

Este middleware verifica se o cabeçalho "Authorization" está presente e se contém as credenciais corretas.
A validação será feita pelo protocolo LDAP, mas por enquanto, estamos usando um usuário e senha fixos.

*/

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"go-api/database"
	"go-api/models"
	"go-api/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	log "github.com/sirupsen/logrus"
)

type UserLogin struct {
	Username string
	Email string
	Password string
   }

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		payload, err := base64.StdEncoding.DecodeString(auth[len("Basic "):])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.unauthenticated"),
			})
			c.Abort()
			return
		}

		pair := strings.SplitN(string(payload), ":", 2)
		//fmt.Println("Payload: ",string(payload), "Tamanho: ", len(pair), "Pair: ", pair)
		user := UserLogin{
			Username: "",
			Email: pair[0],
			Password: pair[1],
		}
		if len(pair) != 2 || len(user.Email) <= 0 || len(user.Password) <= 0{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": utils.GetMessage("geral", "error.bad_request_auth"),
			})
			c.Abort()
			return
		}
		conn, err := Connect()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": utils.GetMessage("geral", "error.internal_ldap"),
			})
			c.Abort()
			return
		}
		defer conn.Close()

		user.Username, _ = searchUser(conn, user.Email)
		searchRequest := ldap.NewSearchRequest(
		 os.Getenv("LDAP_BASE_DN"),
		 ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		 fmt.Sprintf("(sAMAccountName=%s)", user.Username),
		 []string{"dn"},
		 nil,
		)
	   
		searchResp, err := conn.Search(searchRequest)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.credentials_invalid"),
			})
			c.Abort()
			return
		}
	   
		if len(searchResp.Entries) != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.credentials_invalid"),
			})
			c.Abort()
			return
		}
	   
		userDN := searchResp.Entries[0].DN
		
		err = conn.Bind(userDN, user.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.credentials_invalid"),
			})
			c.Abort()
			return
		}

		status, err := GetStatusByEmail(user.Email)

		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Algo deu errado, procure o administrador da API.",
			})
			c.Abort()
			return
		}else if strings.TrimSpace(strings.ToUpper(status[0])) == "N" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.user_inactive"),
			})
			c.Abort()
			return
		}
		c.Set("email", user.Email)
		c.Next()

	}
}

func GetStatusByEmail(email string) ([]string, error) {
	var roles []string
	err := database.DB.Model(&models.Tb_log_usuario{}).
		Where("DS_EMAIL = ?", email).
		Pluck("FL_ATIVO", &roles).Error
	return roles, err
}

func Connect() (*ldap.Conn, error) {
	conn, err := ldap.DialTLS("tcp", os.Getenv("LDAP_ADDRESS"), &tls.Config{InsecureSkipVerify: false})
	if err != nil {
	 log.Errorf("LDAP connection failed, error details: %v", err)
	 return nil, err
	}
   
	if err := conn.Bind(os.Getenv("BIND_USER"), os.Getenv("BIND_PASSWORD")); err != nil {
	 log.Errorf("LDAP bind failed while connecting, error details: %v", err)
	 return nil, err
	}
   
	return conn, nil
}

func Auth(conn *ldap.Conn, user UserLogin) (bool, error) {
	username, err := searchUser(conn, user.Email)
	
	if err != nil{
		log.Errorf("Erro qualquer coisa: %v", err)
		return false, err
	}

	user.Username = username
	
	searchRequest := ldap.NewSearchRequest(
	 os.Getenv("LDAP_BASE_DN"),
	 ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	 fmt.Sprintf("(sAMAccountName=%s)", user.Username),
	 []string{"dn"},
	 nil,
	)
   
	searchResp, err := conn.Search(searchRequest)
	if err != nil {
	 log.Errorf("LDAP search failed for user %s, error details: %v", user.Username, err)
	 return false, err
	}
   
	if len(searchResp.Entries) != 1 {
	 log.Errorf("User: %s not found or multiple entries found", user.Username)
	 err = fmt.Errorf("user: %s not found or multiple entries found", user.Username)
	 return false, err
	}
   
	userDN := searchResp.Entries[0].DN
	fmt.Println("TESTE DO search")
	
	err = conn.Bind(userDN, user.Password)
	if err != nil {
	 log.Errorf("LDAP authentication failed for user %s, error details: %v", user.Username, err)
	 err = fmt.Errorf("LDAP authentication failed for user %s", user.Username)
	 return false, err
	}
   
	return true, nil
}

func searchUser(conn *ldap.Conn, email string) (string, error){
	//_email := "thiago.leite@aviva.com.br"
	searchRequest := ldap.NewSearchRequest(
		os.Getenv("LDAP_BASE_DN"),
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(UserPrincipalName=%s)", email),
		[]string{"dn", "cn", "mail", "sAMAccountName"},
		nil,
	)

	// Executa a busca do SamAccountName pelo email passado
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return "", err
	}

	return sr.Entries[0].GetAttributeValue("sAMAccountName"), nil
}