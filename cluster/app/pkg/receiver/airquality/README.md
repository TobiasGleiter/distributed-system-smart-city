# Database test commands

## From terminal with CURL

Create the table for air_quality

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d '[
    "CREATE TABLE air_quality (sensor_id UUID PRIMARY KEY, value FLOAT NOT NULL, unit CHAR NOT NULL)"
]'
```

Read from database

```bash
curl -G 'localhost:4001/db/query' --data-urlencode 'q=SELECT * FROM air_quality'
```

Insert

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d '[
    "INSERT INTO air_quality (sensor_id, value, unit) VALUES(\"1\", 76.23520235921171, \"celsius\") ON CONFLICT(sensor_id) DO UPDATE SET value = EXCLUDED.value, unit = EXCLUDED.unit"
]'
```

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d '[
    "INSERT INTO air_quality (sensor_id, value, unit) VALUES(\"9348c0fa-cc25-4dfe-a3e1-803925c990dd\", 39.1, \"AQI\") ON CONFLICT(sensor_id) DO UPDATE SET value = EXCLUDED.value, unit = EXCLUDED.unit"
]'
```

## From the RQLite CLI:

Insert and update

```sql
INSERT INTO air_quality (sensor_id, value, unit) VALUES("1", 2.3, "celsius") ON CONFLICT(sensor_id) DO UPDATE SET value = EXCLUDED.value, unit = EXCLUDED.unit
```
