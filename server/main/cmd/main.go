package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"bytes"

	"server/main/config"
	"server/main/pkg/mongodb"
	"server/main/pkg/sensor"
	//"server/main/models"
	//"server/main/pkg/bully"
)

type Message struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Sender    string    `json:"sender"`
	Timestamp time.Time `json:"timestamp"`
}

type NodeInfo struct {
	ID   string `json:"ID"`
	Port string `json:"port"`
}

const (
	ElectionMsg = "ELECTION"
	AnswerMsg   = "ANSWER"
	CoordinatorMsg = "COORDINATOR"
)


var (
	coordinatorID  string
	nodeID         string
	nodes          []string
	mutex          sync.Mutex
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
	//http.HandleFunc("/message", bully.HandleMessage)


	go func() {
		http.HandleFunc("/election", handleElection)
		http.HandleFunc("/answer", handleAnswer)
		http.HandleFunc("/coordinator", handleCoordinator)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	select {}
}

func handleElection(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Received election message from", msg.Sender)

	if msg.ID > nodeID {
		// Respond with an answer message
		sendMessage(msg.Sender, AnswerMsg)
	} else {
		// Initiate own election
		go startElection()
	}
}

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Received answer message from", msg.Sender)
}

func handleCoordinator(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Received coordinator message from", msg.Sender)
	coordinatorID = msg.Sender
}

func sendMessage(target, msgType string) {
	msg := Message{
		ID:        nodeID,
		Type:      msgType,
		Sender:    nodeID,
		Timestamp: time.Now(),
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling message:", err)
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/%s", target, msgType), "application/json", bytes.NewBuffer(jsonMsg))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	defer resp.Body.Close()
}

func startElection() {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Starting election...")
	highestID := nodeID
	for _, n := range nodes {
		if n > highestID {
			highestID = n
		}
	}

	if highestID == nodeID {
		// Become the coordinator
		coordinatorID = nodeID
		fmt.Println("I am the coordinator")
		broadcastCoordinator()
	} else {
		// Send election messages to higher IDs
		for _, n := range nodes {
			if n > nodeID {
				sendMessage(n, ElectionMsg)
			}
		}
	}
}

func broadcastCoordinator() {
	for _, n := range nodes {
		if n != nodeID {
			sendMessage(n, CoordinatorMsg)
		}
	}
}




