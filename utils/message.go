package utils

/*
Criar a pasta utils/assets e dentro dela criar o arquivo message.ini com o seguinte conteúdo exemplo:
[geral]
error.not_found = Personalidade não encontrada
error.decode_json = Erro ao decodificar JSON

Para carregar mensagens de um arquivo INI, você pode usar a biblioteca `gopkg.in/ini.v1`.
Comando: go get gopkg.in/ini.v1
*/

import (
	"log"
	"path/filepath"
	"runtime"

	"gopkg.in/ini.v1"
)

var config *ini.File

func LoadMessages() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	filePath := filepath.Join(basePath, "assets", "message.ini")

	var err error
	config, err = ini.Load(filePath)
	if err != nil {
		log.Fatalf("Erro ao carregar mensagens: %v", err)
	}
}

func GetMessage(section, key string) string {
	sec := config.Section(section)
	return sec.Key(key).MustString("Mensagem padrão não definida")
}
