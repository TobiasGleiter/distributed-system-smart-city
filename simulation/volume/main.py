from mqtt import setup_mqtt_client
import asyncio
from sensors import VolumeSensor


async def main():
    try:
        mqtt_client = setup_mqtt_client()

        volume_sensor = VolumeSensor()
        await volume_sensor.set_client(mqtt_client)

        while True:
            await volume_sensor.send()
            await asyncio.sleep(5)

    except KeyboardInterrupt:
        print("Keyboard interrupt detected. Exiting...")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == '__main__':
    asyncio.run(main())
