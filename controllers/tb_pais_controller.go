package controllers

import (
	"go-api/database"
	"go-api/logSystem"
	"go-api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTb_pais(c *gin.Context) {
	startTime := time.Now()
	var tb_pais []models.Tb_pais

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset := (page - 1) * limit
	mdbUrl := c.Request.Host + c.Request.URL.Path

	// Filtros
	country_name := c.Query("country_name")
	mdbParameterField := []string{
		"country_name&%like%",
		c.Query("country_name"),
	}

	mdbParameterDate := []string{}

	query := database.DB.Model(&models.Tb_pais{})

	if country_name != "" {
		query = query.Where("country_name ILIKE ?", "%"+country_name+"%")
	}

	err := query.Offset(offset).Limit(limit).Find(&tb_pais).Error
	endTime := time.Now()
	if err != nil {
		errMd0 := logSystem.WriteLogMongoDB("SHADOW", "shadow_generico", "GetTb_pais", "ronaldo.costa@aviva.com.br", "generico", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "404")
		if errMd0 != nil {
			logSystem.WriteLogFile("200:SHADOW:shadow_generico:generico:GetTb_pais:200:ronaldo.costa@aviva.com.br:" + startTime.String() + ":" + endTime.String())
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar Pais"})
		return
	} else {
		errMd := logSystem.WriteLogMongoDB("SHADOW", "shadow_generico", "GetTb_pais", "ronaldo.costa@aviva.com.br", "generico", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "200")
		logSystem.WriteLogFile("SHADOW:shadow_generico:GetTb_pais:ronaldo.costa@aviva.com.br:generico:200" + startTime.String())
		if errMd != nil {
			logSystem.WriteLogFile("SHADOW:shadow_generico:GetTb_pais:ronaldo.costa@aviva.com.br:generico:200" + startTime.String())
		}
	}
	c.JSON(http.StatusOK, tb_pais)
}
