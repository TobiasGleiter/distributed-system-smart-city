import asyncio
from mqtt import setup_mqtt_client
from sensors import WaterQualitySensor


async def main():
    try:
        mqtt_client = setup_mqtt_client()

        water_quality_sensor = WaterQualitySensor()
        await water_quality_sensor.set_client(mqtt_client)

        while True:
            await water_quality_sensor.send()
            await asyncio.sleep(5)

    except KeyboardInterrupt:
        print("Keyboard interrupt detected. Exiting...")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == '__main__':
    asyncio.run(main())
