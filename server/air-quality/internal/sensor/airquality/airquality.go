package airquality

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
    "sync"
    "fmt"
    "bytes"

    "server/air-quality/pkg/db"
    //"server/air-quality/pkg/cpu"
    "server/air-quality/shared"
    "server/air-quality/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/mongo/options"
)

type SensorData struct {
    SensorID string  `json:"sensor_id"`
    Value    float64 `json:"value"`
    Unit     string  `json:"unit"`
    Timestamp int64 `json:"timestamp"`
}

type LeaderResponse struct {
    IsLeader bool `json:"isLeader"`
    LeaderID int  `json:"leaderID,omitempty"`
}

type WorkerResponse struct {
    Acknowledged bool `json:acknowledged`
}

var (
    cache      = make(map[string]SensorData)
    cacheMutex sync.Mutex
    savingData bool
    currentIdx = 0
)

func RoundRobinBalancer() models.Node {
    nodes := shared.GetNodes()
    thisNode := models.Node{ID: shared.NodeID, IP: shared.NodeIP} // Assuming 1 is the ID of the current node
    allNodes := append(nodes, thisNode)

    node := allNodes[currentIdx]
    currentIdx = (currentIdx + 1) % len(allNodes)
	return node
}

func DistributeSensorData() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        if !shared.IsLeader() {
            response := LeaderResponse{IsLeader: false}
            sendJSONResponse(w, response)
            return
        }

        var newSensorData SensorData
        if err := json.NewDecoder(r.Body).Decode(&newSensorData); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        jsonData, err := json.Marshal(newSensorData)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        fmt.Println("Sensor Data:", newSensorData)

        // Forward the request to the next node using round-robin
        nextNode := RoundRobinBalancer()
        fmt.Println(nextNode)
        // Assuming models.Node has an endpoint field
        resp, err := http.Post(fmt.Sprintf("http://%s/sensor/air_quality/worker", nextNode.IP), "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            fmt.Println("I save it to my cache...")
            err = saveSensorToCache(newSensorData)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            response := LeaderResponse{IsLeader: true}
            sendJSONResponse(w, response)
            return
        }

        // Check the response from the worker node
        var workerResp WorkerResponse
        if err := json.NewDecoder(resp.Body).Decode(&workerResp); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        if !workerResp.Acknowledged {
            http.Error(w, "Worker did not acknowledge the request", http.StatusInternalServerError)
            return
        }

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



func SaveCacheToDatabase(mc *db.MongoDBClient) {
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

                
                go InsertAirQualityIntoDatabase(mc, dataToSave)
            } else {
                cacheMutex.Unlock()
            }
        }
    }
}


func saveSensorToCache(data SensorData) error {
    fmt.Println("Saving to cache")
    cacheMutex.Lock()
    defer cacheMutex.Unlock()

    data.Timestamp = time.Now().Unix()

    if _, exists := cache[data.SensorID]; exists {
        fmt.Println("Sensor data already exists in cache:", data.SensorID)
        return nil
    }

    cache[data.SensorID] = data
    fmt.Println("Cached", data.SensorID)
    return nil
}


func SensorCacheHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Cache Handler")
        var newSensorData SensorData
        if err := json.NewDecoder(r.Body).Decode(&newSensorData); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        saveSensorToCache(newSensorData)

        response := WorkerResponse{Acknowledged: true}
        sendJSONResponse(w, response)
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

    fmt.Println("Saved to database")

    savingData = false
}