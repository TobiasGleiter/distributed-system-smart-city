# Distributed Systems - Smart City 🏙️

Simple smart city server application and simulation of sensors for a distributed system.

## API-Endpoints of the Server Application

Test Nodes with CURL:

`curl -X POST -H "Content-Type: application/json" -d '{"sensor_id": "sensor123", "value": 25, "unit": "AQI"}' http://localhost:8080/sensor/air_quality
`

Returns either `{"isLeader":false}` or `{"isLeader":true}`
