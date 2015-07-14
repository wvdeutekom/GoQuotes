package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/labstack/echo"
	st "github.com/wvdeutekom/webhookproject/storage"
)

type Error struct {
	Message string
}

type Meta struct {
	Total string
}

type Response struct {
	Username  string `json:"username,omitempty"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
}

func (a *AppContext) NewQuote(c *echo.Context) error {

	r := c.Request()

	//Parse post values
	r.ParseForm()
	isValid := len(r.Form["text"]) > 0 && len(r.Form["team_id"]) > 0
	if !isValid {
		log.Println("Invalid form (empty?)\nI'm a doctor Jim, not a magician!")
	}

	fmt.Printf("form:: %s\n", r.Form)

	//Transfer post values to quote variable
	quote := new(st.Quote)
	decoder := schema.NewDecoder()
	if err := decoder.Decode(quote, c.Request().PostForm); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Filled quote: %#v\n", quote)

	//Save the quote in the database
	a.Storage.SaveQuote(quote)
	resp := "Saving quote: " + quote.Text

	fmt.Println("\n\n")

	return c.JSON(http.StatusOK, resp)
}

//GET /quotes
func (a *AppContext) GetQuotes(c *echo.Context) error {
	var quotes []st.Quote
	var err error

	var query = c.Request.URL.Query().Get("q")

	//Get quote from database
	if query != "" {
		quotes, err = a.Storage.SearchQuotes(query)
	} else {
		quotes, err = a.Storage.FindAllQuotes()
	}

	fmt.Printf("GetQuotes: %s\n\n", quotes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Quotes could not be found."})
	}
	return c.JSON(http.StatusOK, quotes)
}

func (a *AppContext) FindOneQuote(c *echo.Context) error {

	quote, err := a.Storage.FindOneQuote(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Quotes could not be found."})
	}

	return c.JSON(http.StatusOK, quote)
}

func (a *AppContext) EditQuote(c *echo.Context) error {

	return c.JSON(http.StatusOK, "in development")
}

func (a *AppContext) DeleteQuote(c *echo.Context) error {

	return c.JSON(http.StatusOK, "in development")
}

func (a *AppContext) SearchQuote(c *echo.Context) error {

	r := c.Request()

	//Parse post values
	r.ParseForm()
	isValid := len(r.Form["text"]) > 0 && len(r.Form["team_id"]) > 0
	if !isValid {
		log.Println("Invalid form (empty?)\nI'm a doctor Jim, not a magician!")
	}

	fmt.Printf("form:: %s\n", r.Form)

	//Transfer post values to quote variable
	quote := new(st.Quote)
	decoder := schema.NewDecoder()

	var err interface{}
	if r.Method == "GET" {
		err = decoder.Decode(quote, c.Request().Form)
	} else {
		err = decoder.Decode(quote, c.Request().PostForm)
	}

	if err != nil {
		fmt.Println(err)
		//log.Printf("error %s", string.err.Error)
	}
	fmt.Printf("Filled quote: %#v\n", quote)

	resultQuote := a.Storage.SearchQuote(quote.Text)

	resp := Response{
		Username: resultQuote.UserName,
		Text:     resultQuote.Text,
	}

	return c.JSON(http.StatusOK, resp)
}

// Dev stuff
func (a *AppContext) SendQuote(c *echo.Context) error {
	fmt.Printf("Sending quote with slack: %s\n", a.Slack)
	if err := a.Slack.ChatPostMessage("C02QG1PDQ", "Karlo, of jij even je mondtd wilt houden.", nil); err != nil {
		fmt.Printf("Error sending quote: %s\n", err)
	}

	return c.JSON(http.StatusOK, "SendQuote fired, sir.")
}
