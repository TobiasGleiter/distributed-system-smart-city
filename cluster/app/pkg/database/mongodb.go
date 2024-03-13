package database

import (
    "smartcity/server/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var sensorsCollection *mongo.Collection = config.GetCollection(config.DB, "sensors")