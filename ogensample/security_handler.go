package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"githu.com/ma91n/summer2024/ogensample/api"
	"github.com/golang-jwt/jwt/v5"
)

var ErrClaimsInvalid = errors.New("provided claims do not match expected scopes")

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

type MySecurityHandler struct{}

func (o MySecurityHandler) HandleBearer(ctx context.Context, operationName string, t api.Bearer) (context.Context, error) {
	return o.handleToken(ctx, t.Token, []string{})
}

func (o MySecurityHandler) HandleOAuth2(ctx context.Context, operationName string, t api.OAuth2) (context.Context, error) {
	return o.handleToken(ctx, t.Token, t.Scopes)
}

func (o MySecurityHandler) handleToken(ctx context.Context, jwtToken string, expectedClaims []string) (context.Context, error) {
	var userClaim UserClaim
	token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (any, error) {
		// 参考: https://stackoverflow.com/questions/21322182/how-to-store-ecdsa-private-key-in-go
		blockPub, _ := pem.Decode([]byte(jwtKey))
		return x509.ParsePKIXPublicKey(blockPub.Bytes)
	})
	if err != nil {
		return ctx, err
	}

	if !token.Valid {
		return ctx, errors.New("invalid token")
	}

	if err = checkTokenClaims(expectedClaims, userClaim); err != nil {
		return ctx, fmt.Errorf("token claims don't match: %w", err)
	}

	return nil, nil
}

// checkTokenClaims はスコープのチェックを行う

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
