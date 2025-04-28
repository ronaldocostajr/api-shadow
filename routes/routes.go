package routes

import (
	"go-api/controllers"
	middleware "go-api/middlewares"
	routes "go-api/routes/rotas"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rota pública
	router.POST("/login", controllers.Login)
	
	// Rotas protegidas com JWT
	api := router.Group("/api")
	//api.Use(middleware.JWTAuth()) // protege todas as rotas abaixo
	routes.SwaggerRoutes(api)
	// Rate Limiter 5 requisições por segundo
	api.Use(middleware.RateLimiter())
	api.Use(middleware.BasicAuth())
	{
		routes.UnidadeRoutes(api)
		routes.UsuarioRoutes(api)
		routes.Tb_clienteRoutes(api)
		routes.WriteGeneratorRoutes(api)
		routes.Tb_paisRoutes(api)
		routes.Tb_cepRoutes(api)
		routes.SaudacaoRoutes(api)
		// NÃO RETIRAR ESSA LINHA
	}
}
