package bully

import (
	"sync"
    "bytes"
    "fmt"
    "net/http"
)

type Node struct {
    ID   int
    Port string
    Alive bool
}

type Bully struct {
    Nodes    map[int]*Node
    CoordinatorID int
    mutex    sync.Mutex
}

func NewBully(nodes []Node) *Bully {
    bully := &Bully{
        Nodes: make(map[int]*Node),
    }

    for _, node := range nodes {
        bully.Nodes[node.ID] = &node
    }

    return bully
}

func (b *Bully) StartElection() {
    b.mutex.Lock()
    defer b.mutex.Unlock()

    highestID := b.findHighestID()
    for id, node := range b.Nodes {
        if id > b.CoordinatorID && node.Alive {
			// Send election message to nodes with higher IDs
			url := fmt.Sprintf("http://localhost:%s/message", node.Port)
			message := "ELECTION" // Example message, you can customize it
			_, err := http.Post(url, "text/plain", bytes.NewBufferString(message))
			if err != nil {
				fmt.Printf("Failed to send election message to node %d: %v\n", id, err)
				continue
			}
			fmt.Printf("Election message sent to node %d successfully\n", id)
        }
    }

    // If the node is the highest ID, it becomes the coordinator
    if highestID == b.CoordinatorID {
        b.becomeCoordinator()
    }
}

func (b *Bully) findHighestID() int {
    maxID := 0
    for id := range b.Nodes {
        if id > maxID {
            maxID = id
        }
    }
    return maxID
}

func (b *Bully) becomeCoordinator() {
    // Update coordinator ID
    b.CoordinatorID = b.findHighestID()
    // Broadcast coordinator message to other nodes
    // Here you should implement the logic to send coordinator message to other nodes
}

func (b *Bully) HandleCoordinatorMessage(senderID int) {
    // Update coordinator ID if received coordinator ID is higher
    if senderID > b.CoordinatorID {
        b.CoordinatorID = senderID
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
