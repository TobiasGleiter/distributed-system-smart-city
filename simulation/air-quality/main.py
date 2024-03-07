import asyncio
import nats
from sensors import AirQuality


async def main():
    try:
        nc = await nats.connect("demo.nats.io")

        air_quality_sensor = AirQuality()
        await air_quality_sensor.set_client(nc)

        while True:
            await air_quality_sensor.send()
            await asyncio.sleep(5)

    except KeyboardInterrupt:
        print("Keyboard interrupt detected. Exiting...")
    except Exception as e:
        print(f"An error occurred: {e}")
    finally:
        if nc.is_connected:
            await nc.close()


if __name__ == '__main__':
    asyncio.run(main())
