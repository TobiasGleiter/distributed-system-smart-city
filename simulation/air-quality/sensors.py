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


class AirQualitySensor(Sensor):
    async def set_client(self, client):
        self.client = client

    async def send(self):
        air_quality_value = random.uniform(0, 100)
        await self.client.publish("air_quality", str(air_quality_value).encode())
