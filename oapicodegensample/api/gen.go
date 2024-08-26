//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

const (
	BearerScopes = "Bearer.Scopes"
	OAuth2Scopes = "OAuth2.Scopes"
	OIDCScopes   = "OIDC.Scopes"
)

// Hello defines model for Hello.
type Hello struct {
	Message *string `json:"message,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// hello👋
	// (GET /hello)
	Hello(w http.ResponseWriter, r *http.Request)
	// hello bearer👋
	// (GET /hello-bearer)
	HelloBearer(w http.ResponseWriter, r *http.Request)
	// hello oauth2👋
	// (GET /hello-oauth2)
	HelloOAuth2(w http.ResponseWriter, r *http.Request)
	// hello openid connect👋
	// (GET /hello-oidc)
	HelloOIDC(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// Hello operation middleware
func (siw *ServerInterfaceWrapper) Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Hello(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// HelloBearer operation middleware
func (siw *ServerInterfaceWrapper) HelloBearer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{"write:hellos", "read:hellos"})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HelloBearer(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// HelloOAuth2 operation middleware
func (siw *ServerInterfaceWrapper) HelloOAuth2(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, OAuth2Scopes, []string{"write:hellos", "read:hellos"})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HelloOAuth2(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// HelloOIDC operation middleware
func (siw *ServerInterfaceWrapper) HelloOIDC(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, OIDCScopes, []string{"write:hellos", "read:hellos"})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HelloOIDC(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/hello", wrapper.Hello)
	m.HandleFunc("GET "+options.BaseURL+"/hello-bearer", wrapper.HelloBearer)
	m.HandleFunc("GET "+options.BaseURL+"/hello-oauth2", wrapper.HelloOAuth2)
	m.HandleFunc("GET "+options.BaseURL+"/hello-oidc", wrapper.HelloOIDC)

	return m
}

type HelloJSONResponse Hello

type HelloRequestObject struct {
}

type HelloResponseObject interface {
	VisitHelloResponse(w http.ResponseWriter) error
}

type Hello200JSONResponse struct{ HelloJSONResponse }

func (response Hello200JSONResponse) VisitHelloResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type HelloBearerRequestObject struct {
}

type HelloBearerResponseObject interface {
	VisitHelloBearerResponse(w http.ResponseWriter) error
}

type HelloBearer200JSONResponse struct{ HelloJSONResponse }

func (response HelloBearer200JSONResponse) VisitHelloBearerResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type HelloOAuth2RequestObject struct {
}

type HelloOAuth2ResponseObject interface {
	VisitHelloOAuth2Response(w http.ResponseWriter) error
}

type HelloOAuth2200JSONResponse struct{ HelloJSONResponse }

func (response HelloOAuth2200JSONResponse) VisitHelloOAuth2Response(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type HelloOIDCRequestObject struct {
}

type HelloOIDCResponseObject interface {
	VisitHelloOIDCResponse(w http.ResponseWriter) error
}

type HelloOIDC200JSONResponse struct{ HelloJSONResponse }

func (response HelloOIDC200JSONResponse) VisitHelloOIDCResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// hello👋
	// (GET /hello)
	Hello(ctx context.Context, request HelloRequestObject) (HelloResponseObject, error)
	// hello bearer👋
	// (GET /hello-bearer)
	HelloBearer(ctx context.Context, request HelloBearerRequestObject) (HelloBearerResponseObject, error)
	// hello oauth2👋
	// (GET /hello-oauth2)
	HelloOAuth2(ctx context.Context, request HelloOAuth2RequestObject) (HelloOAuth2ResponseObject, error)
	// hello openid connect👋
	// (GET /hello-oidc)
	HelloOIDC(ctx context.Context, request HelloOIDCRequestObject) (HelloOIDCResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// Hello operation middleware
func (sh *strictHandler) Hello(w http.ResponseWriter, r *http.Request) {
	var request HelloRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Hello(ctx, request.(HelloRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Hello")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HelloResponseObject); ok {
		if err := validResponse.VisitHelloResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// HelloBearer operation middleware
func (sh *strictHandler) HelloBearer(w http.ResponseWriter, r *http.Request) {
	var request HelloBearerRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.HelloBearer(ctx, request.(HelloBearerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HelloBearer")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HelloBearerResponseObject); ok {
		if err := validResponse.VisitHelloBearerResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// HelloOAuth2 operation middleware
func (sh *strictHandler) HelloOAuth2(w http.ResponseWriter, r *http.Request) {
	var request HelloOAuth2RequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.HelloOAuth2(ctx, request.(HelloOAuth2RequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HelloOAuth2")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HelloOAuth2ResponseObject); ok {
		if err := validResponse.VisitHelloOAuth2Response(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// HelloOIDC operation middleware
func (sh *strictHandler) HelloOIDC(w http.ResponseWriter, r *http.Request) {
	var request HelloOIDCRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.HelloOIDC(ctx, request.(HelloOIDCRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HelloOIDC")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HelloOIDCResponseObject); ok {
		if err := validResponse.VisitHelloOIDCResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUsW4UMRD9FWug3Owel247kggRGoqAKKIrHO/srZNd27K9HEe0DWmQEEJ8AT1NClr+",
	"5oREySegsX13XMJxAtKNZ59n3hvP20sQujNaofIOykuw6IxWDsPhMbatpkBo5VF5CrkxrRTcS62Kc6cV",
	"5ZxosOMU3bdYQwn3inXVIn51Raw2DEMGFTphpaEiUMJJLwQ6V/ct0wZtqA0ESzc3qBhLGC8jww6d41Ok",
	"EF/xzrQIJTQBm4GfGzo6b6WahoIpo8/OUfjYAkVvpZ+fUKtY8wC5RUvRWYgeadtxDyU8efEMbnKP4MXV",
	"28XV18Wb68XVl++f33/7cA2JPbWLZdaEGu8NDBk8fdj7ZkyNNmuGPBvno1WputWzwI33vtFWvg4zOtQV",
	"3ko+t21q4cqiSEPJhe4KTbhxsUQjZGCxtuianXcSLojSBtOm8KoMo3ZQhhMLJyYVm+veMi6E7pWHDGZW",
	"elxjO13Jer4FTa+kL3C3joDaeNWQD4M9Pjr8zVgNquMjdqiVotfPQFOiSuft/fIZtu3ehdIzVdAVWe0J",
	"rWo57dOqrhn8WjBuulR1WFsvfbuEcCPZfj5ixJcJXSGboqK915a5uMQZvETrIu8H+Sgfkax0F0rYz0f5",
	"GDIw3DfhMYpm6Y8pBpuujHRcQZnck226ezwabbPsCrcybQau7zpu50t//fj08R0p51MH5SkYstiEcJHK",
	"3tnKRdsZHSyN8V+8koOhPF1793Rz57KNbZ0Mk1tyWOS7S1XasT+qSra+Q1XLH8Xfq4p8d6qSldihiRx1",
	"l4qCQ/9BT/AfE9FiW3QNq9TNX0CDvCXPNSguIAPFw/853Bsmw88AAAD//2EKF94MBwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
