package main

import (
	"fmt"
	"log"
	"upload-service/cmd/api"
)

func main() {
	server := api.NewApiServer(":8000")
	fmt.Println("Image upload service started on port :8000")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
