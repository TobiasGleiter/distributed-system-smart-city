package bully

import (
    "fmt"
    "time"
	"net/http"
    "encoding/json"
    "bytes"

    "server/main/models"
)

var (
    NodeID int
    NodePort int
    Nodes []models.Node
)

func HandleHeartbeatAsLeader(w http.ResponseWriter, r *http.Request) { 
    w.WriteHeader(http.StatusOK)
}

func CheckHeartbeatFromLeader(leaderPORT int) {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            _, err := http.Get(fmt.Sprintf("http://localhost:%d/bully/heartbeat", leaderPORT))
            if err != nil {
                fmt.Println("Leader is not alive start election!")
                StartElection()

            } else {
                fmt.Println(fmt.Sprintf("Leader on port %d is alive", leaderPORT))
            }
        }
    }
}

func StartElection() {
    // Get higher-ranked nodes
    higherNodes := GetHigherNodes()

    // Send election messages to higher-ranked nodes
    for _, node := range higherNodes {
        fmt.Println(fmt.Sprintf("Send Election Message", node.ID))
        sendElectionMessage(node)
    }
}

func GetHigherNodes() []models.Node {
    higherNodes := make([]models.Node, 0)
    for _, node := range Nodes {
        if node.ID > NodeID {
            higherNodes = append(higherNodes, node)
        }
    }
    return higherNodes
}

func sendElectionMessage(node models.Node) {
    // Implement logic to send an election message to a node
    fmt.Println(fmt.Sprintf("Send Election Message to Node %d on port %d", node.ID, node.Port))

    // Construct the election message
    electionMsg := ElectionMessage{
        SenderID: NodeID,
        SenderPort: NodePort, // Assuming you have NodePort defined somewhere
    }

    // Convert the election message to JSON
    jsonData, err := json.Marshal(electionMsg)
    if err != nil {
        fmt.Println("Error marshalling election message:", err)
        return
    }

    // Send the election message via HTTP POST request
    resp, err := http.Post(fmt.Sprintf("http://localhost:%d/bully/election/message", node.Port), "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error sending election message:", err)
        return
    }
    defer resp.Body.Close()

    // Check the response status
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: unexpected response status:", resp.Status)
        return
    }

    fmt.Println("Election message sent successfully")
}

func HandleElectionMessage(w http.ResponseWriter, r *http.Request) {
    var electionMsg ElectionMessage
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&electionMsg); err != nil {
        fmt.Println("Error decoding election message:", err)
        http.Error(w, "Failed to decode election message", http.StatusBadRequest)
        return
    }

    // Handle the election message
    // Check if the sender's ID is higher than the local node's ID
    // If yes, respond to the election message by sending an "OK" response
    // Otherwise, ignore the election message

    // Example response
    fmt.Fprint(w, "OK")
}

type ElectionMessage struct {
    SenderID   int `json:"sender_id"`
    SenderPort int `json:"sender_port"`
}