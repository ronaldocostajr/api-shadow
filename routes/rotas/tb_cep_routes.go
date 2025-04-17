//Desenvolvedor: Ronaldo Costa
//Versão: 1.0.0 V
//Compilação: 2025-04-17 08:35:03.9165257 -0300 -03 m=+5.739813101
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

