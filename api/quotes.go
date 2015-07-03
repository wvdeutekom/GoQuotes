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

	r := c.Request()

	r.ParseMultipartForm(5120)
	isValid := len(r.Form["text"]) > 0 && len(r.Form["team_id"]) > 0
	if !isValid {
		log.Println("see ya homies")
	}
	//	isCmd := len(r.Form["trigger_word"]) > 0

	fmt.Printf("TRIGGERRR: %s", r.Form["trigger_word"][0])

	body, _ := ioutil.ReadAll(c.Request().Body)
	contentType := c.Request().Header.Get("Content-Type")

	log.Printf("form parsed1: ", r.Form)
	r.ParseForm()
	log.Printf("form parsed2: ", r.Form)
	log.Printf("form value§: ", r.FormValue("token"))
	log.Printf("form value§: ", r.PostFormValue("token"))

	x := r.Form.Get("token")
	fmt.Printf("param_name: ", x)

	fmt.Printf("request method: %s\n", c.Request().Method)
	fmt.Printf("formvalue: %s\n", c.Request().FormValue("token"))
	fmt.Printf("content-type: %s\n", contentType)
	fmt.Printf("response Body: %s\n", string(body))
	fmt.Printf("Post form get token: %s\n", c.Request().PostForm.Get("token"))
	fmt.Printf("Post form get token222: %s\n", c.Request().PostFormValue("token"))

	// var u Quote

	// defer c.Request().Body.Close()

	// if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
	// 	fmt.Printf("json decoded inside if: %#v\n", u)
	// }

	// fmt.Printf("json decoded: %#v\n", u)

	// d := form.NewDecoder(c.Request().Body)
	// if err := d.Decode(&u); err != nil {
	// 	fmt.Printf("ERROR JIM! ERROR!: ", err)
	// }

	// fmt.Printf("Decoded: %#v", u)

	// payload := make(map[string]interface{})

	// if strings.Contains(contentType, "form") {
	// 	fmt.Println("Looks like youre sending a form")

	// 	formData, err := url.ParseQuery(string(body))
	// 	if err != nil {
	// 		log.Printf("error parsing form payload %+v\n", err)
	// 	} else {
	// 		payload = valuesToMap(formData)
	// 	}
	// }

	// fmt.Printf("payload payload paylod: %#v\n", payload)

	quote := new(Quote)
	decoder := schema.NewDecoder()
	err := decoder.Decode(quote, c.Request().PostForm)

	if err != nil {
		fmt.Println(err)
		//log.Printf("error %s", string.err.Error)
		// Handle error
	}

	fmt.Printf("THIS ISNT EVEN MY FINAL FORM: %#v\n", quote)
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

	fmt.Println("\n\n\n")

	return c.JSON(http.StatusOK, resp)
	//return c.String(http.StatusOK, "looks like a new quote to me!")
}

func valuesToMap(values map[string][]string) map[string]interface{} {
	ret := make(map[string]interface{})

	for key, value := range values {
		if len(value) > 0 {
			ret[key] = value[0]
		}
	}

	return ret
}
