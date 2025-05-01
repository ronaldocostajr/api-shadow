package middleware

import (
	"fmt"
	"go-api/database"
	"go-api/models"
	"go-api/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)


func CheckDays() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailInterface, exists := c.Get("email")
		if !exists {
			c.JSON(400, gin.H{"message": "Usuário não autenticado"})
			return
		}

		fl_check_day, err := GetCheckDayByEmail(emailInterface.(string))

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": utils.GetMessage("geral", "error.internal"),
			})
			c.Abort()
			return
		}else if strings.TrimSpace(strings.ToUpper(fl_check_day[0])) == "S" {
			fmt.Println("FL_CHECK_DAY: ", fl_check_day[0])
			fl_ativo, err := GetDayByDate("30/04/2025")
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": utils.GetMessage("geral", "error.internal"),
				})
				c.Abort()
				return
			}else if strings.TrimSpace(strings.ToUpper(fl_ativo[0])) == "F" {
				c.JSON(http.StatusUnavailableForLegalReasons, gin.H{
					"message": utils.GetMessage("geral", "error.unauthorized_holiday"),
				})
				c.Abort()
				return
			}else {
				fmt.Println("FL_ATIVO ELSE: ", fl_ativo[0])
			}
		}else {
			fmt.Println("FL_CHECK_DAY ELSE: ", fl_check_day[0])
		}

		c.Next()
	}
}

func GetCheckDayByEmail(email string) ([]string, error) {
	var roles []string
	err := database.DB.Model(&models.Tb_log_usuario{}).
		Where("DS_EMAIL = ?", email).
		Pluck("FL_CHECK_DAY", &roles).Error
	return roles, err
}

func GetDayByDate(date string) ([]string, error) {
	var roles []string
	fmt.Println("ENTROU AQUI NO GET DAY BY DATE")
	// Supondo que `date` seja uma string:
	dateStr := "2025-04-30" // exemplo
	parsedDate, _ := time.Parse("2006-01-02", dateStr)
	// if err != nil {
	// 	return errors.New("data inválida: " + err.Error())
	// }
	err := database.DB.Model(&models.Tb_api_dia_feriado{}).
		Where("TRUNC(DT_DIA_FERIADO) = TRUNC(?)", parsedDate).
		Pluck("FL_ATIVO", &roles).Error
	return roles, err
}
