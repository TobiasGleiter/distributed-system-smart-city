# Distributed Systems - Smart City 🏙️

Simple smart city server application and simulation of sensors for a distributed system.

## API-Endpoints of the Server Application

The server application writes the data to the databse (mongodb)

| Endpoint                    | Method | Data                                                                 |
| --------------------------- | ------ | -------------------------------------------------------------------- |
| `/sensor/air_quality/add`   | POST   | `{ "sensor_id": "<sensor_id>", "value": <value>, "unit": "<unit>" }` |
| `/sensor/water_quality/add` | POST   | `{ "sensor_id": "<sensor_id>", "value": <value>, "unit": "<unit>" }` |
| `/sensor/volume/add`        | POST   | `{ "sensor_id": "<sensor_id>", "value": <value>, "unit": "<unit>" }` |
| `/sensor/temperature/add`   | POST   | `{ "sensor_id": "<sensor_id>", "value": <value>, "unit": "<unit>" }` |
