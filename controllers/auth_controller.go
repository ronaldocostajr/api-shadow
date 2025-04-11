package controllers

import (
	"net/http"
	"time"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("minhachavesecreta")

func Login(c *gin.Context) {
	var input models.Usuario
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var user models.Usuario
	if err := database.DB.Where("nm_usuario = ?", input.NmUsuario).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos (usuário)"})
		return
	}

	// Comparar senha fornecida com hash
	if user.DsSenha != input.DsSenha {
		//bcrypt.CompareHashAndPassword([]byte(user.DsSenha), []byte(input.DsSenha)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos (senha)"})
		return
	}

	// Criar token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nm_usuario": user.NmUsuario,
		"cd_usuario": user.CdUsuario,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
