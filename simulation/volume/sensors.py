import asyncio
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


class VolumeSensor(Sensor):
    async def set_client(self, client):
        self.client = client

    async def send(self):
        volume_value = random.uniform(60.0, 110.0)
        await self.client.publish("volume", str(volume_value).encode())
