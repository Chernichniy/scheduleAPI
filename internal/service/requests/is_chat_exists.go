package requests

import (
	"context"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func getChatID(username string, ctx context.Context, b *bot.Bot, update *models.Update, r *http.Request) (int64, error) {

	chatIDRaw, err := b.GetChat(ctx, &bot.GetChatParams{
		ChatID: username,
	})

	if err != nil {
		//Log(r).WithError(err).Error("failed to find chat with needed user")
		return 0, err
	}

	return chatIDRaw.ID, err
}
