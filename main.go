package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	commonLog "github.com/revett/common/log"
	commonMiddleware "github.com/revett/common/middleware"
	"github.com/revett/rss-explore/internal/handler"
	"github.com/rs/zerolog/log"
)

const port = ":5691"

func main() {
	log.Logger = commonLog.New()

	e := echo.New() //nolint:varnamelen
	e.Use(commonMiddleware.LoggerUsingZerolog(log.Logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.RecoverWithConfig(
		middleware.RecoverConfig{
			DisablePrintStack: true,
		},
	))

	e.POST("/api/youtube/convert", handler.Convert)

	if err := e.Start(port); err != nil {
		log.Fatal().Err(err).Send()
	}
}
