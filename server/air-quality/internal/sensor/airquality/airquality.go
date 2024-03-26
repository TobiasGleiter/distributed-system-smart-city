package airquality

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
    "sync"
    "fmt"

    "server/air-quality/pkg/db"
    "server/air-quality/pkg/cpu"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/mongo/options"
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

var (
    cache      = make(map[string]SensorData)
    cacheMutex sync.Mutex
    savingData bool
)

func AirQualityHandler(mc *db.MongoDBClient) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var newSensorData SensorData
        if err := json.NewDecoder(r.Body).Decode(&newSensorData); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        cacheMutex.Lock()
        cache[newSensorData.SensorID] = newSensorData
        cacheMutex.Unlock()

        fmt.Println("Cached")

        response := LeaderResponse{IsLeader: true}
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(response); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func SaveCachedDataToDB(mc *db.MongoDBClient, cpuStats *cpu.Stats) {
	go cpuStats.GetCPUUsage()

    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()
    for {
        <-ticker.C
        if !savingData {
            cacheMutex.Lock()
            if len(cache) > 0 {
                savingData = true
                dataToSave := make(map[string]SensorData, len(cache))
                for sensorID, data := range cache {
                    dataToSave[sensorID] = data
                    delete(cache, sensorID)
                }
                cacheMutex.Unlock()

                
                go saveToDatabase(mc, dataToSave)
            } else {
                cacheMutex.Unlock()
            }
        }
    }
}

func saveToDatabase(mc *db.MongoDBClient, data map[string]SensorData) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    collection := mc.Database("sensor").Collection("air_quality")

    var operations []mongo.WriteModel
    for sensorID, data := range data {
        filter := bson.M{"sensorid": sensorID}
        update := bson.M{"$set": data}
        model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
        operations = append(operations, model)
    }

    _, err := collection.BulkWrite(ctx, operations)
    if err != nil {
        // Handle error
    }

    savingData = false
}