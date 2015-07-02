package api

import (
	//"encoding/json"
	"net/http"

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
	Token       string `json:"token"`
	TeamID      string `json:"team_id"`
	TeamDomain  string `json:"team_domain"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Timestamp   int    `json:"timestamp"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Text        string `json:"text"`
	TriggerWord string `json:"trigger_word"`
}

type Response struct {
	Text string `json:"text"`
}

func (a *AppContext) newQuote(c *echo.Context) error {

	uID := "U02QG1PBN"  //wijnand
	chID := "C02QG1PDQ" //Generaal

	quote := Quote{
		Token:       "lXofu0WHVRJpB0P1efGSNVyl",
		TeamID:      "teamid931",
		TeamDomain:  "itlounge",
		ChannelID:   chID,
		ChannelName: "generaal",
		Timestamp:   1435839624,
		UserID:      uID,
		UserName:    "wijnand",
		Text:        "JA HALLO!",
		TriggerWord: "quote",
	}
	resp := QuoteJSON{
		Data: quote,
	}

	//bytes, _ := json.Marshal(quote)

	return c.JSON(http.StatusOK, resp)
	//return c.String(http.StatusOK, "looks like a new quote to me!")
}
