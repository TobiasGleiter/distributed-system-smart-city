import asyncio
from sensors import VolumeSensor


async def main():
    try:
        sensor = VolumeSensor()
        sensor.set_send_url("http://localhost:8080/sensor/volume/add")

        while True:
            sensor.send()
            await asyncio.sleep(5)

    except KeyboardInterrupt:
        print("Keyboard interrupt detected. Exiting...")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == '__main__':
    asyncio.run(main())
