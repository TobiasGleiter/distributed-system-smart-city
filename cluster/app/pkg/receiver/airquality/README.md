Create the table for air_quality

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d '[
    "CREATE TABLE air_quality (id INTEGER NOT NULL PRIMARY KEY, sensor_id CHAR NOT NULL, value FLOAT NOT NULL, unit CHAR NOT NULL)"
]'
```

Read from database

```bash
curl -G 'localhost:4001/db/query' --data-urlencode 'q=SELECT * FROM air_quality'
```

Insert

```bash
curl -XPOST 'localhost:4001/db/execute?pretty&timings' -H "Content-Type: application/json" -d '[
    "INSERT INTO air_quality (sensor_id, value, unit) VALUES(\"1\", 2.3, \"celsius\")"
]'
```
