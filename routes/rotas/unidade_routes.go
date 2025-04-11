package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func UnidadeRoutes(r *gin.RouterGroup) {
	unidade := r.Group("/unidades")
	{
		unidade.GET("/", controllers.GetUnidades)
		unidade.GET("/:id", controllers.GetUnidadeByID)
		unidade.POST("/", controllers.CreateUnidade)
		unidade.PUT("/:id", controllers.UpdateUnidade)
		unidade.DELETE("/:id", controllers.DeleteUnidade)
	}
}
