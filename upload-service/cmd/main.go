package main

import (
	"log"
	"upload-service/cmd/api"
	"upload-service/db"
)

func main() {
	minioClient, err := db.NewFileStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(":8000", minioClient)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
