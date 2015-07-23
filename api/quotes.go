package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
	"github.com/labstack/echo"
	"github.com/nlopes/slack"
	st "github.com/wvdeutekom/webhookproject/storage"
)

type Error struct {
	Message string
	Error   error
}

type Meta struct {
	Authors []string `json:"authors,omitempty"`
	Github  string   `json:"github,omitempty"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Meta   Meta        `json:"meta,omitempty"`
}

//POST /quotes
//POST /slack/insertQuote
func (a *AppContext) NewQuote(c *echo.Context) error {

	r := c.Request()

	//Add header for angular CORS support
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	//Parse post values
	r.ParseForm()
	isValid := len(r.Form["text"]) > 0 && len(r.Form["team_id"]) > 0
	if !isValid {
		log.Println("Invalid form (empty?)\nI'm a doctor Jim, not a magician!")
		return c.JSON(http.StatusBadRequest, "Looks like I'm missing some parameters, sir.")
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

	//Check for token header

	//Add header for angular CORS support
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	//Get quote from database
	if query != "" {
		//Seperate search terms and put them into a string array
		quotes, err = a.Storage.SearchQuotes(strings.Split(query, ","))
	} else {
		quotes, err = a.Storage.FindAllQuotes()
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Quotes could not be found.", err})
	}

	return c.JSON(http.StatusOK, FormatResponse("Fetched", quotes))
}

//GET /quotes/:id
func (a *AppContext) FindOneQuote(c *echo.Context) error {

	//Add header for angular CORS support
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	quote, err := a.Storage.FindOneQuote(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Quotes could not be found.", err})
	}

	return c.JSON(http.StatusOK, FormatResponse("Fetched", quote))
}

func (a *AppContext) EditQuote(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}

//DELETE /quotes/:id
func (a *AppContext) DeleteQuote(c *echo.Context) error {

	fmt.Printf("Delete api quote\n")
	quote, err := a.Storage.DeleteQuote(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, FormatResponse("Deleted", quote))
}

//GET /slack/searchQuote
func (a *AppContext) SearchQuote(c *echo.Context) error {

	r := c.Request()

	//Parse post values
	r.ParseForm()
	isValid := len(r.Form["text"]) > 0 && len(r.Form["team_id"]) > 0
	if !isValid {
		log.Println("Invalid form (empty?)\nI'm a doctor Jim, not a magician!")
		return c.JSON(http.StatusBadRequest, Error{"Invalid form, missing arguments", nil})
	}

	//Transfer post values to quote variable
	quote := new(st.Quote)
	decoder := schema.NewDecoder()

	if err := decoder.Decode(quote, c.Request().Form); err != nil {
		fmt.Println(err)
	}

	resultQuote, err := a.Storage.SearchQuotes(strings.Split(quote.Text, ","))

	if err != nil {
		return c.JSON(http.StatusNotFound, Error{"Quotes could not be found.", err})
	}

	quoteText := "\"" + resultQuote[0].Text + "\"" + " ~" + resultQuote[0].UserName

	params := slack.PostMessageParameters{}
	_, _, err = a.Slack.PostMessage(quote.ChannelID, quoteText, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{"Could not post to Slack channel", err})
	}

	return echo.NewHTTPError(http.StatusOK, "There's your quote sir!")
}

// Dev stuff
func (a *AppContext) SendQuote(c *echo.Context) error {
	fmt.Printf("Sending quote with slack: %s\n", a.Slack)

	params := slack.PostMessageParameters{}
	channelID, timestamp, err := a.Slack.PostMessage("C02R0LWRQ", "You are Zaphod Beeblebrox?", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return c.JSON(http.StatusInternalServerError, "Oh crap, something went wrong sending the quote\n")
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return c.JSON(http.StatusOK, "SendQuote fired, sir.")
}
