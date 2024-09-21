package api

import (
	"log"
	"net/http"
	"upload-service/service/images"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
)

type ApiServer struct {
	addr        string
	minioClient *minio.Client
}

func NewApiServer(addr string, minioClient *minio.Client) *ApiServer {
	return &ApiServer{addr: addr, minioClient: minioClient}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	imageService := images.NewService(s.minioClient)
	imageHandler := images.NewHandler(imageService)
	imageHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
