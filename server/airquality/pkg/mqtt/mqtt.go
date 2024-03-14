package mqtt

import (
    "context"
    "fmt"
    "log"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
    Broker   string
    ClientID string
    Topic    string
}

type MessageHandler struct{}

func (mh *MessageHandler) ConnectAndListen(ctx context.Context, config Config) {
    opts := MQTT.NewClientOptions().AddBroker(config.Broker).SetClientID(config.ClientID)
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    if token := client.Subscribe(config.Topic, 0, mh.messageHandler); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    fmt.Printf("Connected to MQTT broker. Listening to topic: %s\n", config.Topic)

    select {}
}

func (mh *MessageHandler) messageHandler(client MQTT.Client, msg MQTT.Message) {
    var airQualityMsg AirQualityMessage
    err := json.Unmarshal(msg.Payload(), &airQualityMsg)
    if err != nil {
        log.Println("Error unmarshaling message:", err)
        return
    }

    filter := bson.M{"sensor_id": airQualityMsg.SensorID}
    update := bson.M{"$set": airQualityMsg}

    ctx := context.Background()
    _, err = mh.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
    if err != nil {
        log.Println("Error updating/inserting message into MongoDB:", err)
        return
    }

    fmt.Printf("Received message: %+v\n", airQualityMsg)
}
