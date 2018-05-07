package main

import (
	"fmt"

	"github.com/maria-robobug/jwt-go-api/src/router"
)

func main() {
	fmt.Println("Welcome to the server!")
	e := router.New()

	e.Logger.Fatal(e.Start(":8000"))
}
