package images

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upload-service/types"
	"upload-service/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	service              types.ImageUploadService
	kafkaProducerService types.KafkaProducerService
}

func NewHandler(service types.ImageUploadService, kafkaProducerService types.KafkaProducerService) *Handler {
	return &Handler{service: service, kafkaProducerService: kafkaProducerService}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/image-upload", h.handleImageUpload).Methods(http.MethodPost)
}

func (h *Handler) handleImageUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200 << 20)
	f, header, err := r.FormFile("image")
	if err != nil {
		fmt.Printf("Error reading file of 'image' form data. Reason %s\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	imageMetadata, err := h.service.SaveImage(f, header)
	if err != nil {
		fmt.Printf("Error saving file to minio is not successfully completed. Reason %s\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	imageMetadataBytes, err := json.Marshal(imageMetadata)
	if err != nil {
		fmt.Printf("Error converting image metadata object to bytes. Reason %s\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	h.kafkaProducerService.PushMessage("image-upload", imageMetadataBytes)
	utils.WriteJSON(w, http.StatusAccepted, map[string]string{"status": "success"})
}
