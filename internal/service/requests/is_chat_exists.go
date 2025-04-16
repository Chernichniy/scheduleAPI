package requests

import (
	"context"

	"github.com/go-telegram/bot"
)

func getChatID(username string, ctx context.Context, b *bot.Bot) (int64, error) {

	chatIDRaw, err := b.GetChat(ctx, &bot.GetChatParams{
		ChatID: username,
	})

	if err != nil {
		//Log(r).WithError(err).Error("failed to find chat with needed user")
		return 0, err
	}

	return chatIDRaw.ID, err
}
