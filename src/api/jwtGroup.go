package api

import (
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api/handlers"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
