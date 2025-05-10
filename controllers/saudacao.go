package controllers

// Comando completo do swag init: swag init --parseDependency --parseInternal --parseDepth 1

import (
	"github.com/gin-gonic/gin"
	// Removed unused import as it caused a compilation error
)

// SaudacaoResponse representa a resposta do endpoint de saudação
type SaudacaoResponse struct {
	Message string `json:"mensagem"`
}

// Struct HTTPError é criada para representar a mensagem de erro em na doc do swagger
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"mensagem"`
}



// Saudacao godoc
//
//	@Summary		Endpoint de saudação ao usuário
//	@Description	Exibe uma saudação ao usuário que insere o nome como parâmetro
//	@Tags			Saudacao
//	@Produce		json
//	@Security		BasicAuth
//	@Param			nome	path		string	true	"Nome"
//	@Success		200	{object}	SaudacaoResponse
//	@Failure		400	{object}	HTTPError
//	@Failure		404	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
//	@Router			/api/saudacao/{nome} [get]
func Saudacao(c *gin.Context) {
	nome := c.Param("nome")
	c.JSON(200, gin.H{
		"mensagem": "Olá, " + nome + "! Saudações 👋",
	})
}