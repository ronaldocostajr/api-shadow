package main

import (
	"fmt"
	"go-api/database"
	"go-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// // Define a rate limiter with 1 request per second and a burst of 5.
// var limiter = rate.NewLimiter(1, 5)

// // Middleware to check the rate limit.
// func rateLimiter(c *gin.Context) {
// 	if !limiter.Allow() {
// 		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
// 		c.Abort()
// 		return
// 	}
// 	c.Next()
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	database.Connect()

	router := gin.Default()

	// // Apply the rate limiting middleware
	// router.Use(rateLimiter)

	routes.SetupRoutes(router)

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s | %3d | %13v | %15s | %-7s %#v\nHeaders: %v\n\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Header, // imprime os headers
		)
	}))

	router.RunTLS(":8080", "STAR_aviva_com_br.crt", "aviva.com.br.key")
}

func HashSenha(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(hash), err
}
