package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ContentTypeJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		headers := ""

		for name, values := range c.Request.Header {
			for _, value := range values {
				headers += fmt.Sprintf(` -H "%s: %s"`, name, value)
			}
		}

		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		c.Next()
	}
}