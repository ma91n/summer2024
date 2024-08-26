package main

import (
	"log"
	"net/http"

	"githu.com/ma91n/summer2024/ogensample/api"
)

func main() {
	srv, err := api.NewServer(&HelloHandler{}, MySecurityHandler{})
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
