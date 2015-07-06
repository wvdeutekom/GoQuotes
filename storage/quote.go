package storage

import "fmt"

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

func SaveQuote(quote *Quote) {
	fmt.Printf("Looks like you're saving a quote: %#v\n", quote)
}
