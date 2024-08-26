package main

import (
	"githu.com/ma91n/summer2024/ogensample/api"
	"log"
	"net/http"
)

func main() {
	srv, err := api.NewServer(&HelloHandler{}, SecurityHandler{})
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
