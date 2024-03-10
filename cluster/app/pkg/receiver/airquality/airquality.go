package airquality

import (
    "bytes"
    "log"
    "net/http"
    "fmt"
    "encoding/json"

    "github.com/nats-io/nats.go"
)

type AirQualityMessage struct {
    SensorID string `json:"sensor_id"`
    Value    float64 `json:"value"`
    Unit     string `json:"unit"`
}

type AirQualityReceiver struct {
    client       *nats.Conn
}

func (a *AirQualityReceiver) SetClient(nc *nats.Conn) {
    a.client = nc
}

func (a *AirQualityReceiver) SaveIncomingAirQualityToDatabase() {
    _, err := a.client.QueueSubscribe("air_quality", "airquality_nodes", func(msg *nats.Msg) {
        var airQualityMsg AirQualityMessage
        if err := json.Unmarshal(msg.Data, &airQualityMsg); err != nil {
            log.Println("Error parsing JSON message:", err)
            return
        }

        
        query := fmt.Sprintf(`[
            ["INSERT INTO air_quality (sensor_id, value, unit) VALUES(\"%s\", %.1f, \"%s\")"]
        ]`, airQualityMsg.SensorID, airQualityMsg.Value, airQualityMsg.Unit)

        // Send the SQL statement to RQLite
        if err := SaveToRQLite([]byte(query)); err != nil {
            log.Println("Error sending SQL statement to RQLite:", err)
            return
        }

        log.Printf("Worker: Saved air quality message to database: %s\n", string(msg.Data))
    })
    if err != nil {
        log.Fatal(err)
    }
}


func SaveToRQLite(data []byte) error {
    url := "http://localhost:4001/db/execute?timeout=2m"
    reqBody := bytes.NewBuffer(data)

    // Send POST request to RQLite server
    resp, err := http.Post(url, "application/json", reqBody)
    if err != nil {
        return fmt.Errorf("unexpected response status: %s", err)

    }
    defer resp.Body.Close()

    log.Printf("Response: %s", resp.Status)

    // Check response status
    if resp.StatusCode != http.StatusOK {
        log.Printf("unexpected response status: %s", resp.Status)
        return fmt.Errorf("unexpected response status: %s", resp.Status)
    }

    return nil
}
