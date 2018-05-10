package api

import (
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api/handlers"
)

func UnauthenticatedGroup(e *echo.Echo) {
	e.GET("/", handlers.Root)
	e.GET("/login", handlers.Login)
}
