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
    LeaderID int
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

    fmt.Println("Election message sent successfully to port", node.Port)
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
    if electionMsg.SenderID < NodeID {
        // If sender's ID is lower than local node's ID, respond with "OK"
        fmt.Fprint(w, "OK")
    } else if electionMsg.SenderID > NodeID {
        // If sender's ID is higher than local node's ID, start an election
        StartElection()
        fmt.Fprint(w, "OK")
    } else {
        // If sender's ID is the same as local node's ID, handle as needed
        // For example, you can start an election or respond with "OK"
        // Here, we respond with "OK"
        fmt.Fprint(w, "OK")
    }

    if electionMsg.SenderID > NodeID && len(GetHigherNodes()) == 0 {
        // Declare this node as the leader
        LeaderID = NodeID
        fmt.Println("I am the leader now:", NodeID)

        // Notify all other nodes about the new leader
        NotifyNewLeader()
    }
}

type ElectionMessage struct {
    SenderID   int `json:"sender_id"`
    SenderPort int `json:"sender_port"`
}

type NotificationMessage struct {
    LeaderID int `json:"leader_id"`
}

func NotifyNewLeader() {
    for _, node := range Nodes {
        if node.ID != NodeID { // Skip sending notification to oneself
            // Construct the notification message
            notificationMsg := NotificationMessage{
                LeaderID: LeaderID,
            }

            // Convert the notification message to JSON
            jsonData, err := json.Marshal(notificationMsg)
            if err != nil {
                fmt.Println("Error marshalling notification message:", err)
                return
            }

            // Send the notification message via HTTP POST request
            resp, err := http.Post(fmt.Sprintf("http://localhost:%d/bully/new_leader", node.Port), "application/json", bytes.NewBuffer(jsonData))
            if err != nil {
                fmt.Println("Error sending notification message to node:", err)
                continue
            }
            defer resp.Body.Close()

            // Check the response status
            if resp.StatusCode != http.StatusOK {
                fmt.Println("Error: unexpected response status from node:", resp.Status)
                continue
            }

            fmt.Println("Notification message sent to node", node.ID)
        }
    }
}

func HandleNewLeaderNotification(w http.ResponseWriter, r *http.Request) {
    // Decode the incoming JSON notification
    var notificationMsg NotificationMessage
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&notificationMsg); err != nil {
        fmt.Println("Error decoding new leader notification:", err)
        http.Error(w, "Failed to decode new leader notification", http.StatusBadRequest)
        return
    }

    // Update the local LeaderID variable with the new leader's ID
    LeaderID = notificationMsg.LeaderID

    // Respond with a success message
    fmt.Fprint(w, "Leader updated successfully")
}

