package storage

type Message struct {
	ID          string `schema:"id" json:"id" gorethink:"id,omitempty" mapstructure:"id"`
	Token       string `schema:"token" json:"-" gorethink:"token" mapstructure:"token"`
	TeamID      string `schema:"team_id" json:"team_id,omitempty" gorethink:"team_id" mapstructure:"team_id"`
	TeamDomain  string `schema:"team_domain" json:"team_domain,omitempty" gorethink:"team_domain" mapstructure:"team_domain"`
	ChannelID   string `schema:"channel_id" json:"channel_id,omitempty" gorethink:"channel_id" mapstructure:"channel_id"`
	ChannelName string `schema:"channel_name" json:"channel_name,omitempty" gorethink:"channel_name" mapstructure:"channel_name"`
	UserID      string `schema:"user_id" json:"user_id,omitempty" gorethink:"user_id" mapstructure"user_id"`
	UserName    string `schema:"user_name" json:"user_name,omitempty" gorethink:"user_name" mapstructure:"user_name"`
	Text        string `schema:"text" json:"text" gorethink:"text" mapstructure:"text"`
	Command     string `schema:"command" json:"command,omitempty" gorethink:"command" mapstructure:"command"`
	Timestamp   int    `schema:"-" json:"timestamp,omitempty" gorethink:"timestamp" mapstructure:"timestamp"`
}
