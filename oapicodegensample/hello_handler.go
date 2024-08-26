package main

import (
	"context"

	"github.com/ma91n/summer2024/oapicodegensample/api"
)

type Server struct{}

func (s Server) Hello(_ context.Context, _ api.HelloRequestObject) (api.HelloResponseObject, error) {
	return api.Hello200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func (s Server) HelloBearer(_ context.Context, _ api.HelloBearerRequestObject) (api.HelloBearerResponseObject, error) {
	return api.HelloBearer200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func (s Server) HelloOAuth2(_ context.Context, _ api.HelloOAuth2RequestObject) (api.HelloOAuth2ResponseObject, error) {
	return api.HelloOAuth2200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func (s Server) HelloOIDC(_ context.Context, _ api.HelloOIDCRequestObject) (api.HelloOIDCResponseObject, error) {
	return api.HelloOIDC200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func ptr[T any](t T) *T {
	return &t
}
