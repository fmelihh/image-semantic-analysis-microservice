import uvicorn
import asyncio

from src.app.emotion_detection.api import *
from src.app.emotion_detection.kafka import KafkaConsumerClient


@backend_app.on_event("startup")
def startup():
    asyncio.create_task(KafkaConsumerClient.consume_kafka_messages())


if __name__ == "__main__":
    uvicorn.run("main:backend_app", host="0.0.0.0", port=8000, reload=True)
