package images

import (
	"net/http"
	"upload-service/types"

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

}
