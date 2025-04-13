package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func SaudacaoRoutes(r *gin.RouterGroup) {
	saudacao := r.Group("/saudacao/")
	{
		saudacao.GET("/:nome", controllers.Saudacao)
	}
}