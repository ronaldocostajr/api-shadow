package database

import (
	"fmt"
	"log"
	"os"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func ConnectORA() {
	host := os.Getenv("DB_ORA_HOST")
	//port := os.Getenv("DB_ORA_PORT")
	service := os.Getenv("DB_ORA_SERVICE")
	user := os.Getenv("DB_ORA_USER")
	pass := os.Getenv("DB_ORA_PASS")

	options := map[string]string{
	}

	url := oracle.BuildUrl(host, 1521, service, user, pass, options)

	fmt.Println("Valor de URL: ", url)
	dialector := oracle.New(oracle.Config{
		DSN:	url,
		IgnoreCase: 	false,
		NamingCaseSensitive: true,
		VarcharSizeIsCharLength: true,
		RowNumberAliasForOracle11: "ROW_NUM",
	})

	database, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt: false,
		CreateBatchSize: 50,
	})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}

	fmt.Println("Conexão com banco de dados bem-sucedida!")
	DB = database
}
