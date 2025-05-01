package models

type Vw_columns struct {
	Surrogate_key   int    `gorm:"column:SURROGATE_KEY" json:"surrogate_key"`
	Column_id       int    `gorm:"column:COLUMN_ID" json:"column_id"`
	Owner           string `gorm:"column:OWNER" json:"owner"`
	Table_name      string `gorm:"column:TABLE_NAME" json:"table_name"`
	Column_name     string `gorm:"column:COLUMN_NAME" json:"column_name"`
	Data_type       string `gorm:"column:DATA_TYPE" json:"data_type"`
	Data_type_front string `gorm:"column:DATA_TYPE_FRONT" json:"data_type_front"`
	Data_length     int    `gorm:"column:DATA_LENGTH" json:"data_length"`
	Data_scale      int    `gorm:"column:DATA_SCALE" json:"data_scale"`
	Nullable        string `gorm:"column:NULLABLE" json:"nullable"`
	Comments        string `gorm:"column:COMMENTS" json:"comments"`
	Field_label     string `gorm:"column:FIELD_LABEL" json:"field_label"`
	Field_comment   string `gorm:"column:FIELD_COMMENT" json:"field_comment"`
}

// Nome da tabela no banco
func (Vw_columns) TableName() string {
	return "VW_COLUMNS"
}
