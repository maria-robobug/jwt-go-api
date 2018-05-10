package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/maria-robobug/jwt-go-api/src/api/config"
)

func SetMainMiddlewares(e *echo.Echo) {
	// Echo: Logging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           config.LOG_FORMAT,
		CustomTimeFormat: "2006/01/02 15:04:05",
		Output:           os.Stdout,
	}))

	// Echo: Request ID
	e.Use(middleware.RequestID())

	// Custom: Custom Header
	e.Use(serverHeader)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "RoboBug/1.0")
		return next(c)
	}
}
