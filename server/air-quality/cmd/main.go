package main

import (
	"fmt"
	"net/http"
	"flag"
	"log"

	"server/air-quality/config"
	"server/air-quality/models"
	"server/air-quality/internal/bully/health"
)


func main() {
	configFile := flag.String("config", "config.json", "Path to the configuration file")
    flag.Parse()

    cfg, err := config.LoadConfig(*configFile)
    if err != nil {
        log.Fatal(err)
    }


	var leader = cfg.Leader
	// Start health checking routine for each node
	go func() {
		for _, node := range cfg.Nodes {
			if leader == node.ID {
				go health.Check(models.Config{ID: node.ID, IP: node.IP})
			}	
		}
	}()


	// Convert nodess to []models.Node
	var nodes []models.Node
	for _, node := range cfg.Nodes {
		nodes = append(nodes, models.Node{ID: node.ID, IP: node.IP})
	}

	// Now you have nodes in the format of []models.Node
	// Assuming getHigherNodes returns a slice of models.Node
	//nodes = getHigherNodes(cfg.ID, nodes)

	fmt.Printf("Nodes: %+v\n", nodes)

	http.HandleFunc("/bully/health", health.HandleHealthOfNode)
	http.HandleFunc("/bully/election", handleElectionRequest)

	fmt.Println("Server listening on ip", cfg.IP)
	if err := http.ListenAndServe(cfg.IP, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}



func handleElectionRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func start() {

}

func getHigherNodes(thisNodeId int, nodes []models.Node) []models.Node {
    higherNodes := make([]models.Node, 0)
    for _, node := range nodes {
        if node.ID > thisNodeId {
            higherNodes = append(higherNodes, node)
        }
    }
    return higherNodes
}




