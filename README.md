# Distributed Systems - Smart City 🏙️

**Transforming Cities with Smart Technology**

- **Enhancing Infrastructure:** Transportation, water supply, energy management, and waste disposal are optimized through smart technology.
- **Improving Quality of Life:** Intelligent information systems promote community engagement and elevate living standards.
- **Sustainable Development:** Smart city initiatives prioritize economic, environmental, and social sustainability.

## About this Project 📋

This project utilizes NATS, rqlite, and Python to establish a distributed smart city infrastructure with sensors and clients. In our case, it runs on three Raspberry Pi devices with 8GB of memory.

- **NATS:** Acts as a message broker, handling requests on each device.
- **Rqlite:** Manages database access, essential in a distributed system where various applications access the database.
- **Python:** Chosen for its simplicity, making entry into the codebase easy.

Besides that we using Docker to scale the number of sensors (simulated)
