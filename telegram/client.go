package telegram

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

// Telegram client that can execute our high-level commands
type TelegramClient struct {
	// Telegram client
	client *telegram.Client

	// Information about this user (when available)
	me *tg.User

	// Server ready?
	isServerReady chan struct{}
}

// Create a new Telegram client, manually
func NewTelegramClient(api_id int, api_hash string, session_file string) TelegramClient {
	return newTelegramClient(
		telegram.NewClient(api_id, api_hash, telegram.Options{
			SessionStorage: &session.FileStorage{Path: session_file},
			Device:         deviceConfig,
		}),
	)
}

// Create a new Telegram client from environment:
// APP_ID, APP_HASH, SESSION_DIR or SESSION_FILE
func NewTelegramClientFromEnvironment() (TelegramClient, error) {
	client, err := telegram.ClientFromEnvironment(telegram.Options{
		Device: deviceConfig,
	})
	if err != nil {
		return TelegramClient{}, errors.Wrapf(err, "failed to create Telegram Client")
	}
	return newTelegramClient(client), nil
}

// Authenticate, if necessary, and get the current user.
func (c TelegramClient) RunAuthenticationIfNecessary(ctx context.Context) (user *tg.User, err error) {
	err = c.client.Run(ctx, func(ctx context.Context) error {
		user, err = authenticate(ctx, c.client)
		return err
	})
	return
}

// Run server. Authenticate if necessary.
func (c TelegramClient) Run(ctx context.Context) error {
	return c.client.Run(ctx, func(ctx context.Context) error {
		// Authenticate
		if me, err := authenticate(ctx, c.client); err != nil {
			return err
		} else {
			c.me = me
		}

		// Server's ready: close the channel to signal it
		// All readers will be getting the channel's zero value
		close(c.isServerReady)

		// Wait for the context
		<-ctx.Done()
		return nil
	})
}

// Wait until the server is ready
func (c TelegramClient) WaitServerReady() {
	<-c.isServerReady
}

// ---

// Default device config for the Telegram client
var deviceConfig = telegram.DeviceConfig{
	DeviceModel: "Go StatusBot",
}

// Create a new Telegram client
func newTelegramClient(client *telegram.Client) TelegramClient {
	return TelegramClient{
		client:        client,
		me:            nil,
		isServerReady: make(chan struct{}),
	}
}

// Helper: authenticate, get information about self
func authenticate(ctx context.Context, client *telegram.Client) (*tg.User, error) {
	// Auth flow: use phone number, ask interactively
	flow := auth.NewFlow(
		interactiveTerminalAuth{},
		auth.SendCodeOptions{},
	)

	// Authenticate
	err := client.Auth().IfNecessary(ctx, flow)
	if err != nil {
		return nil, errors.Wrapf(err, "authentication failed")
	}

	// Get user
	if status, err := client.Auth().Status(ctx); err != nil {
		return nil, errors.Wrapf(err, "failed to load my user data")
	} else {
		return status.User, nil
	}
}
