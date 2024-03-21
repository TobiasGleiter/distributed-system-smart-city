package sensor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

var (
	client       *mongo.Client
	databaseName string
)

const (
	database = "sensor"
	mongoCollection  = "air_quality"
)


type SensorData struct {
	SensorID string  `json:"sensor_id"`
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
}

func Initialize(mongoClient *mongo.Client) {
	client = mongoClient
}

func UpdateSensorData(w http.ResponseWriter, r *http.Request) {
	var newSensorData SensorData
	err := json.NewDecoder(r.Body).Decode(&newSensorData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database(database).Collection(mongoCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    filter := bson.M{"sensorid": newSensorData.SensorID}
    update := bson.M{"$set": newSensorData}

    _, err = collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
    if err != nil {
        log.Println("Error updating/inserting message into MongoDB:", err)
        return
    }

	fmt.Println("Updated sensor data in", mongoCollection, "collection:", newSensorData.SensorID)
	w.WriteHeader(http.StatusOK)
}