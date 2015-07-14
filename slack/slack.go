package slack

import (
	"net/url"
)

const apiBaseUrl = "https://slack.com/api/"

type Slack struct {
	Token string
}

func New(token string) *Slack {
	return &Slack{
		Token: token,
	}
}

func (sl *Slack) UrlValues() *url.Values {
	uv := url.Values{}
	uv.Add("token", sl.Token)
	return &uv
}
