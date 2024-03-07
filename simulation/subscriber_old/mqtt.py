import paho.mqtt.client as mqtt


def on_connect(client, userdata, flags, reason_code, properties):
    print(f"Connected with result code {reason_code}")
    # Subscribing in on_connect() means that if we lose the connection and
    # reconnect then subscriptions will be renewed.
    client.subscribe("air_quality")
    client.subscribe("temperature")
    client.subscribe("water_quality")
    client.subscribe("volume")


def on_message(client, userdata, msg):
    print(msg.topic, msg.payload.decode())


def setup_mqtt_client():
    mqttc = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
    mqttc.on_connect = on_connect
    mqttc.on_message = on_message
    mqttc.connect("mqtt.eclipseprojects.io", 1883, 60)
    mqttc.loop_start()
    return mqttc
