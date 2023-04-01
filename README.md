Telegram status setter
======================

This is a Telegram client (userbot) that can do things for you.
It has an HTTP interface so that you can command it remotely.

Features:

* Set Emoji status

## Build

```console
$ go build
```

## Run

Start server:

```console
$ APP_ID= APP_HASH= SESSION_FILE=./session.json LISTEN=:8080 ./go-telegram-statusbot
```

Set emoji status:

```console
$ curl http://localhost:8080/set-status/icq-online
$ curl http://localhost:8080/set-status/icq-online
```
