import random
import time
from abc import ABC, abstractmethod


class Sensor(ABC):
    @abstractmethod
    def set_client(self):
        pass

    @abstractmethod
    def send(self):
        pass


class AirQuality(Sensor):
    def set_client(self, client):
        self.client = client

    def send(self):
        air_quality_value = random.uniform(0, 100)
        self.client.publish("air_quality", payload=str(
            air_quality_value), qos=1, retain=False)
        print("Published air quality:", air_quality_value)
        time.sleep(5)
