import asyncio
from sensors import AirQualitySensor
import requests.exceptions


async def main():
    try:
        sensor = AirQualitySensor()
        cluster_ips = ["http://localhost:8080/sensor/air_quality",
                       "http://localhost:8081/sensor/air_quality",
                       "http://localhost:8082/sensor/air_quality",
                       "http://localhost:8083/sensor/air_quality",
                       "http://localhost:8084/sensor/air_quality"]

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
