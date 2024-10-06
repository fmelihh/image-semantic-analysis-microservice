import os
import cv2
import requests
import cv2.data
import numpy as np
from keras.src.models.model import model_from_json

from ..kafka.producer import KafkaProducerClient


class EmotionService:
    def __init__(self):
        self.emotion_dict = {
            0: "Angry",
            1: "Disgusted",
            2: "Fearful",
            3: "Happy",
            4: "Neutral",
            5: "Sad",
            6: "Surprised",
        }
        self._face_cascade = None
        self._emotion_model = None

    @property
    def face_cascade(self) -> cv2.CascadeClassifier:
        if self._face_cascade is None:
            self._face_cascade = cv2.CascadeClassifier(
                cv2.data.haarcascades + "haarcascade_frontalface_default.xml"
            )
        return self._face_cascade

    @property
    def emotion_model(self):
        if self._emotion_model is None:
            curr_path = (
                f'{os.getcwd().split("/src")[0]}/src/app/emotion_detection/model'
            )
            with open(curr_path + "/emotion_model.json", "rb") as f:
                loaded_model_json = f.read()

            self._emotion_model = model_from_json(loaded_model_json)
            self._emotion_model.load_weights(curr_path + "/emotion_model.weights.h5")

        return self._emotion_model

    async def predict(self, image_name: str, mime_type: str, image_url: str):
        print("message received from upload service", image_name, mime_type)
        response = requests.get(image_url)
        if response.status_code != 200:
            return

        image_np = np.frombuffer(response.content, np.uint8)
        img = cv2.imdecode(image_np, cv2.IMREAD_COLOR)
        img = cv2.resize(img, (600, 600))
        img = cv2.convertScaleAbs(img, beta=50)
        gray_img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
        faces = self.face_cascade.detectMultiScale(gray_img, scaleFactor=1.3, minNeighbors=5)
        if len(faces) == 0:
            print("No face detected!")
            return

        (x, y, w, h) = faces[0]
        face_img = gray_img[y: y + h, x:x + w]
        resized_img = cv2.resize(face_img, (48, 48))
        img_pixels = resized_img.astype('float32') / 255.0
        img_pixels = np.expand_dims(img_pixels, axis=-1)
        img_pixels = np.expand_dims(img_pixels, axis=0)
        predictions = self.emotion_model.predict(img_pixels)
        max_index = np.argmax(predictions[0])
        predicted_emotion = self.emotion_dict[max_index]
        message_body = {
            "ImageURL": image_url,
            "MimeType": mime_type,
            "ImageName": image_name,
            "Emotion": predicted_emotion,
            "AvailableEmotions": ','.join(list(self.emotion_dict.values())),
        }
        print("Predicted emotion is", predicted_emotion)
        await KafkaProducerClient.produce_kafka_messages(message_body)
