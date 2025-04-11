package routes

import (
	"go-api/generator"

	"github.com/gin-gonic/gin"
)

func WriteGeneratorRoutes(r *gin.RouterGroup) {
	writeGenerator := r.Group("/writeclass")
	{
		writeGenerator.GET("/", generator.GetWriteGenerator)
	}
}
