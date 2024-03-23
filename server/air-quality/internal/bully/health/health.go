package health

import (
	"fmt"
	"net/http"
	"time"

	"server/air-quality/models"
	"server/air-quality/shared"
	"server/air-quality/internal/bully/election"
)

func HandleHealthOfNode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func CheckHealthOfLeader() {
    for {
        currentLeader := shared.GetLeader()

        if currentLeader != shared.NodeID {
            var leaderNode models.Node
            for _, node := range election.Nodes {
                if node.ID == currentLeader {
                    leaderNode = node
                    break
                }
            }

            if err := checkNodeHealth(leaderNode); err != nil {
                fmt.Println("Leader is not alive. Starting election...")
                election.StartElection()

                time.Sleep(5 * time.Second)

                newLeader := shared.GetLeader()
                fmt.Println("New Leader:", newLeader)
            } else {
                fmt.Println("Leader is alive:", shared.GetLeader())
            }
        }
        time.Sleep(7 * time.Second)
    }
}

func checkNodeHealth(node models.Node) error {
    resp, err := http.Get(fmt.Sprintf("http://%s/bully/health", node.IP))
    if err != nil {
        fmt.Println("Error checking leader health:", err)
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Node %d (%s) responded with status code: %d\n", node.ID, node.IP, resp.StatusCode)
        return fmt.Errorf("leader not healthy")
    }

    return nil
}