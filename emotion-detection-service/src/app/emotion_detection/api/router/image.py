from fastapi import APIRouter

from ..application import backend_app

image_router = APIRouter()


@image_router.get("/test")
def test():
    return {"ping": "pong"}


backend_app.include_router(image_router, prefix="/image", tags=["image"])
