package main

import (
	"log"
	"net/http"

	"github.com/ma91n/summer2024/openapigeneratorsample/openapi"
)

func main() {
	log.Printf("Server started")

	controller := openapi.NewPingAPIController(openapi.NewPingAPIService())
	handler := openapi.NewRouter(controller)
	handler.Use(Authentication)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
