from sensors import AirQuality
from mqtt import setup_mqtt_client

mqtt_client = setup_mqtt_client()


while True:
    air_quality_sensor = AirQuality()
    air_quality_sensor.set_client(mqtt_client)
    air_quality_sensor.send()
