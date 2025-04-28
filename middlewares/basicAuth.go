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
	_ "fmt"
	"go-api/utils"
	"net/http"
	"os"
	_ "os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	_ "github.com/go-ldap/ldap/v3"
	log "github.com/sirupsen/logrus"
)

const (
	username = "admin"
	password = "1234"
)

type UserLogin struct {
	Username string
	Email string
	Password string
   }

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.unauthorized"),
			})
			c.Abort()
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[len("Basic "):])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.unauthenticated"),
			})
			c.Abort()
			return
		}

		pair := strings.SplitN(string(payload), ":", 2)
		//fmt.Println("valor de pair", pair)
		user := UserLogin{
			Username: "",
			Email: pair[0],
			Password: pair[1],
		}
		if len(pair) != 2{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": utils.GetMessage("geral", "error.bad_request"),
			})
			c.Abort()
			return
		}
		conn, err := Connect()
		if err != nil {
		    log.Fatal("LDAP connection failed.")
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
		 //log.Errorf("LDAP search failed for user %s, error details: %v", user.Username, err)
		 //return false, err
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.credentials_invalid"),
			})
			c.Abort()
			return
		}
	   
		if len(searchResp.Entries) != 1 {
		 //log.Errorf("User: %s not found or multiple entries found", user.Username)
		 //err = fmt.Errorf("user: %s not found or multiple entries found", user.Username)
		 //return false, err
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.credentials_invalid"),
			})
			c.Abort()
			return
		}
	   
		userDN := searchResp.Entries[0].DN
		fmt.Println("TESTE DO search")
		
		err = conn.Bind(userDN, user.Password)
		if err != nil {
		 //log.Errorf("LDAP authentication failed for user %s, error details: %v", user.Username, err)
		 //err = fmt.Errorf("LDAP authentication failed for user %s", user.Username)
		 //return false, err
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.credentials_invalid"),
			})
			c.Abort()
			return
		}
	   
		c.Next()

		// authenticated, authErr := Auth(conn, user)
		// if authErr != nil {
		//     log.Fatal("Authentication failed.")
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": utils.GetMessage("geral", "error.credentials_invalid"),
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// if authenticated {
		//     log.Info("User authenticated successfully.")
		// 	c.Next()
		// } else {
		//     log.Info("Authentication failed. Invalid credentials.")
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": utils.GetMessage("geral", "error.credentials_invalid"),
		// 	})
		// 	c.Abort()
		// 	return
		// }
		//c.Next()
	}
}


func Connect() (*ldap.Conn, error) {
	conn, err := ldap.DialTLS("tcp", os.Getenv("LDAP_ADDRESS"), &tls.Config{InsecureSkipVerify: true})
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
	fmt.Println("TESTE DO search dentro")
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

	// Executa a busca
	sr, err := conn.Search(searchRequest)
	if err != nil {
		//log.Fatalf("Erro na busca: %v", err)
		return "", err
	}

	fmt.Println("SR", sr.Entries[0].GetAttributeValue("sAMAccountName"))

	return sr.Entries[0].GetAttributeValue("sAMAccountName"), nil

	// Mostra os resultados
	// for _, entry := range sr.Entries {
	// 	fmt.Printf("DN: %s\n", entry.DN)
	// 	fmt.Printf("CN: %s\n", entry.GetAttributeValue("cn"))
	// 	fmt.Printf("Email: %s\n", entry.GetAttributeValue("mail"))
	// 	fmt.Printf("sAMAccountName: %s\n", entry.GetAttributeValue("sAMAccountName"))
	// }
}