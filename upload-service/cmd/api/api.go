package api

import (
	"log"
	"net/http"
	"upload-service/service/images"
	"upload-service/service/kafkaProducer"

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
	kafkaProducerService := kafkaProducer.NewKafkaProducerService()

	imageHandler := images.NewHandler(imageService, kafkaProducerService)
	imageHandler.RegisterRoutes(subrouter)

	log.Println("Image upload service started on port ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
