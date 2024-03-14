package models

type AirQualityMessage struct {
    SensorID string  `json:"sensor_id"`
    Value    float64 `json:"value"`
    Unit     string  `json:"unit"`
}
