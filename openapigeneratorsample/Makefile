install:
	npm install @openapitools/openapi-generator-cli -g
	openapi-generator-cli version-manager set 7.8.0

# https://openapi-generator.tech/docs/generators/go-server/
generate:
	npx @openapitools/openapi-generator-cli generate -i openapi.yaml -g go-server --additional-properties=outputAsLibrary=true,sourceFolder=openapi -o .
	gofmt -l -w .

run:
	go run .

curl:
	curl localhost:8080/hello

curl-bearer:
	curl -H "Authorization: Bearer eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS1hdXRoLXNlcnZlciIsInNjcCI6InJlYWQ6aGVsbG9zIHdyaXRlOmhlbGxvcyIsInN1YiI6IjEyMyJ9.AUF4V06rFT5ZptwBcXrJJPS6GQ_ktZSyps1o6gR8XE7F2YNhyUyl7InmfQQhv6dtc0UBzRY9uIcE5wWgqyQXc1n-AZDHddwjUzWxO2c7gYmoevFClLXim2MGyL7FvimYaIFvOs69awFot9gcIIce6eWSGbAkzAVygrHfH6T7hpAwDF7G" localhost:8080/hello-bearer

curl-oauth2:
	curl -H "Authorization: Bearer eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS1hdXRoLXNlcnZlciIsInNjcCI6InJlYWQ6aGVsbG9zIHdyaXRlOmhlbGxvcyIsInN1YiI6IjEyMyJ9.AUF4V06rFT5ZptwBcXrJJPS6GQ_ktZSyps1o6gR8XE7F2YNhyUyl7InmfQQhv6dtc0UBzRY9uIcE5wWgqyQXc1n-AZDHddwjUzWxO2c7gYmoevFClLXim2MGyL7FvimYaIFvOs69awFot9gcIIce6eWSGbAkzAVygrHfH6T7hpAwDF7G" localhost:8080/hello-oauth2

curl-oidc:
	curl -H "Authorization: Bearer eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS1hdXRoLXNlcnZlciIsInNjcCI6InJlYWQ6aGVsbG9zIHdyaXRlOmhlbGxvcyIsInN1YiI6IjEyMyJ9.AUF4V06rFT5ZptwBcXrJJPS6GQ_ktZSyps1o6gR8XE7F2YNhyUyl7InmfQQhv6dtc0UBzRY9uIcE5wWgqyQXc1n-AZDHddwjUzWxO2c7gYmoevFClLXim2MGyL7FvimYaIFvOs69awFot9gcIIce6eWSGbAkzAVygrHfH6T7hpAwDF7G" localhost:8080/hello-oidc
