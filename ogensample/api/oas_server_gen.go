// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// Hello implements hello operation.
	//
	// Hello👋.
	//
	// GET /hello
	Hello(ctx context.Context) (*Hello, error)
	// HelloBearer implements helloBearer operation.
	//
	// Hello bearer👋.
	//
	// GET /hello-bearer
	HelloBearer(ctx context.Context) (*Hello, error)
	// HelloOAuth2 implements helloOAuth2 operation.
	//
	// Hello oauth2👋.
	//
	// GET /hello-oauth2
	HelloOAuth2(ctx context.Context) (*Hello, error)
	// HelloOIDC implements helloOIDC operation.
	//
	// Hello openid connect👋.
	//
	// GET /hello-oidc
	HelloOIDC(ctx context.Context) (*Hello, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
