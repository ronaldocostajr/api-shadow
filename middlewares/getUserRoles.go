package middleware

import (
	"go-api/database"
	"go-api/models"
	"go-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailInterface, exists := c.Get("email")
		if !exists {
			c.JSON(400, gin.H{"message": "Usuário não autenticado"})
			return
		}


		roles, err := GetRolesByEmail(emailInterface.(string))
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": utils.GetMessage("geral", "error.internal"),
			})
			c.Abort()
			return
		}
		c.Set("roles", roles)
		//c.Set("userFlCheckDays", fl_check_day)
	}
}

func GetRolesByEmail(email string) ([]string, error) {
	var roles []string
	err := database.DB.Model(&models.Tb_log_usuario{}).
		Where("DS_EMAIL = ?", email).
		Pluck("NM_ROLE", &roles).Error
	return roles, err
}


// func GetCheckDayByEmail(email string) ([]string, error) {
// 	var roles []string
// 	err := database.DB.Model(&models.Tb_log_usuario{}).
// 		Where("DS_EMAIL = ?", email).
// 		Pluck("FL_CHECK_DAY", &roles).Error
// 	return roles, err
// }