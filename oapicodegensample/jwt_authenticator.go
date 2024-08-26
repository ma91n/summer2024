package main

// https://github.com/oapi-codegen/oapi-codegen/blob/main/examples/authenticated-api/stdhttp/server/jwt_authenticator.go

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
	ErrClaimsInvalid     = errors.New("provided claims do not match expected scopes")
)

var jwtKey = `-----BEGIN PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBhvYD1tKPtdH9fVFr3LeYWF7D43CX
DVDa+7smiWOM1RkHFxcA04jJ4T0p7YZJAZBqdvMdtk07fuU8xOwTG3VWSWYBsQvl
iYiIKdui7jFjcN2bd4/wZ/KonULTKXalvXQYYL49Z2HH1hxip1C9Na+ATu8/bjZl
o9p9OSb9SK51XoDF6xo=
-----END PUBLIC KEY-----`

type UserClaim struct {
	jwt.RegisteredClaims
	Scope string `json:"scp"`
}

func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		securitySchemeName := input.SecuritySchemeName
		switch securitySchemeName {
		case "": // 認証なし
			return nil
		case "Bearer":
			return validateSecurityScheme(input)
		case "OAuth2":
			return validateSecurityScheme(input)
		case "OIDC":
			return validateSecurityScheme(input)
		default:
			panic("not supported security scheme " + securitySchemeName)
		}
	}
}

func validateSecurityScheme(input *openapi3filter.AuthenticationInput) error {
	jws, err := getJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	var userClaim UserClaim
	token, err := jwt.ParseWithClaims(jws, &userClaim, func(token *jwt.Token) (any, error) {
		// 参考: https://stackoverflow.com/questions/21322182/how-to-store-ecdsa-private-key-in-go
		blockPub, _ := pem.Decode([]byte(jwtKey))
		return x509.ParsePKIXPublicKey(blockPub.Bytes)
	})
	if err != nil {
		return fmt.Errorf("validating JWS: %w", err)
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	if err = checkTokenClaims(input.Scopes, userClaim); err != nil {
		return fmt.Errorf("token claims don't match: %w", err)
	}

	return nil
}

func getJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	if authHdr == "" {
		return "", ErrNoAuthHeader
	}

	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}

	return strings.TrimPrefix(authHdr, prefix), nil
}

func checkTokenClaims(expectedClaims []string, t UserClaim) error {
	claims := strings.Split(t.Scope, " ")
	claimsMap := make(map[string]bool, len(claims))
	for _, c := range claims {
		claimsMap[c] = true
	}

	for _, e := range expectedClaims {
		if !claimsMap[e] {
			return ErrClaimsInvalid
		}
	}

	return nil
}
