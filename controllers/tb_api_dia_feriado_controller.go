// Desenvolvedor: Thiago Leite
// Versão: 1.0.0 V
// Compilação: 2025-04-30 14:49:26.711964 -0300 -03 m=+6.645700201
package controllers

import (
	"go-api/database"
	"go-api/logSystem"
	"go-api/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTb_api_dia_feriado(c *gin.Context) {

	// Pega as Roles do usuário vindas do middleware
	rolesInterface, exists := c.Get("roles")
	if !exists {
		c.JSON(400, gin.H{"message": "Usuário não autenticado"})
		return
	}
	roles, _ := rolesInterface.([]string)
	userRoles := strings.Split(roles[0], ",")

	// Roles requeridas do endpoint
	userRolesRequired := strings.Split("RL_ADMIN", ",")

	// Busca itens em comum nos slices
	roleSet := make(map[string]struct{})
	boolRoles := false
	for _, item := range userRoles {
		roleSet[item] = struct{}{}
	}

	for _, item := range userRolesRequired {
		if _, exists := roleSet[item]; exists {

			boolRoles = true
		}
	}

	if !boolRoles{
		c.JSON(400, "Sem direito a acessar a API")
		return
	}
	startTime := time.Now()
	var tb_api_dia_feriado []models.Tb_api_dia_feriado

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset := (page - 1) * limit
	mdbUrl := c.Request.Host + c.Request.URL.Path

	// Filtros passados na URL
	dt_dia_feriado := c.Query("dt_dia_feriado")

	//Parâmetros incluídos no log do sistema
	mdbParameterField := []string{
		"dt_dia_feriado&=",
		c.Query("dt_dia_feriado"),
	}

	// Parâmetros passados na URL
	false := c.Query("false")

	//Parâmetros incluídos no log do sistema
	mdbParameterDate := []string{
		"false",
		c.Query("false"),
	}

	//Montando o modelo e preparando os filtros
	query := database.DB.Model(&models.Tb_api_dia_feriado{})

	if dt_dia_feriado != "" {
		query = query.Where("dt_dia_feriado = ?", dt_dia_feriado)
	}

	if len(false) > 0  {
		query = query.Where("false =  ? ", false +" 00:00:00")
	}

	// Efetua a consulta no banco de dados
	err := query.Offset(offset).Limit(limit).Find(&tb_api_dia_feriado).Error
	endTime := time.Now()

	// Log do sistema - OBRIGATÓRIO
	if err != nil {
		errMd0 := logSystem.WriteLogMongoDB("SHADOW", "shadow_controller", "GetTb_api_dia_feriado", "ronaldo.costa@aviva.com.br", "controller", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "404")
		if errMd0 != nil {
			logSystem.WriteLogFile(":404:SHADOW:shadow_controller:controller:ronaldo.costa@aviva.com.br:"+startTime.String()+":"+endTime.String())
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar a tabela : DIAS_FERIADO"})
		return
	} else {
		errMd1 := logSystem.WriteLogMongoDB("SHADOW", "shadow_controller", "GetTb_api_dia_feriado", "ronaldo.costa@aviva.com.br", "controller", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "202")
		if errMd1 != nil {
			logSystem.WriteLogFile(":202:SHADOW:shadow_controller:controller:ronaldo.costa@aviva.com.br:"+startTime.String()+":"+endTime.String())
		}
	}
	c.JSON(http.StatusOK, tb_api_dia_feriado)
}
