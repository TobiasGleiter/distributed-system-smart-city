import asyncio
import nats


async def message_handler(msg):
    subject = msg.subject
    data = msg.data.decode()
    print(f"Received a message on '{subject}': {data}")


async def subscribe():
    nc = await nats.connect("demo.nats.io")

    # Subscribe to a subject
    await nc.subscribe("air_quality", cb=message_handler)

    # Keep the connection open
    await asyncio.sleep(3600)

if __name__ == '__main__':
    asyncio.run(subscribe())
