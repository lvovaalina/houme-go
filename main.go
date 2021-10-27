package main

import (
	"os"

	_ "github.com/heroku/x/hmetrics/onload"

	"bitbucket.org/houmeteam/houme-go/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	router.RouterInitialize(port)
}
