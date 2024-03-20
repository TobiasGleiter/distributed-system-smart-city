package main

import (
    "log"

    "smartcity/server/pkg/receiver/airquality"
    "smartcity/server/config"
)


func main() {
    mqttClient, err := config.ConnectToMqttServer()
    if err != nil {
        log.Fatal(err)
    }
    defer mqttClient.Disconnect(250)

	airQualityReceiver := &airquality.AirQualityReceiver{}
    airQualityReceiver.SetClient(mqttClient)


    airQualityReceiver.SaveIncomingAirQualityToDatabase()
   
    log.Println("Main function is running...")

    select {}
}
