package auth

import (
	"context"
	"telemonitor/config"

	"slices"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func AuthMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if slices.Contains(config.Cfg.AuthParams.Whitelist, update.Message.From.Username) && !update.Message.From.IsBot {
			next(ctx, b, update)
		}
	}
}
