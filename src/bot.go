package main

import (
    "strings"
    "github.com/nlopes/slack"
)

type Slackparams struct {
	tokenID   string
	botID     string
	channelID string
	rtm       *slack.RTM
}


func main() {
    params := Slackparams{
        tokenID:   "aaaaa",
        botID:     "<@bbbbb>",
        channelID: "ccccc",
    }

    api := slack.New(params.tokenID)
    params.rtm = api.NewRTM()
    go rtm.ManageConnection()

    for msg := range rtm.IncomingEvents {
        switch ev := msg.Data.(type) {
            case *slack.MessageEvent:
                if srtings(ev.Msg.Text, params.botID) && ev.Channel == params.channelID {
                    params.rtm.SendMessage(rtm.NewOutgoingMessage("hoge", ev.Channel))
                }
        }
    }
}
