from sensors import Temperature
from mqtt import setup_mqtt_client

mqtt_client = setup_mqtt_client()

while True:
    temperature_sensor = Temperature()
    temperature_sensor.set_client(mqtt_client)
    temperature_sensor.send()
