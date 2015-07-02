package main

import (
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/wvdeutekom/webhookproject/api"
	"log"
)

func main() {
	config := api.NewConfig()
	echo := echo.New()

	// Middleware
	echo.Use(mw.Logger())
	echo.Use(mw.Recover())

	appcontext := &api.AppContext{
		Config: config,
	}

	//Routes
	api.Route(echo, appcontext)

	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("Starting server on: %s", addr)
	echo.Run(addr)
}
