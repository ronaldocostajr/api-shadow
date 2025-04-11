package models

type Unidade struct {
	CdUnidade int    `json:"cd_unidade" gorm:"primaryKey" validate:"required"`
	FlUnidade string `json:"fl_unidade" validate:"required,oneof=C B Os"` // Supondo que só aceite S ou N
	DsUnidade string `json:"ds_unidade" validate:"required,max=40"`
	DsSigla   string `json:"ds_sigla" validate:"required,max=20"`
}

func (Unidade) TableName() string {
	return "tb_unidade"
}
