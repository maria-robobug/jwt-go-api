package handlers

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api/config"
)

func GetSecurityQuestions(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		log.Println("User was not valid")

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": config.GENERIC_FAILURE_MESSAGE,
		})
	}

	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if claims["aud"] != "localhost:3000" {
		log.Println("Audience was invalid!")

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": config.GENERIC_FAILURE_MESSAGE,
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"security_question": "Who is your favourite superhero?",
	})
}
