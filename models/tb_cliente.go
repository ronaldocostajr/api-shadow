package models

type Tb_cliente struct {
	Sq_cliente string `gorm:"column:sq_cliente" json:"sq_cliente" gorm:"primaryKey"  gorm:"->" validate:"required,max=7" `
	Nm_cliente string `gorm:"column:nm_cliente" json:"nm_cliente" gorm:"->" validate:"required,max=200" `
	Fl_estado_civil string `gorm:"column:fl_estado_civil" json:"fl_estado_civil" gorm:"->" validate:"max=30" `
	Fl_sexo string `gorm:"column:fl_sexo" json:"fl_sexo" gorm:"->" validate:"max=1" `
	Tp_pessoa string `gorm:"column:tp_pessoa" json:"tp_pessoa" gorm:"->" validate:"max=1" `
	Dt_nascimento string `gorm:"column:dt_nascimento" json:"dt_nascimento" gorm:"->" validate:"max=7" `
	Nm_state string `gorm:"column:nm_state" json:"nm_state" gorm:"->" validate:"max=40" `
	Nm_country string `gorm:"column:nm_country" json:"nm_country" gorm:"->" validate:"max=20" `
	Cd_nacionalidade string `gorm:"column:cd_nacionalidade" json:"cd_nacionalidade" gorm:"->" validate:"max=30" `
	Ds_company string `gorm:"column:ds_company" json:"ds_company" gorm:"->" validate:"max=100" `
	Nm_contact string `gorm:"column:nm_contact" json:"nm_contact" gorm:"->" validate:"max=2000" `
	Nu_cnpj string `gorm:"column:nu_cnpj" json:"nu_cnpj" gorm:"->" validate:"max=50" `
	Fl_cnpj_valido string `gorm:"column:fl_cnpj_valido" json:"fl_cnpj_valido" gorm:"->" validate:"max=1" `
	Nu_cpf string `gorm:"column:nu_cpf" json:"nu_cpf" gorm:"->" validate:"max=50" `
	Fl_cpf_valido string `gorm:"column:fl_cpf_valido" json:"fl_cpf_valido" gorm:"->" validate:"max=1" `
	Nu_rg string `gorm:"column:nu_rg" json:"nu_rg" gorm:"->" validate:"max=50" `
	Nu_passaporte string `gorm:"column:nu_passaporte" json:"nu_passaporte" gorm:"->" validate:"max=50" `
	Ds_emails string `gorm:"column:ds_emails" json:"ds_emails" gorm:"->" validate:"max=1500" `
	Sq_profissao int `gorm:"column:sq_profissao" json:"sq_profissao" gorm:"->" validate:"max=7" `
	Nm_profissao string `gorm:"column:nm_profissao" json:"nm_profissao" gorm:"->" validate:"max=350" `
	Cd_uf string `gorm:"column:cd_uf" json:"cd_uf" gorm:"->" validate:"max=20" `
	Nm_endereco string `gorm:"column:nm_endereco" json:"nm_endereco" gorm:"->" validate:"max=1500" `
	Nm_bairro string `gorm:"column:nm_bairro" json:"nm_bairro" gorm:"->" validate:"max=200" `
	Nu_telefones string `gorm:"column:nu_telefones" json:"nu_telefones" gorm:"->" validate:"max=500" `
	Nu_cep string `gorm:"column:nu_cep" json:"nu_cep" gorm:"->" validate:"max=30" `
	Fl_cep_valido string `gorm:"column:fl_cep_valido" json:"fl_cep_valido" gorm:"->" validate:"max=1" `
	Nm_cidade string `gorm:"column:nm_cidade" json:"nm_cidade" gorm:"->" validate:"max=250" `
	Sq_municipio int `gorm:"column:sq_municipio" json:"sq_municipio" gorm:"->" validate:"max=7" `
	Vl_score int `gorm:"column:vl_score" json:"vl_score" gorm:"->" validate:"max=7" `
	Ds_sistema string `gorm:"column:ds_sistema" json:"ds_sistema" gorm:"->" validate:"max=50" `
	Ds_site string `gorm:"column:ds_site" json:"ds_site" gorm:"->" validate:"max=2" `
	Nm_tabela string `gorm:"column:nm_tabela" json:"nm_tabela" gorm:"->" validate:"max=150" `
	Ds_sistema_importacao string `gorm:"column:ds_sistema_importacao" json:"ds_sistema_importacao" gorm:"->" validate:"max=150" `
	Qt_processamento int `gorm:"column:qt_processamento" json:"qt_processamento" gorm:"->" validate:"max=7" `
	Dt_ingestao string `gorm:"column:dt_ingestao" json:"dt_ingestao" gorm:"->" validate:"max=7" `
	Dt_alteracao string `gorm:"column:dt_alteracao" json:"dt_alteracao" gorm:"->" validate:"max=7" `
	Sq_parametro int `gorm:"column:sq_parametro" json:"sq_parametro" gorm:"->" validate:"max=7" `
	Id_status_integracao int `gorm:"column:id_status_integracao" json:"id_status_integracao" gorm:"->" validate:"max=7" `
	Nm_mae string `gorm:"column:nm_mae" json:"nm_mae" gorm:"->" validate:"max=200" `
	Nm_pai string `gorm:"column:nm_pai" json:"nm_pai" gorm:"->" validate:"max=200" `
	Nr_endereco string `gorm:"column:nr_endereco" json:"nr_endereco" gorm:"->" validate:"max=20" `
	Ds_cep_tipo string `gorm:"column:ds_cep_tipo" json:"ds_cep_tipo" gorm:"->" validate:"max=20" `
	Ds_cep_logradouro_completo string `gorm:"column:ds_cep_logradouro_completo" json:"ds_cep_logradouro_completo" gorm:"->" validate:"max=200" `
	Ds_cep_logradouro_complemento string `gorm:"column:ds_cep_logradouro_complemento" json:"ds_cep_logradouro_complemento" gorm:"->" validate:"max=100" `
	Ds_cep_logradouro string `gorm:"column:ds_cep_logradouro" json:"ds_cep_logradouro" gorm:"->" validate:"max=120" `
	Ds_cep_bairro string `gorm:"column:ds_cep_bairro" json:"ds_cep_bairro" gorm:"->" validate:"max=100" `
	Ds_cep_cidade string `gorm:"column:ds_cep_cidade" json:"ds_cep_cidade" gorm:"->" validate:"max=32" `
	Ds_cep_uf string `gorm:"column:ds_cep_uf" json:"ds_cep_uf" gorm:"->" validate:"max=2" `
}

//Nome da tabela no banco de dados
func (Tb_cliente) TableName() string {
	return "tb_cliente"
}

