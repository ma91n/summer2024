package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"githu.com/ma91n/summer2024/ogensample/api"
	"github.com/golang-jwt/jwt/v5"
	"log"
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

type SecurityHandler struct{}

func (o SecurityHandler) HandleBearer(ctx context.Context, operationName string, t api.Bearer) (context.Context, error) {
	fmt.Printf("%v\n", operationName)
	return o.handleToken(ctx, t.Token)
}

func (o SecurityHandler) HandleOAuth2(ctx context.Context, operationName string, t api.OAuth2) (context.Context, error) {
	fmt.Printf("%v %v\n", operationName, t.Scopes)
	return o.handleToken(ctx, t.Token)
}

func (o SecurityHandler) handleToken(ctx context.Context, jwtToken string) (context.Context, error) {
	var userClaim UserClaim
	token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (any, error) {
		// 参考: https://stackoverflow.com/questions/21322182/how-to-store-ecdsa-private-key-in-go
		blockPub, _ := pem.Decode([]byte(jwtKey))
		return x509.ParsePKIXPublicKey(blockPub.Bytes)
	})
	if err != nil {
		log.Fatal(err)
	}

	if !token.Valid {
		return ctx, errors.New("invalid token")
	}

	return nil, nil
}
