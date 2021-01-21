# GPU 使用率監視 Bot 

## 概要
Slack Bot を用いた GPU 使用率監視ツールです. 
確認したいサーバに対応した Bot にメンションで問い合わせることで, GPU の使用率の情報を受け取ることができます. 
サーバに予めコンテナを立てておくことで, Slack API からの要求に対して Golang の goroutine 機能を用いた非同期処理を実行します.
なお, コンテナ内部から GPU の情報を参照できるようにするために, サーバに以下の 3 つがインストールされていることを前提としています.

- Nvidia Driver 
- Docker 
- nvidia-docker2

![summary-img](./assets/summary.png)


## 設定

### 1. Bot の作成
各サーバを担当する Bot を作成します. 
Slack の App から Bots の追加を行い, API Token を取得してください.
また, Bot の名前は各サーバに対応したものにしてください. 
作成が完了した後は, Bot が活動するチャンネルを作成し「アプリを追加する」から Bot を招待してください. 

### 2. コンテナの起動
各サーバでコンテナを起動します.
起動の前に, src/bot.go 内の Slackparams インスタンスが持つ 3 つの変数を変更しておいてください.

```go
params := Slackparams{
    tokenID:   "aaaaa",     // Slack の API Token
    keyword:   "<@bbbbb>",  // Bot が起動するためのキーワード ( e.g. グループ名 )
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
