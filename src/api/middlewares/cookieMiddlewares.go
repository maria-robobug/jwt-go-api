package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(validateCookie)
}

// Custom middleware
func validateCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")

		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you do not have any cookie")
			} else {
				log.Println(err)
				return err
			}
		}

		// In reality you would go in the session store to compare
		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you do not have the right cookie, cookie")
	}
}
