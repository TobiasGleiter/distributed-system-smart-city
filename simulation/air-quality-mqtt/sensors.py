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


class AirQualitySensor(Sensor):
    sensor_id = str(uuid.uuid4())

    def set_send_url(self, send_url):
        self.send_url = send_url

    def send(self):
        air_quality_value = random.uniform(0, 100)
        message = {
            "sensor_id": self.sensor_id,
            "value": air_quality_value,
            "unit": "AQI"
        }

        headers = {'Content-Type': 'application/json'}
        response = requests.post(
            self.send_url, data=json.dumps(message), headers=headers)
        if response.status_code == 200:
            response_data = json.loads(response.text)
            if response_data["isLeader"] == False:
                leaderID = response_data["leaderID"]
                print(f"I am not the leader, the Leader is: {leaderID}")
            else:
                print(
                    f"Published air quality to {self.send_url}: {air_quality_value}")
        else:
            print(
                f"Failed to publish air quality to {self.send_url}: {response.status_code}")
