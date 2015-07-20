package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nlopes/slack"
	"github.com/wvdeutekom/webhookproject/storage"
)

func (a *AppContext) Monitor() {

	s := a.Slack

	chSender := make(chan slack.OutgoingMessage)
	chReceiver := make(chan slack.SlackEvent)

	s.SetDebug(true)
	wsAPI, err := s.StartRTM("", "https://slack.com/api/rtm.start")
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	go wsAPI.HandleIncomingEvents(chReceiver)
	go wsAPI.Keepalive(20 * time.Second)
	go func(wsAPI *slack.SlackWS, chSender chan slack.OutgoingMessage) {
		for {
			select {
			case msg := <-chSender:
				wsAPI.SendMessage(&msg)
			}
		}
	}(wsAPI, chSender)
	for {
		select {
		case msg := <-chReceiver:
			fmt.Print("Event Received: ")
			switch msg.Data.(type) {
			case slack.HelloEvent:
				//Ignore hello
			case *slack.MessageEvent:
				event := msg.Data.(*slack.MessageEvent)
				fmt.Printf("Message: %v\n", event)
			case *slack.PresenceChangeEvent:
				event := msg.Data.(*slack.PresenceChangeEvent)
				fmt.Printf("Presence Change: %v\n", event)
			case slack.LatencyReport:
				event := msg.Data.(slack.LatencyReport)
				fmt.Printf("Current latency: %v\n", event.Value)
			case *slack.SlackWSError:
				error := msg.Data.(*slack.SlackWSError)
				fmt.Printf("Error: %d - %s\n", error.Code, error.Msg)
			case *slack.StarAddedEvent:
				event := msg.Data.(*slack.StarAddedEvent)

				tsFloat, err := strconv.ParseFloat(event.Item.Message.Timestamp, 64)
				if err != nil {
					fmt.Printf("ERROR CONVERTING TIMESTAMP JIM, ERROR!: %s\n\n", err)
				}

				quote := new(storage.Quote)
				fmt.Println("\n")
				quote.Text = event.Item.Message.Text
				quote.Timestamp = int(tsFloat)
				quote.ChannelID = event.Item.Message.ChannelId
				quote.UserID = event.Item.Message.UserId

				channelInfo, err := a.Slack.GetChannelInfo(event.Item.Message.ChannelId)
				if err != nil {
					fmt.Printf("GetChannelInfo error: %s", err)
				}

				quote.ChannelName = channelInfo.Name

				userInfo, err := a.Slack.GetUserInfo(event.Item.Message.UserId)
				if err != nil {
					fmt.Printf("GetUserInfo error: %s", err)
				}

				if userInfo.RealName != "" {
					quote.UserName = userInfo.RealName
				} else {
					quote.UserName = userInfo.Name
				}

				searchQuote, err := a.Storage.SearchQuotes([]string{quote.Text})
				if err != nil {
					fmt.Printf("Error searching quote: %s \n", err)
				}
				if len(searchQuote) != 0 {
					fmt.Printf("Quote already exists, not saving.\n\n")
				} else {
					a.Storage.SaveQuote(quote)
				}

			default:
				fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}

func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)

	//Retrieve the part of text that matches <(.*?)>
	test := reg.FindAllString(text, -1)
	indexes := reg.FindAllStringIndex(text, -1)
	fmt.Println(test)
	laststart := 0
	result := make([]string, len(indexes)+1)
	fmt.Println(result)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

func ReplaceText(input string, search string, replacement string) string {

	reg := regexp.MustCompile(search)
	channelStrings := reg.FindAllString(input, -1)
	fmt.Println(channelStrings)

	for _, element := range channelStrings {
		fmt.Println(element)
		s := strings.Replace(input, element, replacement, -1)
		fmt.Println("result stringreplace: ", s)
		input = s
	}

	return input
}

func (a *AppContext) ReplaceTags(input string, search string) string {

	reg := regexp.MustCompile(search)
	channelStrings := reg.FindAllString(input, -1)
	fmt.Println(channelStrings)

	for _, element := range channelStrings {

		//If first three chars start with "<#C", then its a channel tag
		var newElement string
		switch element[1:3] {
		case "#C":
			{
				//Replace entire element with new channel name (fetched from api)
				slackChannel, err := a.Slack.GetChannelInfo(element[2:12])
				if err != nil {
					fmt.Printf("Error getting channel info: %s\n", err)
				}

				fmt.Printf("Got slack channel: %s\n", slackChannel)
				newElement = slackChannel.Name
			}
		case "@U":
			{
				//Replace element with user name (fetched from api)
				slackUser, err := a.Slack.GetUserInfo(element[2:12])
				if err != nil {
					fmt.Printf("Error getting channel info: %s\n", err)
				}

				fmt.Printf("Got slack user: %s\n", slackUser)
				if slackUser.RealName != "" {
					newElement = slackUser.RealName
				} else {
					newElement = slackUser.Name
				}
			}
		default:
			fmt.Println("Not a recognized tag")
		}

		fmt.Println(newElement)
		s := strings.Replace(input, element, newElement, -1)
		fmt.Println("result stringreplace: ", s)
		input = s
	}

	return input
}
