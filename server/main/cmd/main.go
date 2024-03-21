package main

import (
	"log"
	"net/http"
	"bytes"
	"fmt"
	"time"

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
	http.HandleFunc("/message", handleMessage)

	go func() {
        log.Fatal(http.ListenAndServe(":"+config.Port, nil))
    }()

	// Start sending messages to other nodes every 5 seconds
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            sendMessageToNodes(config.Nodes, "Hello, Server!")
        }
    }
}


func sendMessageToNodes(nodes []config.Node, message string) {
    for _, node := range nodes {
        // Send POST request to each node
        resp, err := http.Post("http://localhost:"+node.Port+"/message", "text/plain", bytes.NewBufferString(message))
        if err != nil {
            fmt.Println("Error:", err)
            continue // Continue to the next node if an error occurs
        }
        defer resp.Body.Close()

        // Check response status
        if resp.StatusCode != http.StatusOK {
            fmt.Println("Error:", resp.Status)
            continue // Continue to the next node if an error occurs
        }

        fmt.Println("Message sent to Node", node.Port, "successfully")
    }
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // Read message from request body
        buf := new(bytes.Buffer)
        _, err := buf.ReadFrom(r.Body)
        if err != nil {
            http.Error(w, "Failed to read request body", http.StatusInternalServerError)
            return
        }
        message := buf.String()

        // Process the message (e.g., store it in a database, log it)
        fmt.Println("Received message:", message)

        // Send a response if needed
        fmt.Fprintf(w, "Message received successfully")
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}






