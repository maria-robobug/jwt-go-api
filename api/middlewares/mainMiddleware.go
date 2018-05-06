package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddleware(e *echo.Echo) {
	// Logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${status} ${host}${path} ${latency_human}]` + "\n",
	}))
}
