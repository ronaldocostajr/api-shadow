package routes

import (
	"go-api/controllers"
	"github.com/gin-gonic/gin"
)

func Tb_clienteRoutes(r *gin.RouterGroup) {
	tb_cliente := r.Group("/tb_cliente") 
	{
		tb_cliente.GET("/", controllers.GetTb_cliente)
	}
}

