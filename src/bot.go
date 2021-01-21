package main

import (
    "fmt"
    "strings"
    "os/exec"
    "encoding/csv"
    "github.com/nlopes/slack"
)

type Slackparams struct {
	tokenID   string
	keyword   string
	channelID string
	rtm       *slack.RTM
}


func gpuInfo() []string {
    out, _ := exec.Command(
        "nvidia-smi",
        "--query-gpu=index,name,memory.total,memory.used",
        "--format=csv,noheader,nounits").Output()
    reader := csv.NewReader(strings.NewReader(string(out)))
    record, _ := reader.Read()
    return record
}


func main() {
    params := Slackparams{
        tokenID:   "aaaaa",
        keyword:   "bbbbb",
        channelID: "ccccc",
    }

    api := slack.New(params.tokenID)
    params.rtm = api.NewRTM()
    go params.rtm.ManageConnection()

    for msg := range params.rtm.IncomingEvents {
        switch ev := msg.Data.(type) {
            case *slack.MessageEvent:
                if strings.EqualFold(ev.Msg.Text, params.keyword) && ev.Channel == params.channelID {
                    info := gpuInfo()
                    fmt.Print(ev.Msg.Text)
                    params.rtm.SendMessage(params.rtm.NewOutgoingMessage(info[1], ev.Channel))
                }
        }
    }
}
