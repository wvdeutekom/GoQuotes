package main

import (
	"fmt"
	"log"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/wvdeutekom/webhookproject/api"
)

func main() {
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

	// resp, err := r.DB(config.DbName).TableCreate("quote").RunWrite(session)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	resp, err := r.DB(config.DbName).Table("quote").Insert(map[string]interface{}{
		"title":   "Lorem ipsum",
		"content": "Dolor sit amet",
	}).RunWrite(session)
	if err != nil {
		fmt.Print(err)
		return
	}

	resp, err = r.DB(config.DbName).Table("quote").RunWrite(session)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("get stuff! %#v\n", resp)

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
