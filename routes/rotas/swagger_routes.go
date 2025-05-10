package routes

import (
	docs "go-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutes(r *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = "/"
	swagger := r.Group("/swagger") 
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}