package config

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var broker = "mqtt.eclipseprojects.io:1883"

func ConnectToMqttServer()(MQTT.Client, error) {
    opts := MQTT.NewClientOptions().AddBroker(broker)
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
    }
    return client, nil
}