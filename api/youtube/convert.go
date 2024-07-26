package youtube

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	commonLog "github.com/revett/common/log"
	commonMiddleware "github.com/revett/common/middleware"
	"github.com/revett/rss-explore/handler"
	"github.com/rs/zerolog/log"
)

// Convert is the entrypoint for the Vercel serverless function; it is a wrapper
// around the handler.Convert function.
func Convert(w http.ResponseWriter, r *http.Request) { //nolint:varnamelen
	log.Logger = commonLog.New()

	e := echo.New() //nolint:varnamelen
	e.Use(commonMiddleware.LoggerUsingZerolog(log.Logger))
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())
	e.Use(middleware.RecoverWithConfig(
		middleware.RecoverConfig{
			DisablePrintStack: true,
		},
	))

	e.POST("/*", handler.Convert)
	e.ServeHTTP(w, r)
}
