// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ConvertURL defines model for ConvertURL.
type ConvertURL struct {
	// Url is the YouTube URL to convert.
	URL string `json:"url" validate:"required,url"`
}

// Error defines model for Error.
type Error struct {
	// Code is the HTTP status code for the error.
	Code int32 `json:"code"`

	// Message is the message of the error.
	Message string `json:"message"`
}

// RSSFeed defines model for RSSFeed.
type RSSFeed struct {
	// Url is the RSS feed for the YouTube channel.
	URL string `json:"url"`
}

// ConvertJSONRequestBody defines body for Convert for application/json ContentType.
type ConvertJSONRequestBody = ConvertURL

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Converts a YouTube URL in to a RSS feed for the YouTube channel.
	// (POST /youtube/convert)
	Convert(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Convert converts echo context to params.
func (w *ServerInterfaceWrapper) Convert(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Convert(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/youtube/convert", wrapper.Convert)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUQW/bPAz9KwK/7+jEaYsNmG/b0GEFeiiS9DAUPagyHWuwRU2iggSF//tAOWmaNgU6",
	"oKfYpPn4+PiYRzDUe3LoOEL1CNG02Ov8+J3cGgPfzq/lzQfyGNhizqXQyU+N0QTr2ZKDCmxU3KL6RWmZ",
	"HlDdzq8VkzIjzBQK4K1HqCBysG4FBWwmK5o43UtQ2kiEtLcTQzWu0E1ww0FPWK9y07XubK1Zvg74J9mA",
	"dSFEhmEoniJQ3WV290MBlyFQeE1e0N9k/3O5vFGRNaeo5EPVUMgJFDCZoqHQa5YKxxfnh7GsY1xhgKGA",
	"HmPUq7eb7PKKmmPoY4FejpV5H8Blwvli8QMl+y8Lmi8WqkGsn0bbb8y02jns3rOq05JL1LqGXrdeEnVR",
	"7NBi59WWkkIdbbdVjXW10q5WvXYiyZ5czCwsd9JRgpcb31FA9fXmCgpYY4gj8mx6Np2J6OTRaW+hgosc",
	"KsBrbrMa5ZYSpwcsd2bMelHk1zx3po9KH/nYOuGu3yWdrEEL3FV9AIRRLoz8jert6ELH6DIF7X1nTa4p",
	"f0fhsT9Eefo/YAMV/FceLrXcnWn57Eaz+MfDnL5EeL45DgnzKqMnF0fvnM9mH0Zwb9AT7BbJGIyxEA35",
	"mYZPIkMuaXTq+MP4jH8JJ9gkhxuPhrEezzE7PKa+12H7Eb7ILSMGMS5Udy99dxOoTia/FOPxQsvsY1WW",
	"IcYJjuafBlybemqoh6F4CXFNRndH1VVZdhJsKXL16fOXMxjuh78BAAD//xGRHLfyBQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
