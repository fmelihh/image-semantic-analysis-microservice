package main

import (
	"log"
	"upload-service/cmd/api"
)

func main() {
	server := api.NewApiServer(":8000")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
