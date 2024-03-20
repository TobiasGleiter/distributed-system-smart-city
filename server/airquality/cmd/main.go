package main

import (
    "context"
    "fmt"
    "log"
    "time"
	"encoding/json"

    MQTT "github.com/eclipse/paho.mqtt.golang"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

	"airquality/models"
)

type MQTTConfig struct {
    Broker   string
    ClientID string
    Topic    string
}

type MongoDBConfig struct {
    URI         string
    Database    string
    Collection  string
}

type MessageHandler struct {
    mongoClient *mongo.Client
    collection  *mongo.Collection
}

func (mh *MessageHandler) ConnectAndListen(mqttConfig MQTTConfig, mongoConfig MongoDBConfig) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.URI))
    if err != nil {
        log.Fatal(err)
    }
    defer mongoClient.Disconnect(ctx)

    mh.mongoClient = mongoClient
    mh.collection = mongoClient.Database(mongoConfig.Database).Collection(mongoConfig.Collection)

    opts := MQTT.NewClientOptions().AddBroker(mqttConfig.Broker).SetClientID(mqttConfig.ClientID)
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    if token := client.Subscribe(mqttConfig.Topic, 0, mh.messageHandler); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    fmt.Printf("Connected to MQTT broker. Listening to topic: %s\n", mqttConfig.Topic)

    select {}
}

func (mh *MessageHandler) messageHandler(client MQTT.Client, msg MQTT.Message) {
    var airQualityMsg models.AirQualityMessage
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

func main() {
    mqttConfig := MQTTConfig{
        Broker:   "mqtt.eclipseprojects.io:1883",
        ClientID: "mqtt-client",
        Topic:    "air_quality",
    }

    mongoConfig := MongoDBConfig{
        URI:        "mongodb+srv://test_user2:EOfApjntJgGosIJ6@smartcity.4okvjzf.mongodb.net/?retryWrites=true&w=majority&appName=SmartCity",
        Database:   "sensors",
        Collection: "air_quality",
    }

    mh := &MessageHandler{}

    mh.ConnectAndListen(mqttConfig, mongoConfig)
}