package airquality

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
    "sync"
    "fmt"

    "server/air-quality/pkg/db"
    //"server/air-quality/pkg/cpu"
    "server/air-quality/shared"

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

func Handler(mc *db.MongoDBClient) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var newSensorData SensorData
        if err := json.NewDecoder(r.Body).Decode(&newSensorData); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        cacheMutex.Lock()
        defer cacheMutex.Unlock()

        if _, exists := cache[newSensorData.SensorID]; exists {
            fmt.Println("Sensor data already exists in cache:", newSensorData.SensorID)
            response := LeaderResponse{IsLeader: shared.IsLeader()}
            sendJSONResponse(w, response)
            return
        }

        if !shared.IsLeader() {
            response := LeaderResponse{IsLeader: false}
            sendJSONResponse(w, response)
            return
        }

        cache[newSensorData.SensorID] = newSensorData
        fmt.Println("Cached", newSensorData.SensorID)

        response := LeaderResponse{IsLeader: true}
        sendJSONResponse(w, response)
    }
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


func SaveCachedToDatabase(mc *db.MongoDBClient) {
    ticker := time.NewTicker(10 * time.Second)
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

                fmt.Println("Saving data to database")
                
                go InsertAirQualityIntoDatabase(mc, dataToSave)
            } else {
                cacheMutex.Unlock()
            }
        }
    }
}

func InsertAirQualityIntoDatabase(mc *db.MongoDBClient, data map[string]SensorData) {
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