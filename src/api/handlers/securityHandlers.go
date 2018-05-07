package handlers

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func SecurityQuestions(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["username"], "User ID: ", claims["user_id"])

	return c.String(http.StatusOK, "you are on the security questions page!")
}
