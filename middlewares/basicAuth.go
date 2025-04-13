package middleware

/*
Middleware para autenticação básica com username e password.

Este middleware verifica se o cabeçalho "Authorization" está presente e se contém as credenciais corretas.
A validação será feita pelo protocolo LDAP, mas por enquanto, estamos usando um usuário e senha fixos.

*/

import (
	"encoding/base64"
	"go-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	username = "admin"
	password = "1234"
)

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
		if len(pair) != 2 || pair[0] != username || pair[1] != password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": utils.GetMessage("geral", "error.credentials_invalid"),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}