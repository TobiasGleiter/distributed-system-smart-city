package health

import (
	"fmt"
	"net/http"
	"time"

	"server/air-quality/models"
)

func HandleHealthOfNode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func Check(node models.Config) {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            resp, err := http.Get(fmt.Sprintf("http://%s/bully/health", node.IP))
            if err != nil {
                fmt.Printf("Error healt check Node %d (%s)\n", node.ID, node.IP)
                continue
            }
            defer resp.Body.Close()

            if resp.StatusCode != http.StatusOK {
                fmt.Printf("Node %d (%s) responded with status code: %d\n", node.ID, node.IP, resp.StatusCode)
            } else {
                fmt.Printf("Node %d is healthy\n", node.ID)
            }
        }
    }
}