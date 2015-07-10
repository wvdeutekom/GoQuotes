package main

import (
	"fmt"
	"log"

	"code.google.com/p/gcfg"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/wvdeutekom/webhookproject/api"
	"github.com/wvdeutekom/webhookproject/storage"
)

type ConfigFile struct {
	Slack struct {
		ApiToken string
	}
}

func main() {

	var cfg Config
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("MY CONFIG YEAAHHH: %#v\n", cfg)

	config := api.NewConfig()
	echo := echo.New()

	var session *r.Session

	session, err := r.Connect(r.ConnectOpts{
		Address:  fmt.Sprint(config.DbURL, ":", config.DbPort),
		Database: config.DbName,
		MaxIdle:  10,
		MaxOpen:  10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	session.SetMaxOpenConns(5)

	r.DBCreate(config.DbName).Exec(session)
	if err != nil {
		log.Println(err)
	}

	_, err = r.DB(config.DbName).TableCreate("quote").RunWrite(session)
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

	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("Starting server on: %s", addr)
	echo.Run(addr)
}

func LoadConfiguration()
