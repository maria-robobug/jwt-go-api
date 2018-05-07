package handlers

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:"name"`
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

		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = jwtToken
		jwtCookie.Secure = false
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"token": jwtToken,
		})
	}

	return c.String(http.StatusUnauthorized, "Your username or password was invalid.")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"maria",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
