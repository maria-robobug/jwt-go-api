package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Root(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}
