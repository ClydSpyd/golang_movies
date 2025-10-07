package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("[WARNING] unable to load .env file")
	}

	MongoDb := os.Getenv("MONGO_URI")

	if MongoDb == "" {
		log.Fatal("[CRITICAL] MONGO_URI environment variable not set")
	}

	log.Println("[INFO] MongoDB URI:", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[INFO] Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("[WARNING] unable to load .env file")
	}
	databaseName := os.Getenv("DATABASE_NAME")
	log.Println("[INFO] Database Name:", databaseName)
	collection := Client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		log.Println("[CRITICAL] Collection not found:", collectionName)
		return nil
	}
	return collection
}
