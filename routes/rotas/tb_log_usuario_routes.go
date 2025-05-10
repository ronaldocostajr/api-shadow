// Desenvolvedor: Thiago Leite
// Versão: 1.0.0 V
// Compilação: 2025-04-28 17:00:05.7595771 -0300 -03 m=+6.320551501
package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func Tb_log_usuarioRoutes(r *gin.RouterGroup) {
	tb_log_usuario := r.Group("/tb_log_usuario") 
	{
		tb_log_usuario.GET("/", controllers.GetTb_log_usuario)
	}
}

