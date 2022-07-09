package database

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database = "mongodb-go"
)

func Connection() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client, nil
}

func GetCollection(collection string) *mongo.Collection {
	client, _ := Connection()
	return client.Database(database).Collection(collection)
}
