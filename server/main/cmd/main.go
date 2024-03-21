package main

import (
	"log"
	"net/http"

	"server/main/config"
	"server/main/pkg/mongodb"
	"server/main/pkg/sensor"
)

const (
	databaseName    = "sensor"
	airQualityColl  = "air_quality"
)

func main() {
    config, err := config.LoadConfig("config.json")
    if err != nil {
        log.Fatal(err)
    }

	mongodb.ConnectToMongoDB(config.MongoURI)
	client := mongodb.GetClient()
	sensor.Initialize(client, databaseName)

	http.HandleFunc("/sensor/air_quality/add", addAirQualityData)

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}


func addAirQualityData(w http.ResponseWriter, r *http.Request) {
	sensor.UpdateSensorData(w, r, airQualityColl)
}




