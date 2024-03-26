package db

import (
    "context"
    
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
    Connect(ctx context.Context, uri string) error
    Disconnect(ctx context.Context) error
    Database(name string) *mongo.Database
}

type MongoDBClient struct {
    client *mongo.Client
}

func NewMongoDBClient() *MongoDBClient {
    return &MongoDBClient{}
}

func (m *MongoDBClient) Connect(ctx context.Context, uri string) error {
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return err
    }
    m.client = client
    return nil
}

func (m *MongoDBClient) Disconnect(ctx context.Context) error {
    if m.client != nil {
        err := m.client.Disconnect(ctx)
        if err != nil {
            return err
        }
    }
    return nil
}

func (m *MongoDBClient) Database(name string) *mongo.Database {
    return m.client.Database(name)
}