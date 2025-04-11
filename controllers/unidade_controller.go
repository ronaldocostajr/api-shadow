package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

func GetUnidades(c *gin.Context) {
	var unidades []models.Unidade

	log.Println("Rota GET /unidades foi chamada!")
	log.Printf("HEADERS RECEBIDOS: %+v\n", c.Request.Header)

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Filtros opcionais
	flUnidade := c.Query("fl_unidade")
	dsUnidade := c.Query("ds_unidade")
	dsSigla := c.Query("ds_sigla")
	log.Println(dsUnidade)
	query := database.DB.Model(&models.Unidade{})

	if dsUnidade != "" {
		query = query.Where("ds_unidade ILIKE ?", "%"+dsUnidade+"%")
	}

	if dsSigla != "" {
		query = query.Where("ds_sigla ILIKE ?", "%"+dsSigla+"%")
	}

	if flUnidade != "" {
		query = query.Where("fl_unidade ILIKE ?", "%"+flUnidade+"%")
	}

	// var total int64
	// query.Count(&total) // conta total antes da paginação

	// Aplica paginação e busca resultados
	err := query.Offset(offset).Limit(limit).Find(&unidades).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar unidades"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"page":     page,
	// 	"limit":    limit,
	// 	"total":    total,
	// 	"unidades": unidades,
	// })

	c.JSON(http.StatusOK, unidades)
}

func GetUnidadeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var unidade models.Unidade

	if err := database.DB.First(&unidade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unidade não encontrada"})
		return
	}

	c.JSON(http.StatusOK, unidade)
}

func CreateUnidade(c *gin.Context) {
	var unidade models.Unidade
	var validate = validator.New()

	// Bind JSON
	if err := c.ShouldBindJSON(&unidade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Validação com validator
	if err := validate.Struct(&unidade); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Campo %s inválido (%s)", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
		return
	}

	// Salva no banco
	if err := database.DB.Create(&unidade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar unidade"})
		return
	}

	c.JSON(http.StatusCreated, unidade)
}

func UpdateUnidade(c *gin.Context) {
	id := c.Param("id")
	var unidade models.Unidade
	var validate = validator.New()

	// Verifica se a unidade existe
	if err := database.DB.First(&unidade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unidade não encontrada"})
		return
	}

	// Faz bind dos dados recebidos
	if err := c.ShouldBindJSON(&unidade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Valida os dados
	if err := validate.Struct(&unidade); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Campo %s inválido (%s)", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
		return
	}

	// Atualiza no banco
	if err := database.DB.Model(&unidade).Where("cd_unidade = ?", id).Updates(unidade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar unidade"})
		return
	}

	c.JSON(http.StatusOK, unidade)
}

func DeleteUnidade(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var unidade models.Unidade

	if err := database.DB.First(&unidade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unidade não encontrada"})
		return
	}

	database.DB.Delete(&unidade)

	c.JSON(http.StatusOK, gin.H{"message": "Unidade deletada com sucesso"})
}
