package api

import (
	"fmt"
	"os"

	"code.google.com/p/gcfg"
	"github.com/labstack/echo"
	"github.com/nlopes/slack"
	"github.com/wvdeutekom/webhookproject/storage"
)

const defaultConfig = `
[app]
port = 8000

[database]
name = quotes
url = 192.168.10.10
port = 28015

[slack]
token = nicetry 
`

type Config struct {
	App      App
	Database storage.Database
	Slack    Slack
}

type App struct {
	Port  int
	Token string
}

type AppContext struct {
	Slack   *slack.Slack
	Config  *Config
	Storage *storage.Storage
}

type Slack struct {
	Token string
}

type Data struct {
	Object interface{}
	Meta   bool
}

// Initialize new configuration
func NewConfig(configFile string) *Config {

	//Check if config file is present, else load default config
	if _, err := os.Stat(configFile); err == nil {
		return LoadConfig(configFile)
	} else {
		return LoadConfig("")
	}
}

func LoadConfig(configFile string) *Config {

	var cfg Config
	var err error

	if configFile != "" {
		err = gcfg.ReadFileInto(&cfg, configFile)
	} else {
		err = gcfg.ReadStringInto(&cfg, defaultConfig)
	}

	if err != nil {
		fmt.Printf("Reading config error: %s\n", err)
	}

	return &cfg
}

func SetDefaultHeaders(c *echo.Context) {

	//Add header for angular CORS support
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
}

func FormatResponse(status string, data interface{}) Response {

	return Response{
		Status: status,
		Data:   data,
		Meta: Meta{
			Authors: []string{"Wijnand van Deutekom"},
			Github:  "https://github.com/wvdeutekom/GoQuotes",
		},
	}
}

func Route(e *echo.Echo, a *AppContext) {
	//Quotes
	e.Post("/quotes", a.NewQuote)
	e.Get("/quotes", a.GetQuotes)
	e.Get("/quotes/:id", a.FindOneQuote)
	e.Put("/quotes/:id", a.EditQuote)
	e.Delete("/quotes/:id", a.DeleteQuote)

	//Slack specific api calls, uses incoming x-www-form-urlencoded post data instead of json
	e.Post("/slack/insertQuote", a.NewQuote)
	e.Get("/slack/searchQuote", a.SearchQuote)

	//Activity feed
	e.Post("/activities", a.NewActivity)
	e.Get("/activities", a.GetActivities)
	e.Get("/activities/:id", a.FindOneActivity)
	e.Delete("/activities/:id", a.DeleteActivity)

	//Debug
	e.Get("/debug", a.SendQuote)
}
