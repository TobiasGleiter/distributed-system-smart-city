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


class WaterQualitySensor(Sensor):
    async def set_client(self, client):
        self.client = client

    async def send(self):
        water_quality_value = random.uniform(0, 100)
        await self.client.publish("water_quality", str(water_quality_value).encode())
