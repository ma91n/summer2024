package main

import (
	"context"
	"githu.com/ma91n/summer2024/ogensample/api"
)

type HelloHandler struct{}

func (h HelloHandler) Hello(_ context.Context) (*api.Hello, error) {
	return &api.Hello{Message: api.OptString{Value: "hello", Set: true}}, nil
}

func (h HelloHandler) HelloOIDC(_ context.Context) (*api.Hello, error) {
	return &api.Hello{Message: api.OptString{Value: "hello", Set: true}}, nil
}

func (h HelloHandler) HelloBearer(_ context.Context) (*api.Hello, error) {
	return &api.Hello{Message: api.OptString{Value: "hello", Set: true}}, nil
}

func (h HelloHandler) HelloOAuth2(_ context.Context) (*api.Hello, error) {
	return &api.Hello{Message: api.OptString{Value: "hello", Set: true}}, nil
}
