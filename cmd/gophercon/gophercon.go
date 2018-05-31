package main

import (
	"net/http"
	"log"
	"github.com/mtekkie/workshop-go/pkg/routing"
	"os"
)

// go run ./cmd/gophercon/gophercon.go
func main() {

	log.Printf("Service is starting")
	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		log.Fatal("Servivce port was not set")
	}

	r := routing.BaseRouter()

	http.ListenAndServe(":"+port, r)
}


