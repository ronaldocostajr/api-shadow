package generator

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"go-api/database"
	"go-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetWriteGenerator(c *gin.Context) {
	err := godotenv.Load("generator/generator.properties")
	if err != nil {
		panic(err)
	}

	var vw_columns []models.Vw_columns

	// Paginação: valores padrão
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1000"))
	offset := (page - 1) * limit

	// Filtros opcionais
	owner := c.Query("owner")
	table_name := c.Query("table_name")

	query := database.DB.Model(&models.Vw_columns{})

	if owner != "" {
		query = query.Where("owner = ?", owner)
	}

	if table_name != "" {
		query = query.Where("table_name = ?", table_name)
	}

	err = query.Offset(offset).Limit(limit).Find(&vw_columns).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar Vw_columns"})
		return
	}

	if len(vw_columns) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar a tabela"})
		return
	}

	var linha []string
	if os.Getenv("generator.comments") == "true" {
		if os.Getenv("generator.comments.author") != "false" {
			linha = append(linha, "//Desenvolvedor: "+os.Getenv("generator.comments.author"))
		}
		if os.Getenv("generator.comments.version") != "false" {
			linha = append(linha, "//Versão: "+os.Getenv("generator.comments.version"))
		}
		if os.Getenv("generator.comments.data") == "true" {
			linha = append(linha, "//Compilação: "+time.Now().String())
		}
		if os.Getenv("generator.comments.additional") != "false" {
			linha = append(linha, "//Comentário adicional: "+os.Getenv("generator.comments.additional"))
		}
	}
	linha = append(linha, "package models")
	linha = append(linha, "")
	var tabela string = ""
	for _, v := range vw_columns {
		tabela = v.Table_name
		linha = append(linha, "type "+strings.Title(table_name)+" struct {")
		break
	}
	
	primaryKey := os.Getenv("fields.primaryKey")
	var textPrimarykey string = ""
	var validate = ""
	var readOnly = ""
	if os.Getenv("fields.readOnly") == "true" {
		readOnly = " gorm:\"->\""
	} else {
		readOnly = ""
	}
	for _, v := range vw_columns {

		if v.Nullable == "N" {
			validate = " validate:\"required,"
		} else {
			validate = " validate:\""
		}

		if os.Getenv("fields.maxSize") == "true" {
			validate += "max=" + strconv.Itoa(v.Data_length)
		} else {
			validate = ""
		}

		if len(validate) == 10 {
			validate = ""
		}

		if v.Column_name == primaryKey {
			textPrimarykey = " gorm:\"primaryKey\" "
		} else {
			textPrimarykey = ""
		}
		var data_type = ""
		if v.Data_type_front == "INTEGER"{
			data_type = "int"
		}else if v.Data_type_front == "DATE"{
			data_type = "string"
		}else{
			data_type = strings.ToLower(v.Data_type_front)
		}
		linha = append(linha, "\t"+strings.Title(v.Column_name)+" "+data_type+" `gorm:\"column:"+strings.ToUpper(v.Column_name)+"\" json:\""+v.Column_name+"\""+textPrimarykey+readOnly+validate+"\" `")
	}

	linha = append(linha, "}")
	linha = append(linha, "")
	if os.Getenv("generator.comments") == "true" {
		linha = append(linha, "//Nome da tabela no banco de dados")
	}
	linha = append(linha, "func ("+strings.Title(table_name)+") TableName() string {")
	linha = append(linha, "\treturn \""+strings.ToUpper(table_name)+"\"")
	linha = append(linha, "}")
	linha = append(linha, "")

	err = escreverTexto(linha, os.Getenv("path")+"models/"+tabela+".go")
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	var linhaRoutes []string
	if os.Getenv("generator.comments") == "true" {
		if os.Getenv("generator.comments.author") != "false" {
			linhaRoutes = append(linhaRoutes, "//Desenvolvedor: "+os.Getenv("generator.comments.author"))
		}
		if os.Getenv("generator.comments.version") != "false" {
			linhaRoutes = append(linhaRoutes, "//Versão: "+os.Getenv("generator.comments.version"))
		}
		if os.Getenv("generator.comments.data") == "true" {
			linhaRoutes = append(linhaRoutes, "//Compilação: "+time.Now().String())
		}
		if os.Getenv("generator.comments.additional") != "false" {
			linhaRoutes = append(linhaRoutes, "//Comentário adicional: "+os.Getenv("generator.comments.additional"))
		}
	}
	linhaRoutes = append(linhaRoutes, "package routes")
	linhaRoutes = append(linhaRoutes, "")
	linhaRoutes = append(linhaRoutes, "import (")
	linhaRoutes = append(linhaRoutes, "\t\"go-api/controllers\"")
	linhaRoutes = append(linhaRoutes, "\t\"github.com/gin-gonic/gin\"")
	linhaRoutes = append(linhaRoutes, ")")
	linhaRoutes = append(linhaRoutes, "")
	linhaRoutes = append(linhaRoutes, "func "+strings.Title(table_name)+"Routes(r *gin.RouterGroup) {")
	linhaRoutes = append(linhaRoutes, "\t"+table_name+" := r.Group(\"/"+table_name+"\") ")
	linhaRoutes = append(linhaRoutes, "\t{")
	linhaRoutes = append(linhaRoutes, "\t\t"+table_name+".GET(\"/\", controllers.Get"+strings.Title(table_name)+")")
	linhaRoutes = append(linhaRoutes, "\t}")
	linhaRoutes = append(linhaRoutes, "}")
	linhaRoutes = append(linhaRoutes, "")

	err = escreverTexto(linhaRoutes, os.Getenv("path")+"routes/rotas/"+tabela+"_routes.go")
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	routes := "routes." + strings.Title(table_name) + "Routes(api)"

	writeRoutes(routes)

	var linhaController []string
	fieldsSearch := strings.Split(os.Getenv("fields.search"), ",")
	paramSearchDate := strings.Split(os.Getenv("param.searchDate"), ",")

	if os.Getenv("generator.comments") == "true" {
		if os.Getenv("generator.comments.author") != "false" {
			linhaController = append(linhaController, "//Desenvolvedor: "+os.Getenv("generator.comments.author"))
		}
		if os.Getenv("generator.comments.version") != "false" {
			linhaController = append(linhaController, "//Versão: "+os.Getenv("generator.comments.version"))
		}
		if os.Getenv("generator.comments.data") == "true" {
			linhaController = append(linhaController, "//Compilação: "+time.Now().String())
		}
		if os.Getenv("generator.comments.additional") != "false" {
			linhaController = append(linhaController, "//Comentário adicional: "+os.Getenv("generator.comments.additional"))
		}
	}
	linhaController = append(linhaController, "package controllers")
	linhaController = append(linhaController, "import (")
	linhaController = append(linhaController, "\t\"net/http\"")
	linhaController = append(linhaController, "\t\"strconv\"")
	linhaController = append(linhaController, "\t\"time\"")
	stringsExists := false
	if os.Getenv("security.roles") != "false" {
		linhaController = append(linhaController, "\t\"strings\"")
		stringsExists = true
	}
	for _, column_name := range fieldsSearch {
		for indexW, valueW := range column_name {
			if string(valueW) == "&" {
				if column_name[indexW+1:] == "inS" || column_name[indexW+1:] == "inI" {
					if !stringsExists {
						linhaController = append(linhaController, "\t\"strings\"")
						stringsExists = true
						break
					}
				}
			}
		}
	}
	linhaController = append(linhaController, "\t\"go-api/database\"")
	linhaController = append(linhaController, "\t\"go-api/models\"")
	linhaController = append(linhaController, "\t\"net/http\"")
	linhaController = append(linhaController, "\t\"strconv\"")
	linhaController = append(linhaController, "\t\"strings\"")
	linhaController = append(linhaController, "\t\"time\"")
	linhaController = append(linhaController, "\t\"go-api/logSystem\"")
	linhaController = append(linhaController, "\t\"github.com/gin-gonic/gin\"")
	linhaController = append(linhaController, ")")
	linhaController = append(linhaController, "")
	linhaController = append(linhaController, "func Get"+strings.Title(table_name)+"(c *gin.Context) {")
	linhaRoles := ""
	if os.Getenv("security.roles") != "false" {
		linhaController = append(linhaController, "\n\t// Pega as Roles do usuário vindas do middleware")
		linhaController = append(linhaController, "\trolesInterface, exists := c.Get(\"roles\")")
		linhaController = append(linhaController, "\tif !exists {")
		linhaController = append(linhaController, "\t\tc.JSON(400, gin.H{\"message\": \"Usuário não autenticado\"})")
		linhaController = append(linhaController, "\t\treturn")
		linhaController = append(linhaController, "\t}")
		linhaController = append(linhaController, "\troles, _ := rolesInterface.([]string)")
		linhaController = append(linhaController, "\tuserRoles := strings.Split(roles[0], \",\")")

		securityRoles := strings.Split(os.Getenv("security.roles"), ",")
		for _, value := range securityRoles {
			linhaRoles += " !strings.Contains(userRoles, \"" + value + "\") &&"
		}
		linhaRoles = linhaRoles[0 : len(linhaRoles)-3]
		fmt.Println(linhaRoles)

		linhaController = append(linhaController, "\n\t// Roles requeridas do endpoint")
		linhaController = append(linhaController, "\tuserRolesRequired := strings.Split(\"" + os.Getenv("security.roles") + "\", \",\")")

		linhaController = append(linhaController, "\n\t// Busca itens em comum nos slices")
		linhaController = append(linhaController, "\troleSet := make(map[string]struct{})")
		linhaController = append(linhaController, "\tboolRoles := false")

		linhaController = append(linhaController, "\tfor _, item := range userRoles {")
		linhaController = append(linhaController, "\t\troleSet[item] = struct{}{}")
		linhaController = append(linhaController, "\t}")
		linhaController = append(linhaController, "\n\tfor _, item := range userRolesRequired {")
		linhaController = append(linhaController, "\t\tif _, exists := roleSet[item]; exists {")
		linhaController = append(linhaController, "\n\t\t\tboolRoles = true")
		linhaController = append(linhaController, "\t\t}")
		linhaController = append(linhaController, "\t}\n")

		linhaController = append(linhaController, "\tif !boolRoles{")
		linhaController = append(linhaController, "\t\tc.JSON(400, \"Sem direito a acessar a API\")")
		linhaController = append(linhaController, "\t\treturn")
		linhaController = append(linhaController, "\t}")
	}
	linhaController = append(linhaController, "\tstartTime := time.Now()")
	linhaController = append(linhaController, "\tvar "+table_name+" []models."+strings.Title(table_name))
	linhaController = append(linhaController, "")
	if os.Getenv("generator.comments") == "true" {
		linhaController = append(linhaController, "\t// Paginação: valores padrão")
	}
	linhaController = append(linhaController, "\tpage, _ := strconv.Atoi(c.DefaultQuery(\"page\", \""+os.Getenv("param.page")+"\"))")
	linhaController = append(linhaController, "\tlimit, _ := strconv.Atoi(c.DefaultQuery(\"limit\", \""+os.Getenv("param.limit")+"\"))")
	if os.Getenv("validate.limit") == "true" {
		linhaController = append(linhaController, "\tif limit < 0 || limit > 100 {")
		linhaController = append(linhaController, "\t\tlimit = 100")
		linhaController = append(linhaController, "\t}")
	}
	linhaController = append(linhaController, "\toffset := (page - 1) * limit")
	linhaController = append(linhaController, "\tmdbUrl := c.Request.Host + c.Request.URL.Path")
	linhaController = append(linhaController, "")
	if os.Getenv("generator.comments") == "true" {
		linhaController = append(linhaController, "\t// Filtros passados na URL")
	}

	for _, column_name := range fieldsSearch {
		column_text := ""
		for indexW, valueW := range column_name {
			if string(valueW) == "&" {
				column_text = column_name[:indexW]
				break
			}
		}
		linhaController = append(linhaController, "\t"+column_text+" := c.Query(\""+column_text+"\")")
	}

	for index, column_name := range fieldsSearch {
		column_text_field := ""
		for indexW, valueW := range column_name {
			if string(valueW) == "&" {
				column_text_field = column_name[:indexW]
				break
			}
		}
		if index == 0 {
			linhaController = append(linhaController, "")
			if os.Getenv("generator.comments") == "true" {
				linhaController = append(linhaController, "\t//Parâmetros incluídos no log do sistema")
			}
			linhaController = append(linhaController, "\tmdbParameterField := []string{")
		}
		if column_name != "" {
			linhaController = append(linhaController, "\t\t\""+column_name+"\",")
			linhaController = append(linhaController, "\t\tc.Query(\""+column_text_field+"\"),")
		}
		if len(fieldsSearch) == index+1 {
			linhaController = append(linhaController, "\t}")
		}
	}

	dtSTART := ""
	for index, column_name := range paramSearchDate {
		if index == 0 {
			linhaController = append(linhaController, "")
			if os.Getenv("generator.comments") == "true" {
				linhaController = append(linhaController, "\t// Parâmetros passados na URL")
			}
		}
		if column_name != "" {
			linhaController = append(linhaController, "\t"+column_name+" := c.Query(\""+column_name+"\")")
			if index == 0 {
				dtSTART = column_name
			} else {
				linhaController = append(linhaController, "")
				if os.Getenv("generator.comments") == "true" {
					linhaController = append(linhaController, "\t// Validação do período")
				}
				linhaController = append(linhaController, "\tif len("+dtSTART+") <= 0 || len("+column_name+") <= 0 {")
				linhaController = append(linhaController, "\t\t"+dtSTART+" = \"2100-01-01\"")
				linhaController = append(linhaController, "\t\t"+column_name+" = \"2100-01-01\"")
				linhaController = append(linhaController, "\t}")
				if os.Getenv("validate.searchDate") != "false" {
					linhaController = append(linhaController, "\tdifS, _ := time.Parse(\"2006-01-02\", "+dtSTART+")")
					linhaController = append(linhaController, "\tdifE, _ := time.Parse(\"2006-01-02\", "+column_name+")")
					linhaController = append(linhaController, "\tdiff := int(difE.Sub(difS).Hours())")
					charValidate := os.Getenv("validate.searchDate")
					charValidateTime := string(charValidate[len(charValidate)-1])
					linhaController = append(linhaController, "")
					if os.Getenv("generator.comments") == "true" {
						linhaController = append(linhaController, "\t// Efetuando a validação do período")
					}
					if charValidateTime == "D" {
						linhaController = append(linhaController, "\tif (diff / 24) > "+charValidate[0:len(charValidate)-1]+" {")
						linhaController = append(linhaController, "\t\tc.JSON(http.StatusInternalServerError, gin.H{\"error\": \""+os.Getenv("validate.messageDate")+" : "+charValidate+"\"})")
						linhaController = append(linhaController, "\t\treturn")
						linhaController = append(linhaController, "\t}")
					} else if charValidateTime == "M" {
						linhaController = append(linhaController, "\tif (diff / 24 / 30) > "+charValidate[0:len(charValidate)-1]+" {")
						linhaController = append(linhaController, "\t\tc.JSON(http.StatusInternalServerError, gin.H{\"error\": \""+os.Getenv("validate.messageDate")+" : "+charValidate+"\"})")
						linhaController = append(linhaController, "\t\treturn")
						linhaController = append(linhaController, "\t}")
					} else if charValidateTime == "Y" {
						linhaController = append(linhaController, "\tif (diff / 24 / 30 / 12) > "+charValidate[0:len(charValidate)-1]+" {")
						linhaController = append(linhaController, "\t\tc.JSON(http.StatusInternalServerError, gin.H{\"error\": \""+os.Getenv("validate.messageDate")+" : "+charValidate+"\"})")
						linhaController = append(linhaController, "\t\treturn")
						linhaController = append(linhaController, "\t}")
					} else {
						linhaController = append(linhaController, "\tif (diff / 24) > "+charValidate[0:len(charValidate)-1]+" {")
						linhaController = append(linhaController, "\t\tc.JSON(http.StatusInternalServerError, gin.H{\"error\": \""+os.Getenv("validate.messageDate")+" : "+charValidate+"\"})")
						linhaController = append(linhaController, "\t\treturn")
						linhaController = append(linhaController, "\t}")
					}
				}
			}
		}
	}

	linhaController = append(linhaController, "")

	for index, column_name := range paramSearchDate {
		if index == 0 {
			if os.Getenv("generator.comments") == "true" {
				linhaController = append(linhaController, "\t//Parâmetros incluídos no log do sistema")
			}
			linhaController = append(linhaController, "\tmdbParameterDate := []string{")
		}
		if column_name != "" {
			if os.Getenv("generator.comments") == "true" {
				linhaController = append(linhaController, "\t\t\""+column_name+"\",")
			}
			linhaController = append(linhaController, "\t\tc.Query(\""+column_name+"\"),")
		}
		if len(paramSearchDate) == index+1 {
			linhaController = append(linhaController, "\t}")
		}
	}

	linhaController = append(linhaController, "")
	if os.Getenv("generator.comments") == "true" {
		linhaController = append(linhaController, "\t//Montando o modelo e preparando os filtros")
	}
	linhaController = append(linhaController, "\tquery := database.DB.Model(&models."+strings.Title(table_name)+"{})")
	linhaController = append(linhaController, "")
	paramDateTime := os.Getenv("param.dateTime")
	dataType := ""
	for _, column_name := range fieldsSearch {
		for _, v := range vw_columns {
			if column_name == v.Column_name {
				dataType = v.Data_type
				break
			}
		}
		column_text := ""
		column_operator := ""
		for indexW, valueW := range column_name {
			if string(valueW) == "&" {
				column_text = column_name[:indexW]
				column_operator = column_name[indexW+1:]
				break
			}
		}
		linhaController = append(linhaController, "\tif "+column_text+" != \"\" {")
		if column_operator == "%like%" || column_operator == "like" || column_operator == "%%" {
			linhaController = append(linhaController, "\t\tquery = query.Where(\""+column_text+" "+os.Getenv("param.searchLIKE")+" ?\", \"%\"+"+column_text+"+\"%\")")
		} else if column_operator == "like%" || column_operator == "%" {
			linhaController = append(linhaController, "\t\tquery = query.Where(\""+column_text+" "+os.Getenv("param.searchLIKE")+" ?\", "+column_text+"+\"%\")")
		} else if column_operator == "%like" {
			linhaController = append(linhaController, "\t\tquery = query.Where(\""+column_text+" "+os.Getenv("param.searchLIKE")+" ?\", \"%\"+"+column_text+")")
		} else if column_operator == "inS" || column_operator == "ins" {
			linhaController = append(linhaController, "\t\tvar country_nameSplit = strings.Split("+column_text+", \""+","+"\")")
			linhaController = append(linhaController, "\t\tcountry_nameSlice := []string{}")
			linhaController = append(linhaController, "\t\tfor _, value := range country_nameSplit {")
			linhaController = append(linhaController, "\t\t\tcountry_nameSlice = append(country_nameSlice, string(value))")
			linhaController = append(linhaController, "\t\t}")
			linhaController = append(linhaController, "\t\tquery = query.Where(\""+column_text+" IN ?\", country_nameSlice )")
		} else if column_operator == "inI" || column_operator == "ini" {
			linhaController = append(linhaController, "\t\tvar country_nameSplit = strings.Split("+column_text+", \""+","+"\")")
			linhaController = append(linhaController, "\t\tcountry_nameSlice := []int{}")
			linhaController = append(linhaController, "\t\tfor _, value := range country_nameSplit {")
			linhaController = append(linhaController, "\t\t\tcountry_nameSlice = append(country_nameSlice, int(value))")
			linhaController = append(linhaController, "\t\t}")
			linhaController = append(linhaController, "\t\tquery = query.Where(\""+column_text+" IN ?\", country_nameSlice )")
		} else {
			pTime := ""
			if dataType == "date" {
				pTime = " \"" + os.Getenv("param.dateTime") + "\""
			} else {
				pTime = ""
			}
			linhaController = append(linhaController, "\t\tquery = query.Where(\""+column_text+" "+column_operator+" ?\", "+column_text+pTime+")")
		}
		linhaController = append(linhaController, "\t}")
		linhaController = append(linhaController, "")
	}
	linhaDataStart := ""
	linhaDataEnd := ""
	for index, column_name := range paramSearchDate {
		if index == 0 {
			linhaDataStart = column_name
		}
		if index == 1 {
			linhaDataEnd = column_name
		}
	}

	if linhaDataStart != "" && linhaDataEnd != "" {
		paramDateTime = "+\" " + paramDateTime + "\""
		linhaController = append(linhaController, "\tif len("+linhaDataStart+") > 0 && len("+linhaDataEnd+") > 0 {")
		linhaController = append(linhaController, "\t\tquery = query.Where(\""+os.Getenv("fields.searchDate")+" BETWEEN ? AND ?\", "+linhaDataStart+" "+paramDateTime+", "+linhaDataEnd+" "+paramDateTime+")")
		linhaController = append(linhaController, "\t}")
		linhaController = append(linhaController, "")
	} else if linhaDataStart != "" {
		paramDateTime = "+\" " + paramDateTime + "\""
		linhaController = append(linhaController, "\tif len("+linhaDataStart+") > 0  {")
		linhaController = append(linhaController, "\t\tquery = query.Where(\""+os.Getenv("fields.searchDate")+" =  ? \", "+linhaDataStart+" "+paramDateTime+")")
		linhaController = append(linhaController, "\t}")
		linhaController = append(linhaController, "")
	}

	orderBy := os.Getenv("fields.orderBY")
	if orderBy != "" {
		orderBy = ".Order(\"" + os.Getenv("fields.orderBY") + "\")"
	} else {
		orderBy = ""
	}
	if os.Getenv("generator.comments") == "true" {
		linhaController = append(linhaController, "\t// Efetua a consulta no banco de dados")
	}
	linhaController = append(linhaController, "\terr := query.Offset(offset).Limit(limit)"+orderBy+".Find(&"+table_name+").Error")
	linhaController = append(linhaController, "\tendTime := time.Now()")
	linhaController = append(linhaController, "")
	if os.Getenv("generator.comments") == "true" {
		linhaController = append(linhaController, "\t// Log do sistema - OBRIGATÓRIO")
	}
	linhaController = append(linhaController, "\tif err != nil {")
	linhaController = append(linhaController, "\t\terrMd0 := logSystem.WriteLogMongoDB(\""+os.Getenv("logMongoDb.database")+"\", \""+os.Getenv("logMongoDb.collection")+"\", \"Get"+strings.Title(table_name)+"\", \"ronaldo.costa@aviva.com.br\", \""+os.Getenv("logMongoDb.module")+"\", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, \"404\")")
	linhaController = append(linhaController, "\t\tif errMd0 != nil {")
	linhaController = append(linhaController, "\t\t\tlogSystem.WriteLogFile(\":404:"+os.Getenv("logMongoDb.database")+":"+os.Getenv("logMongoDb.collection")+":"+os.Getenv("logMongoDb.module")+":"+"ronaldo.costa@aviva.com.br:\""+"+startTime.String()+\":\"+endTime.String())")
	linhaController = append(linhaController, "\t\t}")
	linhaController = append(linhaController, "\t\tc.JSON(http.StatusInternalServerError, gin.H{\"error\": \""+os.Getenv("validate.query")+" "+os.Getenv("table.singularName")+"\"})")
	linhaController = append(linhaController, "\t\treturn")
	linhaController = append(linhaController, "\t} else {")
	linhaController = append(linhaController, "\t\terrMd1 := logSystem.WriteLogMongoDB(\""+os.Getenv("logMongoDb.database")+"\", \""+os.Getenv("logMongoDb.collection")+"\", \"Get"+strings.Title(table_name)+"\", \"ronaldo.costa@aviva.com.br\", \""+os.Getenv("logMongoDb.module")+"\", startTime, endTime, mdbParameterField, mdbParameterDate, mdbUrl, \"202\")")
	linhaController = append(linhaController, "\t\tif errMd1 != nil {")
	linhaController = append(linhaController, "\t\t\tlogSystem.WriteLogFile(\":202:"+os.Getenv("logMongoDb.database")+":"+os.Getenv("logMongoDb.collection")+":"+os.Getenv("logMongoDb.module")+":"+"ronaldo.costa@aviva.com.br:\""+"+startTime.String()+\":\"+endTime.String())")
	linhaController = append(linhaController, "\t\t}")
	linhaController = append(linhaController, "\t}")
	linhaController = append(linhaController, "\tc.JSON(http.StatusOK, "+table_name+")")
	linhaController = append(linhaController, "}")

	err = escreverTexto(linhaController, os.Getenv("path")+"controllers/"+tabela+"_controller.go")
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	copy("generator/generator.properties", "generator/files/"+table_name+".properties")

	c.JSON(http.StatusOK, vw_columns)

}

func escreverTexto(linhas []string, caminhoDoArquivo string) error {
	arquivo, err := os.Create(caminhoDoArquivo)

	if err != nil {
		return err
	}
	defer arquivo.Close()

	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}
	return escritor.Flush()
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func writeRoutes(route string) {
	pathFile := os.Getenv("path") + "routes/routes.go"
	input, err := ioutil.ReadFile(pathFile)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	textContains := false
	for _, line := range lines {
		if strings.Contains(line, route) {
			textContains = true
		}
	}

	if !textContains {
		for i, line := range lines {
			if strings.Contains(line, "// NÃO RETIRAR ESSA LINHA") {
				lines[i] = "\t\t" + route + "\n\t\t// NÃO RETIRAR ESSA LINHA"
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(pathFile, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
	
}
