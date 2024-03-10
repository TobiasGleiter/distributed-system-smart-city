import uuid
import random
import json
from abc import ABC, abstractmethod


class Sensor(ABC):
    @abstractmethod
    def set_client(self):
        pass

    @abstractmethod
    def send(self):
        pass


class TemperatureSensor(Sensor):
    sensor_id = str(uuid.uuid4())

    async def set_client(self, client):
        self.client = client

    async def send(self):
        temperature_value = random.uniform(20.0, 30.0)
        message = {
            "sensor_id": self.sensor_id,
            "value": temperature_value,
            "unit": "Celsius"
        }
        await self.client.publish("temperature", json.dumps(message).encode())
