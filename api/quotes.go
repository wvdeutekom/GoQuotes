package api

import (
	"fmt"
	"io/ioutil"
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
	Token       string `schema:"token"`
	TeamID      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelID   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	Timestamp   int    `schema:"timestamp"`
	UserID      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Text        string `schema:"text"`
	TriggerWord string `schema:"trigger_word"`
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

	body, _ := ioutil.ReadAll(c.Request().Body)
	fmt.Println("response Body:", string(body))

	quote := new(Quote)
	decoder := schema.NewDecoder()
	//r.PostForm is a map of our POST form values
	err := decoder.Decode(quote, c.Request().PostForm)

	log.Printf("%#v\n", quote)

	if err != nil {
		fmt.Println(err)
		//log.Printf("error %s", string.err.Error)
		// Handle error
	}

	//c.Request().ParseForm()
	//newQ := Quote{
	//	Token:  c.Request().Form.Get("token"),
	//	TeamID: c.Request().Form.Get("team_id"),
	//	Token:  c.Request().Form.Get("team_domain"),
	//	Token:  c.Request().Form.Get("channel_id"),
	//	Token:  c.Request().Form.Get("channel_name"),
	//	Token:  c.Request().Form.Get("timestamp"),
	//	Token:  c.Request().Form.Get("user_id"),
	//	Token:  c.Request().Form.Get("user_name"),
	//	Token:  c.Request().Form.Get("text"),
	//	Token:  c.Request().Form.Get("trigger_word"),
	//}

	//json.Unmarshal([]byte(body), &quote)
	//log.Printf("%#v\n", quote)

	//uID := "U02QG1PBN"  //wijnand
	//chID := "C02QG1PDQ" //Generaal

	//quote := Quote{
	//	Token:       "lXofu0WHVRJpB0P1efGSNVyl",
	//	TeamID:      "teamid931",
	//	TeamDomain:  "itlounge",
	//	ChannelID:   chID,
	//	ChannelName: "generaal",
	//	Timestamp:   1435839624,
	//	UserID:      uID,
	//	UserName:    "wijnand",
	//	Text:        "JA HALLO!",
	//	TriggerWord: "quote",
	//}
	//resp := QuoteJSON{
	//	Data: quote,
	//}

	//bytes, _ := json.Marshal(quote)

	resp := Response{
		Username: "Gopher :)",
		Text:     "JA HALLO!",
	}

	return c.JSON(http.StatusOK, resp)
	//return c.String(http.StatusOK, "looks like a new quote to me!")
}
