package routes

import (
	"go-api/controllers"
	"github.com/gin-gonic/gin"
)

func Tb_paisRoutes(r *gin.RouterGroup) {
	tb_pais := r.Group("/tb_pais") 
	{
		tb_pais.GET("/", controllers.GetTb_pais)
	}
}

