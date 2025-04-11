package logSystem

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Key   string
	Value string
}

type Connection struct {
	Clients []Client
}

func WriteLogMongoDB(mdbDatabase string, mdbCollection string, mdbEndpoint string, mdbUser string, mdbModule string, mdbStartTime time.Time, mdbEndTime time.Time, mdbParameterField []string, mdbParameterDate []string, mdbUrl string, mdbMessage string) error {
	// Find .evn
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %s", err)
	// }

	// Get value from .env
	MONGO_URI := os.Getenv("mongodb_uri")
	//"mongodb://localhost:27017"
	// Connect to the database.
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Create collection
	databaseMD := ""
	if mdbDatabase != "" {
		databaseMD = mdbDatabase
	} else {
		databaseMD = os.Getenv("mongodb_database")
	}

	collectionMD := ""
	if mdbCollection != "" {
		collectionMD = mdbCollection
	} else {
		collectionMD = os.Getenv("mongodb_database")
	}

	collection := client.Database(databaseMD).Collection(collectionMD)
	if err != nil {
		return err
	}

	// insert a single document into a collection
	// create a bson.D object

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	hosts, _ := os.Hostname()
	registro := bson.D{
		{"endpoint", mdbEndpoint},
		{"httpStatus", mdbMessage},
		{"IP", localAddr.IP.String()},
		{"host", hosts},
		{"url", mdbUrl},
		{"user", mdbUser},
		{"system", mdbCollection},
		{"module", mdbModule},
		{"parameterField", mdbParameterField},
		{"parameterDate", mdbParameterDate},
		{"startTime", mdbStartTime},
		{"endTime", mdbEndTime},
		{"elapsedTimeMinutes", int(mdbEndTime.Sub(mdbStartTime).Minutes())},
		{"createdAt", time.Now()},
	}

	result, err := collection.InsertOne(context.TODO(), registro)

	if err != nil {
		return err
	} else {
		fmt.Println(result.InsertedID)
	}

	return nil
}
