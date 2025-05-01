//Desenvolvedor: Thiago Leite
//Versão: 1.0.0 V
//Compilação: 2025-04-30 13:42:26.6186388 -0300 -03 m=+6464.441404201
package models

type Tb_cdp_cep struct {
	Sq_cep int `gorm:"column:SQ_CEP" json:"sq_cep" gorm:"primaryKey"  gorm:"->" validate:"required,max=22" `
	Nu_cep string `gorm:"column:NU_CEP" json:"nu_cep" gorm:"->" validate:"max=8" `
	Ds_tipo string `gorm:"column:DS_TIPO" json:"ds_tipo" gorm:"->" validate:"max=20" `
	Ds_logradouro_completo string `gorm:"column:DS_LOGRADOURO_COMPLETO" json:"ds_logradouro_completo" gorm:"->" validate:"max=200" `
	Ds_logradouro_complemento string `gorm:"column:DS_LOGRADOURO_COMPLEMENTO" json:"ds_logradouro_complemento" gorm:"->" validate:"max=100" `
	Ds_logradouro string `gorm:"column:DS_LOGRADOURO" json:"ds_logradouro" gorm:"->" validate:"max=120" `
	Ds_bairro string `gorm:"column:DS_BAIRRO" json:"ds_bairro" gorm:"->" validate:"max=100" `
	Ds_cidade string `gorm:"column:DS_CIDADE" json:"ds_cidade" gorm:"->" validate:"max=32" `
	Ds_uf string `gorm:"column:DS_UF" json:"ds_uf" gorm:"->" validate:"max=2" `
	Dt_ingestao string `gorm:"column:DT_INGESTAO" json:"dt_ingestao" gorm:"->" validate:"required,max=7" `
	Dt_alteracao string `gorm:"column:DT_ALTERACAO" json:"dt_alteracao" gorm:"->" validate:"max=7" `
	Ds_logradouro_char string `gorm:"column:DS_LOGRADOURO_CHAR" json:"ds_logradouro_char" gorm:"->" validate:"max=200" `
	Ds_bairro_char string `gorm:"column:DS_BAIRRO_CHAR" json:"ds_bairro_char" gorm:"->" validate:"max=100" `
	Ds_cidade_char string `gorm:"column:DS_CIDADE_CHAR" json:"ds_cidade_char" gorm:"->" validate:"max=32" `
	Uuid string `gorm:"column:UUID" json:"uuid" gorm:"->" validate:"max=36" `
}

//Nome da tabela no banco de dados
func (Tb_cdp_cep) TableName() string {
	return "TB_CDP_CEP"
}

