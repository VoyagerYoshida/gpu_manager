FROM golang:1.15.7

RUN go get github.com/nlopes/slack

WORKDIR /go/src/
ADD ./src/ /go/src/bot

RUN go install bot

ENTRYPOINT /go/bin/bot
