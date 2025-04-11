//Desenvolvedor: Ronaldo Costa
//Versão: 1.0.0 V
//Compilação: 2025-04-11 08:09:39.9245809 -0300 -03 m=+1123.012202101
//Comentário adicional: código adiconal
package routes

import (
	"go-api/controllers"
	"github.com/gin-gonic/gin"
)

func Tb_cepRoutes(r *gin.RouterGroup) {
	tb_cep := r.Group("/tb_cep") 
	{
		tb_cep.GET("/", controllers.GetTb_cep)
	}
}

