import random
import json
import requests
from abc import ABC, abstractmethod
import uuid


class Sensor(ABC):
    @abstractmethod
    def set_send_url(self):
        pass

    @abstractmethod
    def send(self):
        pass


class WaterQualitySensor(Sensor):
    sensor_id = str(uuid.uuid4())

    def set_send_url(self, send_url):
        self.send_url = send_url

    def send(self):
        value = random.uniform(25, 30)
        message = {
            "sensor_id": self.sensor_id,
            "value": value,
            "unit": "auq"
        }

        headers = {'Content-Type': 'application/json'}
        response = requests.post(
            self.send_url, data=json.dumps(message), headers=headers)
        if response.status_code == 200:
            print("Published air quality:", value)
        else:
            print("Failed to publish air quality:", response.status_code)
