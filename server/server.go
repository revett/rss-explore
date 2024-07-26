package server

import (
	"github.com/labstack/echo/v4"
	"github.com/revett/rss-explore/handler"
)

// Server implements the api.ServerInterface interface.
type Server struct{}

// Convert implements the api.ServerInterface.Convert method.
func (h Server) Convert(ctx echo.Context) error {
	return handler.Convert(ctx) //nolint:wrapcheck
}
