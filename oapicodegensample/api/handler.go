package api

import (
	"context"
)

type Server struct{}

func (s Server) Hello(_ context.Context, _ HelloRequestObject) (HelloResponseObject, error) {
	return Hello200JSONResponse{Message: ptr("hello")}, nil
}

func (s Server) HelloBearer(_ context.Context, _ HelloBearerRequestObject) (HelloBearerResponseObject, error) {
	return HelloBearer200JSONResponse{Message: ptr("hello")}, nil
}

func (s Server) HelloOAuth2(_ context.Context, _ HelloOAuth2RequestObject) (HelloOAuth2ResponseObject, error) {
	return HelloOAuth2200JSONResponse{Message: ptr("hello")}, nil
}

func (s Server) HelloOIDC(_ context.Context, _ HelloOIDCRequestObject) (HelloOIDCResponseObject, error) {
	return HelloOIDC200JSONResponse{Message: ptr("hello")}, nil
}

func ptr[T any](t T) *T {
	return &t
}
