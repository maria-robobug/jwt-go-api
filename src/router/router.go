package router

import (
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api"
	"github.com/maria-robobug/jwt-go-api/src/api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	securityGroup := e.Group("/security")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetJwtMiddlewares(securityGroup)

	// unathenticated routes
	api.MainGroup(e)

	// authenticated routes
	api.SecurityGroup(securityGroup)

	return e
}
