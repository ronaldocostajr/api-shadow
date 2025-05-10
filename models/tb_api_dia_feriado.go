//Desenvolvedor: Thiago Leite
//Versão: 1.0.0 V
//Compilação: 2025-04-30 14:49:26.7070062 -0300 -03 m=+6.640742301
package models

type Tb_api_dia_feriado struct {
	Sq_dia_feriado int `gorm:"column:SQ_DIA_FERIADO" json:"sq_dia_feriado" gorm:"primaryKey"  gorm:"->" validate:"required,max=22" `
	Dt_dia_feriado string `gorm:"column:DT_DIA_FERIADO" json:"dt_dia_feriado" gorm:"->" validate:"required,max=7" `
	Hr_inicial int `gorm:"column:HR_INICIAL" json:"hr_inicial" gorm:"->" validate:"max=22" `
	Hr_final int `gorm:"column:HR_FINAL" json:"hr_final" gorm:"->" validate:"max=22" `
	Ds_dia_feriado string `gorm:"column:DS_DIA_FERIADO" json:"ds_dia_feriado" gorm:"->" validate:"required,max=50" `
	Ds_localidade string `gorm:"column:DS_LOCALIDADE" json:"ds_localidade" gorm:"->" validate:"max=100" `
	Fl_ativo string `gorm:"column:FL_ATIVO" json:"fl_ativo" gorm:"->" validate:"required,max=1" `
	Id_usuario_liberado string `gorm:"column:ID_USUARIO_LIBERADO" json:"id_usuario_liberado" gorm:"->" validate:"max=200" `
	Ds_observacao string `gorm:"column:DS_OBSERVACAO" json:"ds_observacao" gorm:"->" validate:"max=2000" `
	Dt_ingestao string `gorm:"column:DT_INGESTAO" json:"dt_ingestao" gorm:"->" validate:"max=7" `
	Dt_alteracao string `gorm:"column:DT_ALTERACAO" json:"dt_alteracao" gorm:"->" validate:"max=7" `
}

//Nome da tabela no banco de dados
func (Tb_api_dia_feriado) TableName() string {
	return "TB_API_DIA_FERIADO"
}

