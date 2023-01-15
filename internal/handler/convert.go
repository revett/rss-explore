package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Convert takes a YouTube URL and hands back an RSS feed URL to the YouTube
// creator's channel.
func Convert(ctx echo.Context) error {
	if err := ctx.String(http.StatusOK, "rss-explore"); err != nil {
		return fmt.Errorf("sending string response: %w", err)
	}

	return nil
}
