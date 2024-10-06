import json
import asyncio
from pykafka import KafkaClient

from ..services.emotion import EmotionService


class KafkaConsumerClient:
    @staticmethod
    async def consume_kafka_messages():
        client = KafkaClient("localhost:29092")
        topic = client.topics["image-upload"]

        consumer = topic.get_simple_consumer(
            auto_commit_enable=True, reset_offset_on_start=False
        )

        emotion_service = EmotionService()
        try:
            async for message in KafkaConsumerClient.consume_messages_async(consumer):
                if message is not None:
                    message_value = json.loads(message.value)
                    await emotion_service.predict(
                        image_name=message_value["Name"],
                        mime_type=message_value["Name"],
                        image_url=message_value["LocationUrl"],
                    )
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
