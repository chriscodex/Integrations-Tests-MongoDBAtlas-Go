package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	// Add the uri of your cluster of MongoDBAtlas
	// https://www.mongodb.com/docs/drivers/go/current/quick-start/#connect-to-your-cluster
	uri = ""
	// Add the name of your database
	database = ""
)

func Connection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	rd, _ := readpref.New(1)
	err = client.Ping(ctx, rd)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCollection(collection string) *mongo.Collection {
	client, _ := Connection()
	return client.Database(database).Collection(collection)
}
