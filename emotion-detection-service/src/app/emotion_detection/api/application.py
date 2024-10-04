import fastapi
import asyncio
from contextlib import asynccontextmanager

from ..kafka.consumer import KafkaConsumerClient


@asynccontextmanager
async def lifespan(app: fastapi.FastAPI):
    asyncio.create_task(KafkaConsumerClient.consume_kafka_messages())
    yield

backend_app = fastapi.FastAPI(lifespan=lifespan)
