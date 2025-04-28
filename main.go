package main

import (
	"fmt"
	"go-api/routes"
	"go-api/utils"
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

// @title           API Shadow
// @version         1.0
// @description     A Shadow API é uma interface robusta e segura, projetada para fornecer dados estratégicos e operacionais a usuários-chave da organização.
// @description		Esta API tem como principal objetivo dar suporte a processos de tomada de decisão, análise de desempenho e integração de dados entre sistemas.
// @description		Usuários-alvo incluem times de operações, lideranças estratégicas, analistas e demais áreas que demandam informações atualizadas e estruturadas.
// @description
// @description		Funcionalidades incluem:
//@description		- Consulta de dados operacionais em tempo real
// @description		- Acesso a indicadores estratégicos
// @description		- Integração com plataformas internas e externas
// @description		- Suporte a autenticação básica (Basic Auth)
// @description
// @description		Para utilizar a API, é necessário autenticação e autorização apropriada.
// @description		Em caso de dúvidas, entre em contato com o time responsável pela governança de dados.
// @securityDefinitions.basic  BasicAuth
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	//database.Connect()
	utils.LoadMessages()

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

	router.RunTLS(":8080", "wildcard_aviva_com_br.crt", "wildcard_aviva_com_br.key")
}

func HashSenha(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(hash), err
}
