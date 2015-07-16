package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
	"github.com/labstack/echo"
	st "github.com/wvdeutekom/webhookproject/storage"
)

type Error struct {
	Message string
	Error   error
}

type Meta struct {
	Total string
}

type Response struct {
	Username  string `json:"username,omitempty"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
}

//POST /quotes
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

	var query = c.Request().URL.Query().Get("q")

	fmt.Printf("GET PARAMS: %#v\n\n", query)

	//Get quote from database
	if query != "" {
		//Seperate search terms and put them into a string array
		quotes, err = a.Storage.SearchQuotes(strings.Split(query, ","))
	} else {
		quotes, err = a.Storage.FindAllQuotes()
	}

	fmt.Printf("GetQuotes: %s\n\n", quotes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Quotes could not be found.", err})
	}
	return c.JSON(http.StatusOK, quotes)
}

//GET /quotes/:id
func (a *AppContext) FindOneQuote(c *echo.Context) error {

	quote, err := a.Storage.FindOneQuote(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Quotes could not be found.", err})
	}

	return c.JSON(http.StatusOK, quote)
}

func (a *AppContext) EditQuote(c *echo.Context) error {

	return c.JSON(http.StatusOK, "in development")
}

//DELETE /quotes/:id
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
		return c.JSON(http.StatusBadRequest, Error{"Invalid form, missing arguments", nil})
	}

	fmt.Printf("form:: %s\n", r.Form)

	//Transfer post values to quote variable
	quote := new(st.Quote)
	decoder := schema.NewDecoder()

	if err := decoder.Decode(quote, c.Request().Form); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Filled quote: %#v\n", quote)

	resultQuote, err := a.Storage.SearchQuotes(strings.Split(quote.Text, ","))

	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Quotes could not be found.", err})
	}

	quoteText := "\"" + resultQuote[0].Text + "\"" + " ~" + resultQuote[0].UserName
	if a.Slack.ChatPostMessage(quote.ChannelID, quoteText, nil); err != nil {
		fmt.Printf("Error sending quote: %s\n", err)
		return c.JSON(http.StatusBadRequest, Error{"Could not post to Slack channel", err})
	}

	return echo.NewHTTPError(http.StatusOK, "There's your quote sir!")
}

// Dev stuff
func (a *AppContext) SendQuote(c *echo.Context) error {
	fmt.Printf("Sending quote with slack: %s\n", a.Slack)
	if err := a.Slack.ChatPostMessage("C02QG1PDQ", "Karlo, of jij even je mondtd wilt houden.", nil); err != nil {
		fmt.Printf("Error sending quote: %s\n", err)
	}

	return c.JSON(http.StatusOK, "SendQuote fired, sir.")
}
