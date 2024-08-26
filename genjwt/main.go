package main

import (
	_ "embed"

	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

//go:embed ecprivatekey.pem
var privateKey []byte

func main() {
	// 参考: https://stackoverflow.com/questions/21322182/how-to-store-ecdsa-private-key-in-go
	block, _ := pem.Decode(privateKey)

	ecPrivateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	s, err := jwt.NewWithClaims(jwt.SigningMethodES512,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "123",
			"scp": "read:hellos write:hellos",
		}).SignedString(ecPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
