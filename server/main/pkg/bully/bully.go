package bully

import (
    "fmt"
    "time"
	"net/http"
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
            } else {
                fmt.Println(fmt.Sprintf("Leader on port %d is alive", leaderPORT))
            }
        }
    }
}



// func SendMessageToNodes(nodes []models.Node, message string) {
//     for _, node := range nodes {
//         // Send POST request to each node
//         resp, err := http.Post("http://localhost:"+node.Port+"/message", "text/plain", bytes.NewBufferString(message))
//         if err != nil {
//             fmt.Println("Error:", err)
//             continue // Continue to the next node if an error occurs
//         }
//         defer resp.Body.Close()

//         // Check response status
//         if resp.StatusCode != http.StatusOK {
//             fmt.Println("Error:", resp.Status)
//             continue // Continue to the next node if an error occurs
//         }

//         fmt.Println("Message sent to Node", node.Port, "successfully")
//     }
// }

// func HandleMessage(w http.ResponseWriter, r *http.Request) {
//     if r.Method == "POST" {
//         // Read message from request body
//         buf := new(bytes.Buffer)
//         _, err := buf.ReadFrom(r.Body)
//         if err != nil {
//             http.Error(w, "Failed to read request body", http.StatusInternalServerError)
//             return
//         }
//         message := buf.String()

//         // Process the message (e.g., store it in a database, log it)
//         fmt.Println("Received message:", message)

//         // Send a response if needed
//         fmt.Fprintf(w, "Message received successfully")
//     } else {
//         http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//     }
// }
