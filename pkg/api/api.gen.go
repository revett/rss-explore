// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
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

	// Invoke the callback with all the unmarshaled arguments
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
