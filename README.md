# Distributed Systems - Smart City 🏙️

Simple smart city server application and simulation of sensors for a distributed system.

## API-Endpoints of the Database-Write-Application

Important from the outside:

- `/sensor/air_quality`
- `/sensor/air_quality/worker`

Internal use:

- `/bully/health`
- `/bully/election`

## Testing with Curl

Send to the leader of the cluster, all else will return isLeader: false. The Leader is the Node with the highest ID. Can be found in the server/air-quality/cmd json config files.

Test Nodes with CURL (here localhost:8080):

`curl -X POST -H "Content-Type: application/json" -d '{"sensor_id": "sensor123", "value": 25, "unit": "AQI"}' http://localhost:8080/sensor/air_quality
`

Returns either `{"isLeader":false}` or `{"isLeader":true}`
