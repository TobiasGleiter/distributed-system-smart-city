import asyncio
import nats
from sensors import WaterQualitySensor


async def main():
    try:
        nats_client_connection = await nats.connect("demo.nats.io")

        water_quality_sensor = WaterQualitySensor()
        await water_quality_sensor.set_client(nats_client_connection)

        while True:
            await water_quality_sensor.send()
            await asyncio.sleep(5)

    except KeyboardInterrupt:
        print("Keyboard interrupt detected. Exiting...")
    except Exception as e:
        print(f"An error occurred: {e}")
    finally:
        if nats_client_connection.is_connected:
            await nats_client_connection.close()


if __name__ == '__main__':
    asyncio.run(main())
