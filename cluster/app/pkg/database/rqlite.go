package database

import (
    "bytes"
    "log"
    "net/http"
    "fmt"
)

func SaveToRQLite(data string) error {
    url := "http://localhost:4001/db/execute?queue"
    reqBody := bytes.NewBufferString(data)

    resp, err := http.Post(url, "application/json", reqBody)
    if err != nil {
        return fmt.Errorf("unexpected response status: %s", err)

    }
    defer resp.Body.Close()

    log.Printf("Response: %s", resp.Status)

    if resp.StatusCode != http.StatusOK {
        log.Printf("unexpected response status: %s", resp.Status)
        return fmt.Errorf("unexpected response status: %s", resp.Status)
    }

    return nil
}
