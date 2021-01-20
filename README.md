# GPU 使用率管理 Bot 

## 概要

![summary-img](./assets/summary.png)

## 設定
src/bot.go 内の Slackparams インスタンスが持つ 3 つの変数を変更する.

```go
params := Slackparams{
    tokenID:   "aaaaa",     // Slack の API Token
    botID:     "<@bbbbb>",  // bot の名前
    channelID: "ccccc",     // 使用する channel の名前
}
```


## コマンド

### Docker image の作成
```
$ make build
```

### Docker container の起動
```
$ make run
```

### Docker container の停止
```
$ make stop
```
