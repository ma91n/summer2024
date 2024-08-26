package main

import (
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/oapi-codegen/nethttp-middleware"

	"github.com/ma91n/summer2024/oapicodegensample/api"
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
				AuthenticationFunc: api.NewAuthenticator(),
			},
		})

	strictHandler := api.NewStrictHandler(api.Server{}, nil)
	handler := api.HandlerFromMux(strictHandler, http.NewServeMux())

	s := &http.Server{
		Handler: mw(handler),
		Addr:    "0.0.0.0:8080",
	}
	log.Fatal(s.ListenAndServe())
}
