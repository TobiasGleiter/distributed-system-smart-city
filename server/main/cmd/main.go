package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type SensorData struct {
    SensorID string  `json:"sensor_id"`
    Value    float64 `json:"value"`
    Unit     string  `json:"unit"`
}

var sensorData []SensorData

func main() {
    http.HandleFunc("/data", getData)
    http.HandleFunc("/data/add", addData)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sensorData)
}

func addData(w http.ResponseWriter, r *http.Request) {
    var newData SensorData
    err := json.NewDecoder(r.Body).Decode(&newData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    sensorData = append(sensorData, newData)
    w.WriteHeader(http.StatusCreated)
}
