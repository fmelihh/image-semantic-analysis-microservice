import json
from typing import Any
from pykafka import KafkaClient

class KafkaProducerClient:
    @staticmethod
    async def produce_kafka_messages(message: dict[str, Any]):
        client = KafkaClient("localhost:29092")
        topic = client.topics["emotion-response"]
        producer = topic.get_sync_producer()

        try:
            producer.produce(json.dumps(message).encode("utf-8"))
            print("message produced.", json.dumps(message, indent=3))
        except Exception as e:
            print(f"Error {e}")
        finally:
            producer.stop()