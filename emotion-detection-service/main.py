import uvicorn

from src.app.emotion_detection.api import *


if __name__ == "__main__":
    uvicorn.run("main:backend_app", host="0.0.0.0", port=8000, reload=True)
