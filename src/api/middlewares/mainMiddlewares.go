package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	// Echo: Logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${status} ${host}${path} ${latency_human}]` + "\n",
	}))

	e.Use(serverHeader)
}

// Custom header middleware
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "RoboBug/1.0")
		return next(c)
	}
}
