//Desenvolvedor: Ronaldo Costa
//Versão: 1.0.0 V
//Compilação: 2025-04-28 10:50:12.882475 -0300 -03 m=+17.982541301
//Comentário adicional: código adiconal
package controllers
import (
	"net/http"
	"strconv"
	"time"
	"strings"
	"go-api/database"
	"go-api/models"
	"go-api/logSystem"
	"github.com/gin-gonic/gin"
)

func GetTb_cep(c *gin.Context) {
	userRoles := "RL_ADMIN"
	if  !strings.Contains(userRoles, "RL_ADMIN") && !strings.Contains(userRoles, "RL_CONTROLADORIA") && !strings.Contains(userRoles, "RL_TESOURARIA") {
		c.JSON(400, "Sem direito a acessar a API")
		return
	}
	startTime := time.Now()
	var tb_cep []models.Tb_cep

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if limit < 0 || limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit
	mdbUrl := c.Request.Host + c.Request.URL.Path

	// Filtros passados na URL
	nu_cep := c.Query("nu_cep")
	ds_logradouro := c.Query("ds_logradouro")

	//Parâmetros incluídos no log do sistema
	mdbParameterField := []string{
		"nu_cep&=",
		c.Query("nu_cep"),
		"ds_logradouro&like%",
		c.Query("ds_logradouro"),
	}

	// Parâmetros passados na URL
	false := c.Query("false")

	//Parâmetros incluídos no log do sistema
	mdbParameterDate := []string{
		"false",
		c.Query("false"),
	}

	//Montando o modelo e preparando os filtros
	query := database.DB.Model(&models.Tb_cep{})

	if nu_cep != "" {
		query = query.Where("nu_cep = ?", nu_cep)
	}

	if ds_logradouro != "" {
		query = query.Where("ds_logradouro ILIKE ?", ds_logradouro+"%")
	}

	if len(false) > 0  {
		query = query.Where("false =  ? ", false +" 00:00:00")
	}

	// Efetua a consulta no banco de dados
	err := query.Offset(offset).Limit(limit).Find(&tb_cep).Error
	endTime := time.Now()

	// Log do sistema - OBRIGATÓRIO
	if err != nil {
		errMd0 := logSystem.WriteLogMongoDB("SHADOW", "shadow_financeiro", "GetTb_cep", "ronaldo.costa@aviva.com.br", "tesouraria", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "404")
		if errMd0 != nil {
			logSystem.WriteLogFile(":404:SHADOW:shadow_financeiro:tesouraria:ronaldo.costa@aviva.com.br:"+startTime.String()+":"+endTime.String())
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar a tabela : CEP"})
		return
	} else {
		errMd1 := logSystem.WriteLogMongoDB("SHADOW", "shadow_financeiro", "GetTb_cep", "ronaldo.costa@aviva.com.br", "tesouraria", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, "202")
		if errMd1 != nil {
			logSystem.WriteLogFile(":202:SHADOW:shadow_financeiro:tesouraria:ronaldo.costa@aviva.com.br:"+startTime.String()+":"+endTime.String())
		}
	}
	c.JSON(http.StatusOK, tb_cep)
}
