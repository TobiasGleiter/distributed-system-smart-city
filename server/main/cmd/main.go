package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
    "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

type SensorData struct {
	SensorID string  `json:"sensor_id"`
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
}

// MongoDB connection URI
const (
	databaseName    = "sensor"
	airQualityColl  = "air_quality"
	waterQualityColl = "water_quality"
	volumeColl      = "volume"
	temperatureColl = "temperature"
)

type Config struct {
    MongoURI string `json:"mongo_uri"`
}

var client *mongo.Client

func main() {
    config, err := loadConfig("config.json")
    if err != nil {
        log.Fatal(err)
    }

	connectToMongoDB(config)

	http.HandleFunc("/sensor/air_quality/add", addAirQualityData)
	http.HandleFunc("/sensor/water_quality/add", addWaterQualityData)
	http.HandleFunc("/sensor/volume/add", addVolumeData)
	http.HandleFunc("/sensor/temperature/add", addTemperatureData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectToMongoDB(config Config) {
    clientOptions := options.Client().ApplyURI(config.MongoURI)

    // Connect to MongoDB
    var err error  // Change client declaration to avoid shadowing
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")
}


func addAirQualityData(w http.ResponseWriter, r *http.Request) {
	updateSensorData(w, r, airQualityColl)
}

func addWaterQualityData(w http.ResponseWriter, r *http.Request) {
	updateSensorData(w, r, waterQualityColl)
}

func addVolumeData(w http.ResponseWriter, r *http.Request) {
	updateSensorData(w, r, volumeColl)
}

func addTemperatureData(w http.ResponseWriter, r *http.Request) {
	updateSensorData(w, r, temperatureColl)
}

func updateSensorData(w http.ResponseWriter, r *http.Request, collectionName string) {
	var newData SensorData
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database(databaseName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    filter := bson.M{"sensorid": newData.SensorID}
    update := bson.M{"$set": newData}

    _, err = collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
    if err != nil {
        log.Println("Error updating/inserting message into MongoDB:", err)
        return
    }

	fmt.Println("Updated sensor data in", collectionName, "collection:", newData.SensorID)
	w.WriteHeader(http.StatusOK)
}


func loadConfig(filename string) (Config, error) {
    var config Config
    configFile, err := os.Open(filename)
    if err != nil {
        return config, err
    }
    defer configFile.Close()

    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    return config, err
}