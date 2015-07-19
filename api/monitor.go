package api

import (
	"fmt"
	"strconv"
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
				fmt.Printf("\nMESSAGE IS STARRED!: %s\n\n", event.Item.Text)
				quote := new(storage.Quote)
				quote.Text = event.Item.Text

				tsFloat, err := strconv.ParseFloat(event.Item.Message.Timestamp, 64)
				if err != nil {
					fmt.Printf("ERROR JIM, ERROR!: %s\n\n", err)
				}

				quote.Timestamp = int(tsFloat)
				quote.ChannelID = event.Item.Message.ChannelId
				quote.UserName = event.Item.Message.Username
				quote.UserID = event.Item.Message.UserId

				fmt.Printf("message timestamp: %#v vs event timestamp: %#v\n", quote.Timestamp, event.EventTimestamp)
				a.Storage.SaveQuote(quote)

			default:
				fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
