import os
import cv2
import cv2.data
import numpy as np
from keras.src.models.model import model_from_json


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
                f'{os.getcwd().split("/src")[0]}/src/app/emotion_detection/models'
            )
            with open(curr_path + "/emotion_model.joblib", "rb") as f:
                loaded_model_json = f.read()

            self._emotion_model = model_from_json(loaded_model_json)
            self._emotion_model.load_weights(curr_path + "/emotion_model.weights.h5")

        return self._emotion_model
