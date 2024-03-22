package main
 import (
	"log"
	"net/http"
	"fmt"
	"flag"

	"server/main/config"
	"server/main/pkg/mongodb"
	"server/main/pkg/sensor"
	//"server/main/models"
	"server/main/pkg/bully"
)

var (
	NodeID int
	LeaderID int
	LeaderPort int
)

func main() {
    configFile := flag.String("config", "config.json", "Path to the configuration file")
    flag.Parse()

    config, err := config.LoadConfig(*configFile)
    if err != nil {
        log.Fatal(err)
    }

	mongodb.ConnectToMongoDB(config.MongoURI)

	client := mongodb.GetClient()
	sensor.Initialize(client)

	NodeID = config.ID
	LeaderID = config.Leader
	LeaderPort = 8080

	if NodeID == LeaderID {
		fmt.Println("I am the Leader! Thats me:", NodeID)
	} else {
		fmt.Println(fmt.Sprintf("I am worker %d and the leader is %d", NodeID, LeaderID))

		bully.NodeID = config.ID
		bully.NodePort = config.Port
		bully.Nodes = config.Nodes
		go bully.CheckHeartbeatFromLeader(LeaderPort)
	}

	http.HandleFunc("/sensor/air_quality/add", sensor.UpdateSensorData)
	http.HandleFunc("/bully/heartbeat", bully.HandleHeartbeatAsLeader)
	http.HandleFunc("/bully/election/message", bully.HandleElectionMessage)

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
	}()

	select {}
}




