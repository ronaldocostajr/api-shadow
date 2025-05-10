//Desenvolvedor: Thiago Leite
//Versão: 1.0.0 V
//Compilação: 2025-04-30 14:49:26.7086248 -0300 -03 m=+6.642361001
package routes

import (
	"go-api/controllers"
	"github.com/gin-gonic/gin"
)

func Tb_api_dia_feriadoRoutes(r *gin.RouterGroup) {
	tb_api_dia_feriado := r.Group("/tb_api_dia_feriado") 
	{
		tb_api_dia_feriado.GET("/", controllers.GetTb_api_dia_feriado)
	}
}

