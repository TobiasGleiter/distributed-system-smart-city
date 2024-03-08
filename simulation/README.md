# Sensor Simulation Programs 🆙

## Sensors

There are four simulation programms:

1. Air-Quality
2. Temperature
3. Volume
4. Water-Quality

Run `docker-compose up -d --build` to run the simulation in docker.
Note: Air-Quality has currently 3 replicas.

## NATS Topics

The sensors send values on the following topics:

- `water_quality`
- `air_quality`
- `temperature`
- `volume`
