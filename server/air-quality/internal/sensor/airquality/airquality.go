package airquality

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "server/air-quality/shared"
    "server/air-quality/pkg/db"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type SensorData struct {
    SensorID string  `json:"sensor_id"`
    Value    float64 `json:"value"`
    Unit     string  `json:"unit"`
}

type LeaderResponse struct {
    IsLeader bool `json:"isLeader"`
    LeaderID int  `json:"leaderID,omitempty"`
}

func PostAirQualityHandler(mc *db.MongoDBClient) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var newSensorData SensorData
        if err := json.NewDecoder(r.Body).Decode(&newSensorData); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        if !shared.IsLeader() {
            response := LeaderResponse{
                IsLeader: false,
            }
            if shared.Leader == 0 {
                http.Error(w, "Leader not yet elected", http.StatusServiceUnavailable)
                return
            }
            response.LeaderID = shared.Leader

            w.Header().Set("Content-Type", "application/json")
            if err := json.NewEncoder(w).Encode(response); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        collection := mc.Database("sensor").Collection("air_quality")
        filter := bson.M{"sensorid": newSensorData.SensorID}
        update := bson.M{"$set": newSensorData}

        _, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        response := LeaderResponse{IsLeader: true}
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}
