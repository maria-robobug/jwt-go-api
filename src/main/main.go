package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand is type: %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "invalid request",
	})
}

// Fastest
func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to read the request body for addCat: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed to unmarshal in addCat: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "we got your cat!")
}

// Close to addCats in terms of performance (probably most ideal)
func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got your dog!")
}

// Slowest, and adds dependency of echo
func addHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster!")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you have found the secret admin page, welcome!")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret cookie page!")
}

func login(c echo.Context) error {
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
		return c.String(http.StatusOK, "You are logged in!")
	}

	return c.String(http.StatusUnauthorized, "Your username or password was invalid.")
}

/////// middlewares ///////
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		return next(c)
	}
}

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

func main() {
	fmt.Println("Welcome to the server!")

	e := echo.New()
	e.Use(ServerHeader)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${status} ${host}${path} ${latency_human}]` + "\n",
	}))

	e.GET("/", helloWorld)
	e.GET("/cats/:data", getCats)
	e.GET("/login", login)

	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hamsters", addHamster)

	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")

	// Adds basic auth to admin group
	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Ideally you would check the DB here, but for simplicity just going to hardcode
		if username == "maria@email.com" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	cookieGroup.Use(validateCookie)

	cookieGroup.GET("/main", mainCookie)
	adminGroup.GET("/main", mainAdmin)

	e.Logger.Fatal(e.Start(":8000"))
}
