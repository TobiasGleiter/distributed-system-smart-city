package mongodb

import (
    "context"


    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
    URI        string
    Database   string
    Collection string
}

type Client struct {
    client     *mongo.Client
    collection *mongo.Collection
}

func NewClient(ctx context.Context, config Config) (*Client, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
    if err != nil {
        return nil, err
    }
    collection := client.Database(config.Database).Collection(config.Collection)
    return &Client{client: client, collection: collection}, nil
}

func (c *Client) Close(ctx context.Context) error {
    return c.client.Disconnect(ctx)
}
