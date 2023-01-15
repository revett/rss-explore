package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/revett/rss-explore/pkg/api"
)

// Convert takes a YouTube URL and hands back an RSS feed URL to the YouTube
// creator's channel.
func Convert(ctx echo.Context) error {
	var body api.ConvertURL
	if err := ctx.Bind(&body); err != nil {
		return returnErrorType(ctx, http.StatusBadRequest, "invalid request body")
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return returnErrorType(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := ctx.String(http.StatusOK, "rss-exploreeee"); err != nil {
		return fmt.Errorf("sending string response: %w", err)
	}

	return nil
}
