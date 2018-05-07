package api

import (
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api/handlers"
)

func SecurityGroup(g *echo.Group) {
	g.GET("/questions", handlers.SecurityQuestions)
}
