package images

import (
	"fmt"
	"net/http"
	"upload-service/types"
	"upload-service/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	service types.ImageUploadService
}

func NewHandler(service types.ImageUploadService) *Handler {
	return &Handler{service: service}
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
	utils.WriteJSON(w, http.StatusAccepted, map[string]any{
		"message":       "image was saved",
		"imageName":     imageMetadata.Name,
		"imageMimeType": imageMetadata.MimeType,
		"imageBytes":    imageMetadata.Bytes,
	})
}
