package api

import (
	"github.com/labstack/echo"
	"github.com/wvdeutekom/webhookproject/storage"
)

type Config struct {
	Port   int
	DbName string
	DbURL  string
	DbPort int
}

type AppContext struct {
	Config  *Config
	Storage *storage.QuoteStorage
}

// Struct used to marshall json
type Data struct {
	Object interface{}
	Meta   bool
}

// Initialize new configuration
func NewConfig() *Config {
	return &Config{
		Port:   8000,
		DbName: "quotes",
		DbURL:  "192.168.10.10",
		DbPort: 28015,
	}
}

func Route(e *echo.Echo, a *AppContext) {
	e.Post("/quote", a.newQuote)
	e.Post("/latestquote", a.GetLatestQuote)
	e.Post("/searchquote", a.SearchQuote)
}
