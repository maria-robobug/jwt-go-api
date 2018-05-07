package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
	// Echo: basic auth middleware
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Ideally you would check the DB here, but for simplicity just going to hardcode
		if username == "maria@email.com" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))
}
