package controllers

import (
	"log"
	"net/http"
	"strconv"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func GetUsuarios(c *gin.Context) {
	var usuarios []models.Usuario

	log.Println("Rota GET /ususario foi chamada!")
	log.Printf("HEADERS RECEBIDOS: %+v\n", c.Request.Header)

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Filtros opcionais
	nm_usuario := c.Query("nm_usuario")
	ds_usuario := c.Query("ds_usuario")

	query := database.DB.Model(&models.Usuario{})

	if nm_usuario != "" {
		query = query.Where("nm_usuario ILIKE ?", "%"+nm_usuario+"%")
	}

	if ds_usuario != "" {
		query = query.Where("ds_usuario ILIKE ?", "%"+ds_usuario+"%")
	}

	var total int64
	query.Count(&total) // conta total antes da paginação

	// Aplica paginação e busca resultados
	err := query.Offset(offset).Limit(limit).Find(&usuarios).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuarios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"limit":    limit,
		"total":    total,
		"usuarios": usuarios,
	})
}

func GetUsuarioByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var usuario models.Usuario

	if err := database.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
