from sensors import WaterQuality
from mqtt import setup_mqtt_client

mqtt_client = setup_mqtt_client()

while True:
    water_quality_sensor = WaterQuality()
    water_quality_sensor.set_client(mqtt_client)
    water_quality_sensor.send()
