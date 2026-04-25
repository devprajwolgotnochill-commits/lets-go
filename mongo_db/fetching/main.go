package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Connecting to:", os.Getenv("DB_URL"))

	uri := os.Getenv("MONGO_DB")

	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"usage-examples/#environment-variable")
	}

	//connecting
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//fetching
	coll := client.Database("local").Collection("oplog.rs")
	title := "oplog.rs"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{}).
		Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}

	if err != nil {
		panic(err)
	}

	//turning into json
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	//printing
	fmt.Printf("%s\n", jsonData)
}
