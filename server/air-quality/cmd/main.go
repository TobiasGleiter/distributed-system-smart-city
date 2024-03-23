package main

import (
	"fmt"
	"net/http"
	"flag"
	"log"
	"time"

	"server/air-quality/config"
	"server/air-quality/models"
	"server/air-quality/shared"
	"server/air-quality/internal/bully/health"
	"server/air-quality/internal/bully/election"
)


func main() {
	configFile := flag.String("config", "config.json", "Path to the configuration file")
    flag.Parse()

    cfg, err := config.LoadConfig(*configFile)
    if err != nil {
        log.Fatal(err)
    }

	var nodes []models.Node
	for _, node := range cfg.Nodes {
		nodes = append(nodes, models.Node{ID: node.ID, IP: node.IP})
	}
	shared.NodeID = cfg.ID
	election.Nodes = nodes
	shared.SetLeader(100)


	go DoTasks()
	go health.CheckHealthOfLeader()


	http.HandleFunc("/bully/health", health.HandleHealthOfNode)
	http.HandleFunc("/bully/election", election.HandleElectionRequest)

	fmt.Println("Server listening on ip", cfg.IP)
	if err := http.ListenAndServe(cfg.IP, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

func DoTasks() {
	for {
		if shared.IsLeader() {
			fmt.Println("I am doing tasks...")
		}
		time.Sleep(2 * time.Second)
	}
}




