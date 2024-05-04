package commands

import (
	"context"
	"log"

	"telemonitor/config"
	"telemonitor/internal/monitoring"
	"telemonitor/internal/system_info"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	welcomeMessage *string  = &config.Cfg.Messages.Welcome
	commandsList   []string = []string{"/mem", "/help", "/info", "/disk", "/net"}
)

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      *welcomeMessage,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func MemoryLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetMemoryLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NetworkLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetNetworkLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func DiskLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetDiskLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func CpuLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetCpuLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func SystemInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      system_info.GetSystemInfo(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}
