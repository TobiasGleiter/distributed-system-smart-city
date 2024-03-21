import asyncio
from sensors import WaterQualitySensor


async def main():
    try:
        sensor = WaterQualitySensor()
        sensor.set_send_url("http://localhost:8080/sensor/water_quality/add")

        while True:
            sensor.send()
            await asyncio.sleep(5)

    except KeyboardInterrupt:
        print("Keyboard interrupt detected. Exiting...")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == '__main__':
    asyncio.run(main())
