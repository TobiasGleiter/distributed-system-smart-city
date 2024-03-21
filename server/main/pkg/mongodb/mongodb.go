package mongodb

import (
	"fmt"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetClient() *mongo.Client {
    return client
}

func ConnectToMongoDB(mongo_uri string) {
    clientOptions := options.Client().ApplyURI(mongo_uri)

    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")
}
