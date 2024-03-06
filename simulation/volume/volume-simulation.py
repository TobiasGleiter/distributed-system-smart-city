from sensors import Volume
from mqtt import setup_mqtt_client

mqtt_client = setup_mqtt_client()


while True:
    volume_sensor = Volume()
    volume_sensor.set_client(mqtt_client)
    volume_sensor.send()
