package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func UsuarioRoutes(r *gin.RouterGroup) {
	usuario := r.Group("/usuarios")
	{
		usuario.GET("/", controllers.GetUsuarios)
		usuario.GET("/:id", controllers.GetUsuarioByID)
	}
}
