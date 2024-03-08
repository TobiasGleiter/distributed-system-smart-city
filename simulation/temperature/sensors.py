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


class TemperatureSensor(Sensor):
    async def set_client(self, client):
        self.client = client

    async def send(self):
        temperature_value = random.uniform(20.0, 30.0)
        await self.client.publish("temperature", str(temperature_value).encode())
