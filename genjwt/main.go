package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func main() {

	// make privatekey
	privateKey := `-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIA64IEwEEBhle3vJ2WxIHcYpuLP/0u1vjUq6WO4l0sVbuojHJ7pb1g
ZLHzpOEpTi0ZdRIBb02hxRe8iqhZTRFRBUqgBwYFK4EEACOhgYkDgYYABAGG9gPW
0o+10f19UWvct5hYXsPjcJcNUNr7uyaJY4zVGQcXFwDTiMnhPSnthkkBkGp28x22
TTt+5TzE7BMbdVZJZgGxC+WJiIgp26LuMWNw3Zt3j/Bn8qidQtMpdqW9dBhgvj1n
YcfWHGKnUL01r4BO7z9uNmWj2n05Jv1IrnVegMXrGg==
-----END EC PRIVATE KEY-----`

	t := jwt.NewWithClaims(jwt.SigningMethodES512,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "123",
			"scp": "read:hellos write:hellos",
		})

	// 参考: https://stackoverflow.com/questions/21322182/how-to-store-ecdsa-private-key-in-go
	block, _ := pem.Decode([]byte(privateKey))

	ecPrivateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	s, err := t.SignedString(ecPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

}
