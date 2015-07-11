package api

import (
	"fmt"
	"os"

	"code.google.com/p/gcfg"
	"github.com/labstack/echo"
	"github.com/wvdeutekom/webhookproject/slack"
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
	App struct {
		Port int
	}
	Database struct {
		Name string
		URL  string
		Port int
	}
	Slack struct {
		Token string
	}
}

type AppContext struct {
	Slack   *slack.Slack
	Config  *Config
	Storage *storage.QuoteStorage
}

// Struct used to marshall json
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
	fmt.Printf("loading config: %s\n", configFile)

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

func Route(e *echo.Echo, a *AppContext) {
	e.Post("/quotes", a.newQuote)
	e.Post("/latestquote", a.GetLatestQuote)
	e.Post("/searchquote", a.SearchQuote)
	e.Get("/searchquote", a.SearchQuote)
	e.Get("/debug", a.SendQuote)
}
