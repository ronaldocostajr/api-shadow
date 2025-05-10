//Desenvolvedor: Thiago Leite
//Versão: 1.0.0 V
//Compilação: 2025-04-28 17:00:05.7567253 -0300 -03 m=+6.317699701
package models

type Tb_log_usuario struct {
	Id_usuario int `gorm:"column:ID_USUARIO" json:"id_usuario" gorm:"primaryKey"  gorm:"->" validate:"required,max=22" `
	Nm_usuario string `gorm:"column:NM_USUARIO" json:"nm_usuario" gorm:"->" validate:"required,max=300" `
	Ds_email string `gorm:"column:DS_EMAIL" json:"ds_email" gorm:"->" validate:"required,max=300" `
	Hs_password string `gorm:"column:HS_PASSWORD" json:"hs_password" gorm:"->" validate:"required,max=300" `
	Dt_ingestao string `gorm:"column:DT_INGESTAO" json:"dt_ingestao" gorm:"->" validate:"max=7" `
	Dt_alteracao string `gorm:"column:DT_ALTERACAO" json:"dt_alteracao" gorm:"->" validate:"max=7" `
	Fl_ativo string `gorm:"column:FL_ATIVO" json:"fl_ativo" gorm:"->" validate:"max=1" `
	Sq_operador int `gorm:"column:SQ_OPERADOR" json:"sq_operador" gorm:"->" validate:"max=22" `
	Sq_setor int `gorm:"column:SQ_SETOR" json:"sq_setor" gorm:"->" validate:"max=22" `
	Nm_role string `gorm:"column:NM_ROLE" json:"nm_role" gorm:"->" validate:"max=1000" `
	Fl_check_day string `gorm:"column:FL_CHECK_DAY" json:"fl_check_day" gorm:"->" validate:"max=1" `
	Fl_check_ip string `gorm:"column:FL_CHECK_IP" json:"fl_check_ip" gorm:"->" validate:"max=1" `
}

//Nome da tabela no banco de dados
func (Tb_log_usuario) TableName() string {
	return "TB_LOG_USUARIO"
}

