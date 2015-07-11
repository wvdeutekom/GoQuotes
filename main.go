package main

import (
	"fmt"
	"log"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/wvdeutekom/webhookproject/api"
	"github.com/wvdeutekom/webhookproject/storage"
)

func main() {

	config := api.NewConfig("config.gcfg")
	echo := echo.New()
	fmt.Printf("MY CONFIG YEAAHHH: %#v\n", config)

	var session *r.Session

	session, err := r.Connect(r.ConnectOpts{
		Address:  fmt.Sprint(config.Database.URL, ":", config.Database.Port),
		Database: config.Database.Name,
		MaxIdle:  10,
		MaxOpen:  10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	session.SetMaxOpenConns(5)

	r.DBCreate(config.Database.Name).Exec(session)
	if err != nil {
		log.Println(err)
	}

	_, err = r.DB(config.Database.Name).TableCreate("quote").RunWrite(session)
	if err != nil {
		fmt.Print(err)
	}

	// Middleware
	echo.Use(mw.Logger())
	echo.Use(mw.Recover())

	appcontext := &api.AppContext{
		Config: config,
		Storage: &storage.QuoteStorage{
			Name:    "quotes",
			URL:     "192.168.10.10",
			Session: session,
		},
	}

	//Routes
	api.Route(echo, appcontext)

	addr := fmt.Sprintf(":%d", config.App.Port)
	log.Printf("Starting server on: %s", addr)
	echo.Run(addr)
}
