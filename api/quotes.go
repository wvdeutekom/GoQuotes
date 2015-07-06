package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/labstack/echo"
	st "github.com/wvdeutekom/webhookproject/storage"
)

type Meta struct {
	Total string
}

// type Response struct {
// 	Username string `json:"username,omitempty"`
// 	Text     string `json:"text"`
// }

func (a *AppContext) newQuote(c *echo.Context) error {

	r := c.Request()

	r.ParseForm()
	isValid := len(r.Form["text"]) > 0 && len(r.Form["team_id"]) > 0
	if !isValid {
		log.Println("Invalid form (empty?)\nI'm a doctor Jim, not a magician!")
	}

	fmt.Printf("form:: %s\n", r.Form)

	quote := new(st.Quote)
	decoder := schema.NewDecoder()
	err := decoder.Decode(quote, c.Request().PostForm)

	if err != nil {
		fmt.Println(err)
		//log.Printf("error %s", string.err.Error)
	}
	fmt.Printf("Filled quote: %#v\n", quote)

	a.Storage.SaveQuote(quote)
	resp := "Saving quote: " + quote.Text

	fmt.Println("\n\n")

	return c.JSON(http.StatusOK, resp)
}
