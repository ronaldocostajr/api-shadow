// Middleware to check the rate limit.
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Define a rate limiter with 1 request per second and a burst of 5.
var limiter = rate.NewLimiter(1, 5)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
