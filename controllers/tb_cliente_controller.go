package controllers
import (
	"net/http"
	"strconv"
	"go-api/database"
	"go-api/models"
	"github.com/gin-gonic/gin"
)

func GetTb_cliente(c *gin.Context) {
	var tb_cliente []models.Tb_cliente

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset := (page - 1) * limit

	// Filtros
	nm_cliente := c.Query("nm_cliente")
	dt_aniversario := c.Query("dt_aniversario")
	fl_ativo := c.Query("fl_ativo")
	nu_cep := c.Query("nu_cep")
	fl_estado_civil := c.Query("fl_estado_civil")
	tp_pessoa := c.Query("tp_pessoa")
	dt_start := c.Query("dt_start")
	dt_end := c.Query("dt_end")

	query := database.DB.Model(&models.Tb_cliente{})

	if nm_cliente != "" {
		query = query.Where("nm_cliente ILIKE ?", nm_cliente+"%")
	}

	if dt_aniversario != "" {
		query = query.Where("dt_aniversario >= ?", dt_aniversario)
	}

	if fl_ativo != "" {
		query = query.Where("fl_ativo ILIKE ?", fl_ativo+"%")
	}

	if nu_cep != "" {
		query = query.Where("nu_cep ILIKE ?", "%"+nu_cep+"%")
	}

	if fl_estado_civil != "" {
		query = query.Where("fl_estado_civil = ?", fl_estado_civil)
	}

	if tp_pessoa != "" {
		query = query.Where("tp_pessoa <> ?", tp_pessoa)
	}

	if dt_start != "" && dt_end != "" {
		query = query.Where("dt_nascimento BETWEEN ? AND ?", dt_start +" 00:00:00", dt_end +" 00:00:00")
	}

	err := query.Offset(offset).Limit(limit).Distinct("nm_cliente").Find(&tb_cliente).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar Cliente"})
		return
	}

	c.JSON(http.StatusOK, tb_cliente)
}
