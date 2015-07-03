package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/labstack/echo"
)

type Meta struct {
	Total string
}

type QuoteJSON struct {
	Data Quote `json:"data"`
	Meta *Meta `json:"meta,omitempty"`
}

type Quote struct {
	Token       string  `schema:"token"`
	TeamID      string  `schema:"team_id"`
	TeamDomain  string  `schema:"team_domain"`
	ChannelID   string  `schema:"channel_id"`
	ChannelName string  `schema:"channel_name"`
	Timestamp   float32 `schema:"timestamp"`
	UserID      string  `schema:"user_id"`
	UserName    string  `schema:"user_name"`
	Text        string  `schema:"text"`
	Command     string  `schema:"command"`
	ServiceID   string  `schema:"service_id"`
}

// type Quote struct {
// 	Token       string `json:"token"`
// 	TeamID      string `json:"team_id"`
// 	TeamDomain  string `json:"team_domain"`
// 	ChannelID   string `json:"channel_id"`
// 	ChannelName string `json:"channel_name"`
// 	Timestamp   int    `json:"timestamp"`
// 	UserID      string `json:"user_id"`
// 	UserName    string `json:"user_name"`
// 	Text        string `json:"text"`
// 	TriggerWord string `json:"trigger_word"`
// }

type Response struct {
	Username string `json:"username,omitempty"`
	Text     string `json:"text"`
}

func (a *AppContext) newQuote(c *echo.Context) error {

	r := c.Request()

	r.ParseForm()
	isValid := len(r.Form["text"]) > 0 && len(r.Form["team_id"]) > 0
	if !isValid {
		log.Println("Invalid form (empty?)\nI'm a doctor Jim, not a magician!")
	}

	fmt.Printf("form:: %s\n", r.Form)

	quote := new(Quote)
	decoder := schema.NewDecoder()
	err := decoder.Decode(quote, c.Request().PostForm)

	if err != nil {
		fmt.Println(err)
		//log.Printf("error %s", string.err.Error)
	}
	fmt.Printf("Filled quote: %#v\n", quote)

	//resp := Response{
	//	Username: quote.UserName,
	//	Text:     "Saving quote: " + quote.Text,
	//}

	resp := "Saving quote: " + quote.Text

	fmt.Println("\n\n")

	return c.JSON(http.StatusOK, resp)
	//return c.String(http.StatusOK, "looks like a new quote to me!")
}
