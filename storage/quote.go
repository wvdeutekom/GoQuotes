package storage

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"time"
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

// Quote struct for slash command
type Quote struct {
	Token       string `schema:"token" json:"token" gorethink:"token"`
	TeamID      string `schema:"team_id" json:"team_id" gorethink:"team_id"`
	TeamDomain  string `schema:"team_domain" json:"team_domain" gorethink:"team_domain"`
	ChannelID   string `schema:"channel_id" json:"channel_id" gorethink:"channel_id"`
	ChannelName string `schema:"channel_name" json:"channel_name" gorethink:"channel_name"`
	UserID      string `schema:"user_id" json:"user_id" gorethink:"user_id"`
	UserName    string `schema:"user_name" json:"user_name" gorethink:"user_name"`
	Text        string `schema:"text" json:"text" gorethink:"text"`
	Command     string `schema:"command" json:"command" gorethink:"command"`
	Timestamp   int32  `schema:"-" json:"-" gorethink:"timestamp"`
}

type QuoteStorage struct {
	Name    string
	URL     string
	Session *r.Session
}

func (s *QuoteStorage) SaveQuote(quote *Quote) {
	fmt.Printf("Looks like you're saving a quote: %#v\n", quote)

	quote.Timestamp = int32(time.Now().Unix())

	resp, err := r.DB(s.Name).Table("quote").Insert(quote).RunWrite(s.Session)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("get stuff! %#v\n", resp)
}

func (s *QuoteStorage) GetLatestQuote() Quote {

	rows, err := r.Table("quote").OrderBy(r.Desc("timestamp")).Run(s.Session)
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
