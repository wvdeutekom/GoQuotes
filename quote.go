package main

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
