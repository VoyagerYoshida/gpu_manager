package main

import (
    "strings"
    "strconv"
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
                    out, _ := exec.Command("nvidia-smi",
                                           "--query-gpu=index,name,memory.total,memory.used",
                                           "--format=csv,noheader,nounits").Output()
                    reader := csv.NewReader(strings.NewReader(string(out)))

                    for {
                        info, err := reader.Read()

                        if err != nil {
                            break
                        }

                        message := "[" + info[0] + "]" + info[1] + " :" + info[3] + "MiB /" + info[2] + "MiB"
                        numer, _ := strconv.ParseFloat(strings.TrimSpace(info[3]), 64)
                        denom, _ := strconv.ParseFloat(strings.TrimSpace(info[2]), 64)
                        message += " -> " + strconv.FormatFloat(numer * 100 / denom, 'f', 2, 64) + "%"
                        params.rtm.SendMessage(params.rtm.NewOutgoingMessage(message, ev.Channel))
                    }
                }
        }
    }
}
