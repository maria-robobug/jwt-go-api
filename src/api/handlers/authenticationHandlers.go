package handlers

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	NextStep string `json:"next_step"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// Ideally you would verify username and password in DB after hashing password
	if username == "maria" && password == "1234" {
		// create jwt token
		jwtToken, err := createJwtToken()
		if err != nil {
			log.Println("Error creating JWT token", err)
			return c.String(http.StatusInternalServerError, "Something went wrong!")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": jwtToken,
		})
	}

	return c.String(http.StatusUnauthorized, "Your username or password was invalid.")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"maria",
		"18995",
		"/security/questions",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Second).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
