package api

import (
	"log"
	"net/http"
	"upload-service/service/images"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	imageService := images.NewService()
	imageHandler := images.NewHandler(imageService)
	imageHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
