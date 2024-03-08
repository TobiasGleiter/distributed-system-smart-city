import asyncio
import nats
from sensors import VolumeSensor


async def main():
    try:
        nats_client_connection = await nats.connect("demo.nats.io")

        volume_sensor = VolumeSensor()
        await volume_sensor.set_client(nats_client_connection)

        while True:
            await volume_sensor.send()
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
