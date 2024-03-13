import random
from abc import ABC, abstractmethod
import json
import uuid


class Sensor(ABC):
    @abstractmethod
    def set_client(self):
        pass

    @abstractmethod
    def send(self):
        pass


class AirQualitySensor(Sensor):
    sensor_id = str(uuid.uuid4())

    async def set_client(self, client):
        self.client = client

    async def send(self):
        air_quality_value = random.uniform(0, 100)
        message = {
            "sensor_id": self.sensor_id,
            "value": air_quality_value,
            "unit": "AQI"
        }
        self.client.publish("air_quality", payload=json.dumps(
            message).encode(), qos=1, retain=False)
        print("Published air quality:", air_quality_value)
