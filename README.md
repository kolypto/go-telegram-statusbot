Telegram status setter
======================

## Build

```console
$ go build
```

## Run

Start server:

```console
$ APP_ID= APP_HASH= SESSION_FILE=session.json LISTEN=:8080 ./go-telegram-statusbot
```

Now make HTTP requests:

```console
$ curl http://localhost:8080/set-status/icq-online
```