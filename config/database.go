package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Missing MONGO_URI environment variable")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDb ping failed ", err)
	}

	DB = client

	log.Println("Connected to MongoDB")

}

func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database client is not initialized. Ensure ConnectDB() is called before using GetCollection.")
	}

	return DB.Database("crypto").Collection(collectionName)
}
