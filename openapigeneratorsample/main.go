package main

import (
	"github.com/ma91n/summer2024/openapigeneratorsample/middleware"
	"github.com/ma91n/summer2024/openapigeneratorsample/openapi"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	controller := openapi.NewPingAPIController(openapi.NewPingAPIService())
	handler := openapi.NewRouter(controller)
	handler.Use(middleware.Authentication)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
