package handlers

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/maria-robobug/jwt-go-api/src/api/models"
)

func GetLogin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// Ideally you would verify username and password in DB after hashing password
	if username == "maria" && password == "1234" {
		// create jwt token
		jwtToken, err := createLoginJwtToken()
		if err != nil {
			log.Println("Error creating JWT token", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Something went wrong...",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": jwtToken,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"error": "Your username or password was invalid.",
	})
}

func createLoginJwtToken() (string, error) {
	claims := models.JwtClaims{
		"maria",
		"18995",
		"/security/questions",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Second).Unix(),
			Issuer:    "localhost:8000",
			Audience:  "localhost:3000",
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
