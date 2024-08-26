package main

import (
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/ma91n/summer2024/oapicodegensample/api"
	"github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	spec, err := api.GetSwagger()
	if err != nil {
		log.Fatalln("loading spec: ", err)
	}
	spec.Servers = nil

	mw := nethttpmiddleware.OapiRequestValidatorWithOptions(spec,
		&nethttpmiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticator(),
			},
		})

	strictHandler := api.NewStrictHandler(Server{}, nil)
	handler := api.HandlerFromMux(strictHandler, http.NewServeMux())

	s := &http.Server{
		Handler: mw(handler),
		Addr:    "0.0.0.0:8080",
	}
	log.Fatal(s.ListenAndServe())
}
