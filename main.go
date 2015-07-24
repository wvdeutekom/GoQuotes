package main

import (
	"fmt"
	"log"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/nlopes/slack"
	a "github.com/wvdeutekom/webhookproject/api"
	"github.com/wvdeutekom/webhookproject/storage"
)

func main() {

	config := a.NewConfig("config.gcfg")
	echo := echo.New()
	s := slack.New(config.Slack.Token)

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

	_, err = r.DB(config.Database.Name).TableCreate("quotes").RunWrite(session)
	if err != nil {
		fmt.Print(err)
	}
	_, err = r.DB(config.Database.Name).TableCreate("activities").RunWrite(session)
	if err != nil {
		fmt.Print(err)
	}

	// Middleware
	echo.Use(mw.Logger())
	echo.Use(mw.Recover())

	appcontext := &a.AppContext{
		Slack:  s,
		Config: config,
		Storage: &storage.Storage{
			Name:    "quotes",
			URL:     "192.168.10.10",
			Session: session,
		},
	}

	go appcontext.Monitor()

	//Routes
	a.Route(echo, appcontext)

	addr := fmt.Sprintf(":%d", config.App.Port)
	log.Printf("Starting server on: %s", addr)
	echo.Run(addr)
}
