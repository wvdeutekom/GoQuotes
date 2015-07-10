package api

import (
	"github.com/labstack/echo"
	"github.com/wvdeutekom/webhookproject/slack"
	"github.com/wvdeutekom/webhookproject/storage"
)

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
func NewConfig() string {
	return `
	[app]
	port = 8000

	[database]
	name = quotes
	url = 192.168.10.10
	port = 28015

	[slack]
	apitoken = xoxb-7510246197-X36KpkPgHgUFTDeIsPXAzcv2 
	`
}

func Route(e *echo.Echo, a *AppContext) {
	e.Post("/quotes", a.newQuote)
	e.Post("/latestquote", a.GetLatestQuote)
	e.Post("/searchquote", a.SearchQuote)
	e.Get("/searchquote", a.SearchQuote)
	e.Get("/debug", a.SendQuote)
}
