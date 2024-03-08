package airquality

import (
    "log"

    "github.com/nats-io/nats.go"
)

type AirQualityMessage struct {
    SensorID string
    Value    float64
    Unit     string
}

type AirQualityReceiver struct {
    client       *nats.Conn
}

func (a *AirQualityReceiver) SetClient(nc *nats.Conn) {
    a.client = nc
}

func (a *AirQualityReceiver) SaveIncomingAirQualityToDatabase() {
    _, err := a.client.Subscribe("air_quality", func(msg *nats.Msg) {
        log.Printf("Saving air quality message to database: %s\n", string(msg.Data))
    })
    if err != nil {
        log.Fatal(err)
    }
}