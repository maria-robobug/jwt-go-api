package main

import (
	"github.com/maria-robobug/jwt-go-api/src/router"
)

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
