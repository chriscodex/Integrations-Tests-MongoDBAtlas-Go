package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	usr      = "christian"
	pwd      = "123"
	host     = "localhost"
	port     = 27017
	database = "mongodb-go"
)

func Connection() (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	rd, _ := readpref.New(1)
	err = client.Ping(ctx, rd)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetCollection(collection string) mongo.Collection {
	client, _ := Connection()
	return *client.Database(database).Collection(collection)
}
