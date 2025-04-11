package models

type Tb_pais struct {
	Country_code string `gorm:"column:country_code" json:"country_code" gorm:"primaryKey"  gorm:"->" validate:"max=7" `
	Country_name string `gorm:"column:country_name" json:"country_name" gorm:"->" validate:"max=7" `
}

//Nome da tabela no banco de dados
func (Tb_pais) TableName() string {
	return "tb_pais"
}

