generate:
	go run .

genprivatekey:
	openssl ecparam -name secp521r1 -genkey -noout -out ecprivatekey.pem

genpubkey:
	openssl ec -in ecprivatekey.pem -pubout -out ecpubkey.pem
