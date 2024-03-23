package airquality

import (
	"fmt"
	"net/http"
	"encoding/json"

	"server/air-quality/shared"
)

type LeaderResponse struct {
	IsLeader bool `json:"isLeader"`
    LeaderID int  `json:"leaderID,omitempty"`
}

func HandleAirQualityRequest(w http.ResponseWriter, r *http.Request) {
	response := LeaderResponse{}

	if shared.IsLeader() {
        response.IsLeader = true
    } else {
        if shared.Leader == 0 {
            response.IsLeader = false
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Fprintf(w, "Leader not yet elected")
            return
        }
        response.IsLeader = false
        response.LeaderID = shared.Leader
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}