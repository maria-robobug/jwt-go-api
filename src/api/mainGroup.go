package api

import (
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Root)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)

	e.POST("/dogs", handlers.AddDog)

	e.POST("/hamsters", handlers.AddHamster)

	e.GET("/login", handlers.Login)
}
