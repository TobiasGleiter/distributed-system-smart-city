package main

import (
	"fmt"
	"net/http"
	"flag"
	"log"
	"context"

	"server/air-quality/config"
	"server/air-quality/models"
	"server/air-quality/shared"
	"server/air-quality/pkg/db"
	"server/air-quality/pkg/cpu"
	"server/air-quality/internal/bully/health"
	"server/air-quality/internal/bully/election"
	"server/air-quality/internal/sensor/airquality"
)


func main() {
	configFile := flag.String("config", "config.json", "Path to the configuration file")
    flag.Parse()

    cfg, err := config.LoadConfig(*configFile)
    if err != nil {
        log.Fatal(err)
    }

	ctx := context.Background()
	mongoURI := cfg.MongoURI
	mongoClient := db.NewMongoDBClient()
	if err := mongoClient.Connect(ctx, mongoURI); err != nil {
		fmt.Println("MongoDB error:", err)
	}
	defer mongoClient.Disconnect(ctx)

	var nodes []models.Node
	for _, node := range cfg.Nodes {
		nodes = append(nodes, models.Node{ID: node.ID, IP: node.IP})
	}
	shared.NodeID = cfg.ID
	shared.Nodes = nodes
	shared.SetLeader(100)


	go health.CheckHealthOfLeader()


	http.HandleFunc("/bully/health", health.HandleHealthOfNode)
	http.HandleFunc("/bully/election", election.HandleElectionRequest)

	http.HandleFunc("/sensor/air_quality", airquality.Handler(mongoClient))


	cpuStats := &cpu.Stats{}
	go cpuStats.GetCPUUsage()
	go airquality.SaveCachedToDatabase(mongoClient)


	fmt.Println("Server listening on ip", cfg.IP)
	if err := http.ListenAndServe(cfg.IP, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}




