package election

import (
	"fmt"
	"encoding/json"
	"bytes"
	"net/http"

	"server/air-quality/models"
	"server/air-quality/shared"
)

type ElectionMessage struct {
	SenderID int `json:"ID"`
}

var (
	Nodes []models.Node
)

func StartElection() {
	fmt.Println("Start Election")

	// nodes := getHigherNodes(NodeID, Nodes)
	// fmt.Printf("Nodes: %+v\n", nodes)

	// I am the Leader now
	shared.SetLeader(shared.NodeID)
	fmt.Printf("Election: I am the Leader now: %d\n", shared.Leader)

	// Send message to all higher nodes
	// I am the leader
	for _, node := range Nodes {
		sendElectionMessage(node)
	}
}

func HandleElectionRequest(w http.ResponseWriter, r *http.Request) {
	var electionMsg ElectionMessage
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&electionMsg); err != nil {
        fmt.Println("Error decoding election message:", err)
        http.Error(w, "Failed to decode election message", http.StatusBadRequest)
        return
    }
	
    // Handle the election message
    if electionMsg.SenderID > shared.NodeID {
		shared.SetLeader(electionMsg.SenderID)
		fmt.Println(fmt.Sprintf("The leader is now %d and I am (%d) a worker.", shared.Leader, shared.NodeID))
		fmt.Fprint(w, "OK")

	} else if electionMsg.SenderID < shared.NodeID {
		fmt.Println(fmt.Sprintf("Handle Election Request: The leader is now %d and I am (%d) a worker.", shared.Leader, shared.NodeID))
		StartElection()
	}

}

	// // use this for the election process
	// nodes = getHigherNodes(cfg.ID, nodes)
	// fmt.Printf("Nodes: %+v\n", nodes)

func sendElectionMessage(node models.Node) {
	// Implement logic to send an election message to a node
	fmt.Println(fmt.Sprintf("Send Election Message to Node %d on port %s", node.ID, node.IP))

	// Construct the election message
	electionMsg := ElectionMessage{
		SenderID: shared.NodeID,
	}

	// Convert the election message to JSON
	jsonData, err := json.Marshal(electionMsg)
	if err != nil {
		fmt.Println("Error marshalling election message:", err)
		return
	}

	// Send the election message via HTTP POST request
	resp, err := http.Post(fmt.Sprintf("http://%s/bully/election", node.IP), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending election message:", err)
		shared.SetLeader(shared.NodeID)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: unexpected response status:", resp.Status)
		return
	}

	fmt.Println("Election message sent successfully to port", node.IP)
}