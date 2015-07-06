package storage

import (
	"fmt"

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

// Quote struct for slash command
type Quote struct {
	Token       string `schema:"token"`
	TeamID      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelID   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	UserID      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Text        string `schema:"text"`
	Command     string `schema:"command"`
}

type QuoteStorage struct {
	Name    string
	URL     string
	Session *r.Session
}

func (s *QuoteStorage) SaveQuote(quote *Quote) {
	fmt.Printf("Looks like you're saving a quote: %#v\n", quote)

	resp, err := r.DB(s.Name).Table("quote").Insert(quote).RunWrite(s.Session)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("get stuff! %#v\n", resp)
}
