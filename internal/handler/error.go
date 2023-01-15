package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/revett/rss-explore/pkg/api"
)

func returnErrorType(ctx echo.Context, code int, message string) error {
	errType := api.Error{
		Code:    int32(code),
		Message: message,
	}

	if err := ctx.JSON(code, errType); err != nil {
		return fmt.Errorf("sending json response with status code: %w", err)
	}

	return nil
}
