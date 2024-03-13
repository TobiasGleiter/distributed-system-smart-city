package airquality

import (
    "log"
    "encoding/json"
    
    MQTT "github.com/eclipse/paho.mqtt.golang"
)

type AirQualityMessage struct {
    SensorID string `json:"sensor_id"`
    Value    float64 `json:"value"`
    Unit     string `json:"unit"`
}

type AirQualityReceiver struct {
    client MQTT.Client
}

func (a *AirQualityReceiver) SetClient(client MQTT.Client) {
    a.client = client
}

func (a *AirQualityReceiver) SaveIncomingAirQualityToDatabase() {
    token := a.client.Subscribe("air_quality", 0, func(client MQTT.Client, msg MQTT.Message) {
        var airQualityMsg AirQualityMessage
        if err := json.Unmarshal(msg.Payload(), &airQualityMsg); err != nil {
            log.Println("Error parsing JSON message:", err)
            return
        }

        log.Println("Message from Sensor:", airQualityMsg.SensorID)

    })
    if token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }
}