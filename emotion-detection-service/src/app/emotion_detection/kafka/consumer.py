import asyncio
from pykafka import KafkaClient as PyKafkaClient


class KafkaConsumerClient:
    @staticmethod
    async def consume_kafka_messages():
        client = PyKafkaClient("localhost:9092")
        topic = client.topics["image-upload"]

        consumer = topic.get_simple_consumer(
            auto_commit_enable=True, reset_offset_on_start=False
        )
        try:
            async for message in KafkaConsumerClient.consume_messages_async(consumer):
                if message is not None:
                    print(f"Received message: {message.value.decode('utf-8')}")
        except Exception as e:
            print(f"Error {e}")
        finally:
            consumer.stop()

    @staticmethod
    async def consume_messages_async(consumer):
        while True:
            message = consumer.consume(block=False)
            if message:
                yield message
            await asyncio.sleep(0.1)
