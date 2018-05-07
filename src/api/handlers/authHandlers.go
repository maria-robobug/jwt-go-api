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
		// You could also write this as:
		// cookie := new(http.Cookie)
		cookie := &http.Cookie{}
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		// cookie.Secure = true
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		// create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Error creating JWT token", err)
			return c.String(http.StatusInternalServerError, "Something went wrong!")
		}

		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You are logged in!",
			"token":   token,
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

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
