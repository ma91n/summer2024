package main

import (
	"context"

	"github.com/ma91n/summer2024/oapicodegensample/api"
)

type HelloServer struct{}

func (s HelloServer) Hello(_ context.Context, _ api.HelloRequestObject) (api.HelloResponseObject, error) {
	return api.Hello200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func (s HelloServer) HelloBearer(_ context.Context, _ api.HelloBearerRequestObject) (api.HelloBearerResponseObject, error) {
	return api.HelloBearer200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func (s HelloServer) HelloOAuth2(_ context.Context, _ api.HelloOAuth2RequestObject) (api.HelloOAuth2ResponseObject, error) {
	return api.HelloOAuth2200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func (s HelloServer) HelloOIDC(_ context.Context, _ api.HelloOIDCRequestObject) (api.HelloOIDCResponseObject, error) {
	return api.HelloOIDC200JSONResponse{HelloJSONResponse: api.HelloJSONResponse{Message: ptr("hello")}}, nil
}

func ptr[T any](t T) *T {
	return &t
}
