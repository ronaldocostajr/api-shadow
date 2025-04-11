package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// err := godotenv.Load("go-api/.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %s", err)
	// }

	host := "host=" + os.Getenv("postgres_host")
	user := "user=" + os.Getenv("postgres_user")
	password := "password=" + os.Getenv("postgres_password")
	dbname := "dbname=" + os.Getenv("postgres_dbname")
	port := "port=" + os.Getenv("postgres_port")

	dsn := host + " " + user + " " + password + " " + dbname + " " + port
	// "host=garrulously-affluent-takin.data-1.use1.tembo.io user=postgres password=Rac134972372715537700 dbname=postgres port=5432 "
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}

	fmt.Println("Conexão com banco de dados bem-sucedida!")
	DB = database
}
