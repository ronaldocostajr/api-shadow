//Desenvolvedor: Thiago Leite
//Versão: 1.0.0 V
//Compilação: 2025-04-30 13:42:26.6206398 -0300 -03 m=+6464.443405201
package routes

import (
	"go-api/controllers"
	"github.com/gin-gonic/gin"
)

func Tb_cdp_cepRoutes(r *gin.RouterGroup) {
	tb_cdp_cep := r.Group("/tb_cdp_cep") 
	{
		tb_cdp_cep.GET("/", controllers.GetTb_cdp_cep)
	}
}

