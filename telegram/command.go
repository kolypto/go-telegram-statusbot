package telegram

import (
	"context"
	"time"

	"github.com/gotd/td/tg"
)

// Command: Set Emoji status
func (c TelegramClient) SetEmojiStatus(ctx context.Context, documentId int64) error {
	return setEmojiStatus(ctx, c, &tg.EmojiStatus{DocumentID: documentId})
}

// Command: Set Emoji status, with a timeout
func (c TelegramClient) SetEmojiStatusFor(ctx context.Context, documentId int64, duration time.Duration) error {
	until := time.Now().Add(duration).Unix()
	return setEmojiStatus(ctx, c, &tg.EmojiStatusUntil{DocumentID: documentId, Until: int(until)})
}

//---

// Helper: set emoji status (one-liner)
func setEmojiStatus(ctx context.Context, client TelegramClient, status tg.EmojiStatusClass) error {
	<-client.isServerReady
	_, err := client.client.API().AccountUpdateEmojiStatus(ctx, status)
	return err
}
