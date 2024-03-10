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


class WaterQualitySensor(Sensor):
    sensor_id = str(uuid.uuid4())

    async def set_client(self, client):
        self.client = client

    async def send(self):
        water_quality_value = random.uniform(0, 100)
        message = {
            "sensor_id": self.sensor_id,
            "value": water_quality_value,
            "unit": "VU"
        }
        await self.client.publish("water_quality", json.dumps(message).encode())
