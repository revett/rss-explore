package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/revett/rss-explore/internal/youtube"
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

	videoID, err := youtube.ExtractVideoID(body.URL)
	if err != nil {
		return returnErrorType(ctx, http.StatusBadRequest, "invalid youtube url")
	}

	resp := api.RSSFeed{
		URL: videoID,
	}

	if err := ctx.JSON(http.StatusOK, resp); err != nil {
		return fmt.Errorf("sending json response: %w", err)
	}

	return nil
}
