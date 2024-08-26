package openapi

import (
	"context"
	"net/http"
)

// Summer2024APIService is a service that implements the logic for the Summer2024APIServicer
// This service should implement the business logic for every endpoint for the Summer2024API API.
// Include any external packages or services that will be required by this service.
type Summer2024APIService struct{}

// NewSummer2024APIService creates a default api service
func NewSummer2024APIService() *Summer2024APIService {
	return &Summer2024APIService{}
}

// Hello - hello👋
func (s *Summer2024APIService) Hello(_ context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}

// HelloBearer - hello bearer👋
func (s *Summer2024APIService) HelloBearer(_ context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}

// HelloOAuth2 - hello oauth2👋
func (s *Summer2024APIService) HelloOAuth2(_ context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}

// HelloOIDC - hello openid connect👋
func (s *Summer2024APIService) HelloOIDC(_ context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, Hello{Message: "Hello"}), nil
}
