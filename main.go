package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	commonLog "github.com/revett/common/log"
	commonMiddleware "github.com/revett/common/middleware"
	"github.com/revett/rss-explore/internal/server"
	"github.com/revett/rss-explore/pkg/api"
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

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal().Err(err).Msg("loading swagger spec")
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match.
	swagger.Servers = nil

	api.RegisterHandlers(e, server.Server{})

	if err := e.Start(port); err != nil {
		log.Fatal().Err(err).Send()
	}
}
