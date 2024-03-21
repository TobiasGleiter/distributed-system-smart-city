package main

import (
	"log"
	"net/http"

	"server/main/config"
	"server/main/pkg/mongodb"
	"server/main/pkg/sensor"
)

func main() {
    config, err := config.LoadConfig("config.json")
    if err != nil {
        log.Fatal(err)
    }

	mongodb.ConnectToMongoDB(config.MongoURI)

	client := mongodb.GetClient()
	sensor.Initialize(client)

	http.HandleFunc("/sensor/air_quality/add", sensor.UpdateSensorData)

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}







