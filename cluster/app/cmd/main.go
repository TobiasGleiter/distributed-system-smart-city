package main

import (
    "log"

    "smartcity/server/pkg/receiver/airquality"
    "smartcity/server/config"
)


func main() {
    nc, err := config.ConnectToNatsClient()
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

	airQualityReceiver := &airquality.AirQualityReceiver{}
    airQualityReceiver.SetClient(nc)
    airQualityReceiver.SaveIncomingAirQualityToDatabase()



	select {}
}
