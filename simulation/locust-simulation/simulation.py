import random
from locust import HttpUser, task, between
import uuid
import json


class MyUser(HttpUser):
    # Adjust this according to your desired wait time
    wait_time = between(1, 3)

    @task
    def send_message(self):
        try:
            sensor_id = str(uuid.uuid4())
            air_quality_value = random.uniform(0, 100)
            message = {
                "sensor_id": sensor_id,
                "value": air_quality_value,
                "unit": "AQI"
            }
            data = json.dumps(message)
            response = self.client.post(
                "/sensor/air_quality", json={"your": "data"})
            if response.status_code != 200:
                response.failure("Failed to send message")
        except Exception as e:
            response.failure(f"Exception: {str(e)}")
