package openapi

import (
	"context"
	"net/http"
)

type PingAPIService struct {
}

func NewPingAPIService() *PingAPIService {
	return &PingAPIService{}
}

// Hello - hello👋
func (s *PingAPIService) Hello(_ context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}

// HelloBearer - hello bearer👋
func (s *PingAPIService) HelloBearer(ctx context.Context) (ImplResponse, error) {
	return s.Hello(ctx)
}

// HelloOAuth2 - hello oauth2👋
func (s *PingAPIService) HelloOAuth2(ctx context.Context) (ImplResponse, error) {
	return s.Hello(ctx)
}

// HelloOIDC - hello openid connect👋
func (s *PingAPIService) HelloOIDC(ctx context.Context) (ImplResponse, error) {
	return s.Hello(ctx)
}
