package router

import (
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddleware(e)
	middlewares.SetAdminMiddleware(adminGroup)

	return e
}
