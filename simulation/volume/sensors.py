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


class Volume(Sensor):
    def set_client(self, client):
        self.client = client

    def send(self):
        volume_value = random.uniform(60.0, 110.0)
        self.client.publish("volume", payload=str(
            volume_value), qos=1, retain=False)
        print("Published volume:", volume_value)
        time.sleep(5)
