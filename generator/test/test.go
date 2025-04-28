package main

import (
	"fmt"
	middleware "go-api/middlewares"
)

//	_ "github.com/sijms/go-ora/v2"

var localDB = map[string]string{
	"service":  "XE",
	"username": "system",
	"server":   "127.0.0.1",
	"port":     "1521",
	"password": "adm",
}

//func main() {
// fields := "nm_cliente#=,dt_aniversario#>=,fl_ativo#like, nu_cep#ilike"
// fields2 := strings.Split(fields, ",")
// for _, value := range fields2 {
// 	for index, valueW := range value {
// 		if string(valueW) == "#" {
// 			fmt.Println(value[:index])
// 			fmt.Println(value[index+1:])
// 			break
// 		}
// 	}
// }
// texto := "BRASIL,ARGENTINA,CHILE"
// texto = "\"" + strings.Replace(texto, ",", "\",\"", -1) + "\""
// fmt.Println(texto)

// CONEXÃO ORACLE
// connectString := "oracle://" + localDB["username"] + ":" + localDB["password"] + "@" + localDB["server"] + ":" + localDB["port"] + "/" + localDB["service"]
// db, err := sql.Open("oracle", connectString)
// if err != nil {
// 	log.Fatal(err)
// }
// defer db.Close()
// err = db.Ping()
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println("Conexão bem sucedida!")

// rows, err := db.Query("select ip_address from c##api_system.TB_API_IP_ADDRESS")
// if err != nil {
// 	log.Fatal(err)
// }
// defer db.Close()

// fmt.Println(rows)

// CONEXÃO MONGODB

// func main() {
// 	connectMOngoDB()
// }

// func connectMOngoDB() {
// 	// Find .evn
// 	// err := godotenv.Load(".env")
// 	// if err != nil {
// 	// 	log.Fatalf("Error loading .env file: %s", err)
// 	// }

// 	// Get value from .env
// 	MONGO_URI := "mongodb://localhost:27017"
// 	// Connect to the database.
// 	clientOptions := options.Client().ApplyURI(MONGO_URI)
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create collection
// 	collection := client.Database("SHADOW").Collection("shadow_cx")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// insert a single document into a collection
// 	// create a bson.D object
// 	registro := bson.D{
// 		{"endpoint", "https://api-aviva.com.br/GetTb_usuario"},
// 		{"user", "ronaldo.costa@aviva.com.br"},
// 		{"system", "tesouraria"},
// 		{"module", "financiero"},
// 		{"startTime", time.Now()},
// 		{"endTime", time.Now()},
// 		{"createdAt", time.Now()},
// 	}

// 	result, err := collection.InsertOne(context.TODO(), registro)

// 	if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println(result.InsertedID)
// 	}

// type Client struct {
// 	Key   string
// 	Value string
// }

// type Connection struct {
// 	Clients []Client
// }

// func main() {
// 	startDate := "2025-01-01"
// 	endDate := "2025-01-05"
// 	startDATE, _ := time.Parse("2006-01-02", startDate)
// 	endDATE, _ := time.Parse("2006-01-02", endDate)

// 	diff := endDATE.Sub(startDATE)

// 	fmt.Println(int(diff.Hours() / 24))
// }

// func main() {
// 	err := godotenv.Load("generator/generator.properties")
// 	if err != nil {
// 		panic(err)
// 	}
// 	charValidate := os.Getenv("validate.searchDate")
// 	fmt.Println(charValidate)
// 	fmt.Println(string(charValidate[len(charValidate)-1]))
// }

// func main() {
// 	difS, _ := time.Parse("2006-01-02", "2025-01-01")
// 	difE, _ := time.Parse("2006-01-02", "2027-01-01")

// 	diff := int(difE.Sub(difS).Hours())
// 	fmt.Println(diff / 24 / 30 / 12)

// }

// func main() {
// 	err := godotenv.Load("generator/generator.properties")
// 	if err != nil {
// 		panic(err)
// 	}
// 	validate := os.Getenv("validate.searchDate")

// 	fmt.Println(validate)
// 	fmt.Println(string(validate[0 : len(validate)-1]))
// }

// func main() {
// 	err := godotenv.Load("generator/generator.properties")
// 	if err != nil {
// 		panic(err)
// 	}
// 	validate := if os.Getenv("generator.comments") == "true" {

// 	if strconv.ParseBool(validate) {
// 		fmt.Println(validate)
// 	} else {
// 		fmt.Println(string(validate[0 : len(validate)-1]))
// 	}
// }

// func main() {
// 	conteudo, _ := ioutil.ReadFile("speed.txt")
// 	linhas := strings.Split(string(conteudo), "|")
// 	for _, v := range linhas {
// 		fmt.Println(v)
// 		fmt.Println(len(v))
// 		if v == "10.0.0.1" {
// 			fmt.Println(v) // preciso mostra a linha toda do registro C100
// 		}
// 	}
// }

// func main() {
// 	var literalLines []string

// 	file, err := os.Open("speed.txt")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	// guarda cada linha em indice diferente do slice
// 	for scanner.Scan() {
// 		literalLines = append(literalLines, scanner.Text())
// 	}

// 	// busca linha
// 	for _, line := range literalLines {
// 		if strings.Contains(line, "// NÃO RETIRAR ESSA LINHA") {
// 			fmt.Println(line)
// 			output := strings.Join(lines, "\n")
// 			err = ioutil.WriteFile("myfile", []byte(output), 0644)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 		}
// 	}
// }

// func main() {
// 	err := godotenv.Load("generator/generator.properties")
// 	if err != nil {
// 		panic(err)
// 	}

// 	pathFile := os.Getenv("path") + "routes/routes.go"
// 	fmt.Println(pathFile, "...........................")
// 	input, err := ioutil.ReadFile(pathFile)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	lines := strings.Split(string(input), "\n")

// 	textContains := false
// 	for _, line := range lines {
// 		if strings.Contains(line, "routes.Tb_cepRoutess(api)") {
// 			textContains = true
// 		}
// 	}

// 	if !textContains {
// 		for i, line := range lines {
// 			if strings.Contains(line, "// NÃO RETIRAR ESSA LINHA") {
// 				fmt.Println("entrei")
// 				lines[i] = "routes.Tb_cepRoutess(api)\n// NÃO RETIRAR ESSA LINHA 2"
// 			}
// 		}
// 		output := strings.Join(lines, "\n")
// 		err = ioutil.WriteFile(pathFile, []byte(output), 0644)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 	}
// }

// func main() {

// 	pathFile := "D:/PROJETOS_DIVERSOS/projetos GO/go-api/go-api/generator/test/feriado.txt"
// 	input, err := ioutil.ReadFile(pathFile)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	lines := strings.Split(string(input), "\n")

// 	dateHoliday := false
// 	for _, line := range lines {
// 		if strings.Contains(line, "01/01/2026") { //time.Now().Format("02/01/2006")) {
// 			dateHoliday = true
// 			break
// 		}
// 	}
// 	fmt.Println(dateHoliday)
// }

// func main() {
// 	ctx := context.Background()

// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "redis-14966.c245.us-east-1-3.ec2.redns.redis-cloud.com:14966",
// 		Password: "Hkm9nejhkHHLp5zUxdUc2hEy3IuIrjHy", // no password set
// 		DB:       0,                                  // use default DB
// 	})

// 	err := client.Set(ctx, "key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := client.Get(ctx, "key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	fmt.Println("REDIS", middleware.RedisRead("key"))

// 	val2, err := client.Get(ctx, "key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist

// }

func main() {
	retRedisInsert := middleware.RedisInsert("teste", "meu valor 123456")
	fmt.Println(retRedisInsert)

	retRedisRead := middleware.RedisRead("teste")
	fmt.Println(retRedisRead)

	// retRedisDelete := middleware.RedisDelete("teste")
	// fmt.Println(retRedisDelete)

	retRedisRead1 := middleware.RedisRead("teste")
	fmt.Println(retRedisRead1)

}
