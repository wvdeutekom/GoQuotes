package api

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Config struct {
	Port   int
	DbName string
	DbURL  string
	DbPort int
}

type AppContext struct {
	Config *Config
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

// Handler
func hello(c *echo.Context) error {
	log.Printf("%#v\n", c)
	return c.String(http.StatusOK, "Hello, ECHO!")
}

func Route(e *echo.Echo, a *AppContext) {
	e.Get("/", hello)
	e.Post("/quote", a.newQuote)
	e.Get("/quote", a.newQuote)
}
