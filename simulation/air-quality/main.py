import asyncio
from sensors import AirQualitySensor
import requests.exceptions
import json


async def main():
    try:
        sensor = AirQualitySensor()

        with open('config.json') as f:
            config = json.load(f)
            cluster_ips = config["cluster_ips"]

        while True:
            for cluster_ip in cluster_ips:
                try:
                    sensor.set_send_url(cluster_ip)
                    sensor.send()
                except requests.exceptions.RequestException as e:
                    print(
                        f"Continuing with the next endpoint.")
            await asyncio.sleep(5)

    except KeyboardInterrupt:
        print("Keyboard interrupt detected. Exiting...")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == '__main__':
    asyncio.run(main())
