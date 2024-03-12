# Cluster 🎯

## App on Raspberries

NATS Messages from Sensors look like this:

```json
{
  "sensor_id": "sensor_id",
  "value": "value",
  "unit": "the unit"
}
```

## Build the go app

Build the go for Linux:

```bash
env GOOS=linux GOARCH=arm go build -o <target_executable> <source_file>
```
