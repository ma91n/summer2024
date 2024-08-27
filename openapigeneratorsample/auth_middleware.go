package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ma91n/summer2024/openapigeneratorsample/openapi"
)

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
	ErrClaimsInvalid     = errors.New("provided claims do not match expected scopes")
)

// genjwtで生成された秘密鍵（本来は認可サーバから取得する）
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

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/hello":
			next.ServeHTTP(w, r)
		default:
			if err := validateToken(r); err != nil {
				status := http.StatusUnauthorized
				_ = openapi.EncodeJSONResponse(map[string]any{"message": err.Error()}, &status, w)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}

func validateToken(r *http.Request) error {
	jws, err := getJWSFromRequest(r)
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

	// scopeの取り方が無い？
	scopes := []string{"read:hellos", "write:hellos"}
	// https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go-server.md#security-feature
	if err = checkTokenClaims(scopes, userClaim); err != nil {
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
