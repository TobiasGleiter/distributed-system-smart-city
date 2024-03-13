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


class VolumeSensor(Sensor):
    sensor_id = str(uuid.uuid4())

    async def set_client(self, client):
        self.client = client

    async def send(self):
        volume_value = random.uniform(60.0, 110.0)
        message = {
            "sensor_id": self.sensor_id,
            "value": volume_value,
            "unit": "VU"
        }
        self.client.publish("temeperature", payload=json.dumps(
            message).encode(), qos=1, retain=False)
        print("Published air quality:", volume_value)
