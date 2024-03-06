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


class Temperature(Sensor):
    def set_client(self, client):
        self.client = client

    def send(self):
        temperature_value = random.uniform(20.0, 30.0)
        self.client.publish("temperature", payload=str(
            temperature_value), qos=1, retain=False)
        print("Published temperature:", temperature_value)
        time.sleep(5)
