package models

type Vw_columns struct {
	Surrogate_key   int    `gorm:"column:surrogate_key" json:"surrogate_key"`
	Owner           string `gorm:"column:owner" json:"owner"`
	Table_name      string `gorm:"table_name:fl_usuario" json:"table_name"`
	Column_name     string `gorm:"column:column_name" json:"column_name"`
	Data_type       string `gorm:"data_type:fl_usuario" json:"data_type"`
	Data_type_front string `gorm:"column:data_type_front" json:"data_type_front"`
	Data_length     int    `gorm:"column:data_length" json:"data_length"`
	Data_scale      int    `gorm:"data_scale:fl_usuario" json:"data_scale"`
	Nullable        string `gorm:"column:nullable" json:"nullable"`
	Comments        string `gorm:"column:comments" json:"comments"`
	Field_label     string `gorm:"column:field_label" json:"field_label"`
	Field_comment   string `gorm:"column:field_comment" json:"field_comment"`
}

// Nome da tabela no banco
func (Vw_columns) TableName() string {
	return "vw_columns"
}
