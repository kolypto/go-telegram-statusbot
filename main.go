package main

import (
	"context"
	"os"

	"github.com/go-faster/errors"
	"github.com/kolypto/go-telegram-statusbot/httpapi"
	"github.com/kolypto/go-telegram-statusbot/telegram"
)

func main() {
	client, err := telegram.NewTelegramClientFromEnvironment()
	if err != nil {
		panic(err)
	}
	listenAddr, ok := os.LookupEnv("LISTEN")
	if !ok {
		panic(errors.Wrap(err, "SESSION_DIR not set or invalid"))
	}

	// Launch Telegram client
	ctx := context.Background()
	go func() {
		err := client.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()
	client.WaitServerReady()

	// Now start HTTP server
	httpapi.StartServer(ctx, listenAddr, client)
}
