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
func (s *PingAPIService) Hello(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}

// HelloBearer - hello bearer👋
func (s *PingAPIService) HelloBearer(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}

// HelloOAuth2 - hello oauth2👋
func (s *PingAPIService) HelloOAuth2(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}

// HelloOIDC - hello openid connect👋
func (s *PingAPIService) HelloOIDC(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}
