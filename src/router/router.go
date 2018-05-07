package router

import (
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api"
	"github.com/maria-robobug/jwt-go-api/src/api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	securityGroup := e.Group("/security")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(securityGroup)

	// unathenticated routes
	api.MainGroup(e)

	// authenticated routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.SecurityGroup(securityGroup)

	return e
}
