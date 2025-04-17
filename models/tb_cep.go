//Desenvolvedor: Ronaldo Costa
//Versão: 1.0.0 V
//Compilação: 2025-04-17 08:35:03.9165257 -0300 -03 m=+5.739813101
//Comentário adicional: código adiconal
package models

type Tb_cep struct {
	Sq_cep string `gorm:"column:sq_cep" json:"sq_cep" gorm:"primaryKey"  gorm:"->" validate:"required,max=7" `
	Nu_cep string `gorm:"column:nu_cep" json:"nu_cep" gorm:"->" validate:"max=8" `
	Ds_tipo string `gorm:"column:ds_tipo" json:"ds_tipo" gorm:"->" validate:"max=20" `
	Ds_logradouro_completo string `gorm:"column:ds_logradouro_completo" json:"ds_logradouro_completo" gorm:"->" validate:"max=200" `
	Ds_logradouro_complemento string `gorm:"column:ds_logradouro_complemento" json:"ds_logradouro_complemento" gorm:"->" validate:"max=100" `
	Ds_logradouro string `gorm:"column:ds_logradouro" json:"ds_logradouro" gorm:"->" validate:"max=120" `
	Ds_bairro string `gorm:"column:ds_bairro" json:"ds_bairro" gorm:"->" validate:"max=100" `
	Ds_cidade string `gorm:"column:ds_cidade" json:"ds_cidade" gorm:"->" validate:"max=32" `
	Ds_uf string `gorm:"column:ds_uf" json:"ds_uf" gorm:"->" validate:"max=2" `
	Dt_ingestao string `gorm:"column:dt_ingestao" json:"dt_ingestao" gorm:"->" validate:"max=7" `
	Dt_alteracao string `gorm:"column:dt_alteracao" json:"dt_alteracao" gorm:"->" validate:"max=7" `
	Ds_logradouro_char string `gorm:"column:ds_logradouro_char" json:"ds_logradouro_char" gorm:"->" validate:"max=200" `
	Ds_bairro_char string `gorm:"column:ds_bairro_char" json:"ds_bairro_char" gorm:"->" validate:"max=100" `
	Ds_cidade_char string `gorm:"column:ds_cidade_char" json:"ds_cidade_char" gorm:"->" validate:"max=32" `
	Uuid string `gorm:"column:uuid" json:"uuid" gorm:"->" validate:"max=36" `
	Fl_ativo string `gorm:"column:fl_ativo" json:"fl_ativo" gorm:"->" validate:"max=1" `
}

//Nome da tabela no banco de dados
func (Tb_cep) TableName() string {
	return "tb_cep"
}

