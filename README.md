Telegram status setter
======================

This is a Telegram client (userbot) that can do things for you.
It has an HTTP interface so that you can command it remotely.

Features:

* Set Emoji status

## Easy Start (with Docker)

Create the container:

* `./statusbot-telegram-data` is where the bot will keep its `session.json` with your credentials
* `-p` is the port the container will be listening on: `7000`
* `APP_ID` and `APP_HASH` are your Telegram API credentials: <https://my.telegram.org/apps>
* `--name` is the name for your container

When the container starts, you'll be asked to sign in interactively.
Prepare your phone number, your 2FA code, and your Telegram Cloud password.
When done, press `Ctrl-C`.

```console
$ mkdir statusbot-telegram-data
$ docker run -it \
    -v $(pwd)/statusbot-telegram-data:/app/data -p7000:8080 \
    -e APP_ID=... -e APP_HASH=... \
    --name "telegram-statusbot" kolypto/go-telegram-statusbot
Phone: +7901...
2FA code: 12345
Telegram cloud password: Password

2023/04/01 15:28:09 Server listening on: :8080
Ctrl-C
```

Now run it for real:

```console
$ docker start telegram-statusbot
```

Anc check that it's running alright:

```console
$ docker ps -a
CONTAINER ID   IMAGE                           STATUS         PORTS                   NAMES
fa372225ed35   kolypto/go-telegram-statusbot   Up 3 minutes   0.0.0.0:7000->8080/tcp, kolypto-telegram-statusbot
$ docker logs kolypto-telegram-statusbot
2023/04/01 15:28:09 Server listening on: :8080
```

## Or build it yourself! (For Go folks)

### Option 1. Locally

```console
$ go build
```

Start server:

```console
$ APP_ID= APP_HASH= SESSION_FILE=./session.json LISTEN=:8080 ./go-telegram-statusbot
```

### Option 2. With Docker

Build and first time run:

```console
$ docker-compose build bot
$ APP_ID=... APP_HASH=... docker-compose run --rm bot
```

Then run:

```console
$ docker-compose up -d bot
```


## Use

To set an emoji status on your account, use HTTP API:

```console
$ curl -X POST http://localhost:8080/set-status/icq-online
$ curl -X POST http://localhost:8080/set-status/5276100343673924208
```
