package models

type Usuario struct {
	CdUsuario int    `gorm:"column:cd_usuario;primaryKey" json:"cd_usuario"`
	NmUsuario string `gorm:"column:nm_usuario" json:"nm_usuario" binding:"required"`
	FlUsuario string `gorm:"column:fl_usuario" json:"fl_usuario"`
	DsUsuario string `gorm:"column:ds_usuario" json:"ds_usuario"`
	DsSenha   string `gorm:"column:ds_senha" json:"ds_senha" binding:"required"`
}

// Nome da tabela no banco
func (Usuario) TableName() string {
	return "tb_usuario"
}
