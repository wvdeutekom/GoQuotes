package storage

import (
	"fmt"
	"strings"
	"time"

	r "github.com/dancannon/gorethink"
)

// Quote struct for outgoing webhook instead of slash command
// type Quote struct {
// 	Token       string `schema:"token"`
// 	TeamID      string `schema:"team_id"`
// 	TeamDomain  string `schema:"team_domain"`
// 	ChannelID   string `schema:"channel_id"`
// 	ChannelName string `schema:"channel_name"`
// 	Timestamp   int    `schema:"timestamp"`
// 	UserID      string `schema:"user_id"`
// 	UserName    string `schema:"user_name"`
// 	Text        string `schema:"text"`
// 	TriggerWord string `schema:"trigger_word"`
// }

type Storage struct {
	Name string
	URL  string
	Port int
}

// Quote struct for slash command
type Quote struct {
	ID          string `schema:"id" json:"id" gorethink:"id,omitempty"`
	Token       string `schema:"token" json:"-" gorethink:"token"`
	TeamID      string `schema:"team_id" json:"team_id,omitempty" gorethink:"team_id"`
	TeamDomain  string `schema:"team_domain" json:"team_domain,omitempty" gorethink:"team_domain"`
	ChannelID   string `schema:"channel_id" json:"channel_id,omitempty" gorethink:"channel_id"`
	ChannelName string `schema:"channel_name" json:"channel_name,omitempty" gorethink:"channel_name"`
	UserID      string `schema:"user_id" json:"user_id,omitempty" gorethink:"user_id"`
	UserName    string `schema:"user_name" json:"user_name,omitempty" gorethink:"user_name"`
	Text        string `schema:"text" json:"text,omitempty" gorethink:"text"`
	Command     string `schema:"command" json:"command,omitempty" gorethink:"command"`
	Timestamp   int    `schema:"-" json:"timestamp,omitempty" gorethink:"timestamp"`
}

type QuoteStorage struct {
	Name    string
	URL     string
	Session *r.Session
}

func (s *QuoteStorage) SaveQuote(quote *Quote) {

	fmt.Printf("\n\nLooks like you're saving a quote: %#v\n\n", quote)

	if quote.Timestamp == 0 {
		quote.Timestamp = int(time.Now().Unix())
	}
	_, err := r.DB(s.Name).Table("quotes").Insert(quote).RunWrite(s.Session)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func (s *QuoteStorage) FindAllQuotes() ([]Quote, error) {
	rows, err := r.DB(s.Name).Table("quotes").Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Quote
	err = rows.All(&list)
	if err == r.ErrEmptyResult {
		return nil, err
	}

	return list, nil
}

func (s *QuoteStorage) FindOneQuote(id string) (*Quote, error) {

	rows, err := r.DB(s.Name).Table("quotes").Filter(
		r.Row.Field("id").Eq(id)).Run(s.Session)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quote Quote
	err = rows.One(&quote)
	if err == r.ErrEmptyResult {
		return nil, err
	}

	return &quote, nil
}

func (s *QuoteStorage) GetLatestQuote() Quote {

	rows, err := r.Table("quotes").OrderBy(r.Desc("timestamp")).Run(s.Session)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var quote Quote
	err2 := rows.One(&quote)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("Fetch one record %#v\n", quote)

	return quote
}

func (s *QuoteStorage) DeleteQuote(id string) (*Quote, error) {

	rows, err := r.DB(s.Name).Table("quotes").Get(id).Delete().Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quote Quote
	err = rows.One(&quote)
	if err == r.ErrEmptyResult {
		return nil, err
	}

	return &quote, nil

}

func (s *QuoteStorage) SearchQuotes(searchStrings []string) ([]Quote, error) {

	fmt.Printf("Searchterms: %s\n", searchStrings)

	//Append the strings into one regex string, e.g. bob|said|bananas
	searchTerms := strings.Join(searchStrings, "|")
	fmt.Printf("Filtered searchterms: %s\n", searchTerms)

	rows, err := r.Table("quotes").Filter(func(quote r.Term) r.Term {
		return quote.Field("text").Match(searchTerms)
	}).Run(s.Session)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []Quote{}
	err2 := rows.All(&list)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("Search result record %#v\n", list)

	return list, nil
}
