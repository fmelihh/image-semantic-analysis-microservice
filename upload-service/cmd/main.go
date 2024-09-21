package main

import (
	"fmt"
	"log"
	"upload-service/cmd/api"
	"upload-service/db"
)

func main() {
	minioClient, err := db.NewFileStorage()
	if err != nil {
		log.Fatal(err)
		return
	}

	server := api.NewApiServer(":8000", minioClient)
	fmt.Println("Image upload service started on port :8000")

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
