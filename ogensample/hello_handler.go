package main

import (
	"context"
	"githu.com/ma91n/summer2024/ogensample/api"
)

type HelloHandler struct{}

func (h HelloHandler) Hello(_ context.Context) (*api.Hello, error) {
	return &api.Hello{Message: api.OptString{Value: "hello", Set: true}}, nil
}

func (h HelloHandler) HelloBearer(ctx context.Context) (*api.Hello, error) {
	return h.Hello(ctx)
}

func (h HelloHandler) HelloOAuth2(ctx context.Context) (*api.Hello, error) {
	return h.Hello(ctx)
}

func (h HelloHandler) HelloOIDC(ctx context.Context) (*api.Hello, error) {
	return h.Hello(ctx)
}
