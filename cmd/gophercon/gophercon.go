package main

import (
	"net/http"
	"log"
	"github.com/mtekkie/workshop-go/pkg/routing"
)

// go run ./cmd/gophercon/gophercon.go
func main() {

	r := routing.BaseRouter()

	log.Printf("Service is starting")
	http.ListenAndServe(":8000", r)
}


