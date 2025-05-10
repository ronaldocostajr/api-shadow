//Desenvolvedor: Thiago Leite
//Versão: 1.0.0 V
//Compilação: 2025-04-28 17:00:05.7622516 -0300 -03 m=+6.323226101
package controllers
import (
	"net/http"
	"strconv"
	"time"
	"go-api/database"
	"go-api/models"
	"go-api/logSystem"
	"github.com/gin-gonic/gin"
)

func GetTb_log_usuario(c *gin.Context) {
	startTime := time.Now()
	var tb_log_usuario []models.Tb_log_usuario

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset := (page - 1) * limit
	mdbUrl := c.Request.Host + c.Request.URL.Path

	// Filtros passados na URL
	ds_email := c.Query("ds_email")

	//Parâmetros incluídos no log do sistema
	mdbParameterField := []string{
		"ds_email&=",
		c.Query("ds_email"),
	}

	// Parâmetros passados na URL
	false := c.Query("false")

	//Parâmetros incluídos no log do sistema
	mdbParameterDate := []string{
		"false",
		c.Query("false"),
	}

	//Montando o modelo e preparando os filtros
	query := database.DB.Model(&models.Tb_log_usuario{})

	if ds_email != "" {
		query = query.Where("ds_email = ?", ds_email)
	}

	if len(false) > 0  {
		query = query.Where("false =  ? ", false +" 00:00:00")
	}

	// Efetua a consulta no banco de dados
	err := query.Offset(offset).Limit(limit).Find(&tb_log_usuario).Error
	endTime := time.Now()

	// Log do sistema - OBRIGATÓRIO
	if err != nil {
		errMd0 := logSystem.WriteLogMongoDB("SHADOW", "shadow_controller", "GetTb_log_usuario", "ronaldo.costa@aviva.com.br", "controller", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "404")
		if errMd0 != nil {
			logSystem.WriteLogFile(":404:SHADOW:shadow_controller:controller:ronaldo.costa@aviva.com.br:"+startTime.String()+":"+endTime.String())
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar a tabela : USUARIO"})
		return
	} else {
		errMd1 := logSystem.WriteLogMongoDB("SHADOW", "shadow_controller", "GetTb_log_usuario", "ronaldo.costa@aviva.com.br", "controller", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "202")
		if errMd1 != nil {
			logSystem.WriteLogFile(":202:SHADOW:shadow_controller:controller:ronaldo.costa@aviva.com.br:"+startTime.String()+":"+endTime.String())
		}
	}
	c.JSON(http.StatusOK, tb_log_usuario)
}
