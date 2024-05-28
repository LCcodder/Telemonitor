package commands

import (
	"context"
	"fmt"
	"log"

	"telemonitor/config"
	"telemonitor/internal/monitoring"
	"telemonitor/internal/system_info"

	"github.com/docker/docker/client"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	welcomeMessage *string = &config.Cfg.Messages.Welcome
	// update this message as long as you adding new commands
	commandsMessage string                   = "<b>Commands:</b>\n/sysinfo - shows your Telemonitor instance's hardware and firmware info\n/cpu - shows cpu loading\n/mem - shows RAM and SWAP usage\n/disk - shows your main directory capacity\n/net - shows network interfaces info\n/rep - shows full system metrics report\n/containers - shows all hosted docker containers\n/images - shows downloaded and created images"
	c, _                                     = client.NewClientWithOpts(client.FromEnv)
	d               monitoring.DockerMetrics = monitoring.DockerMetrics{Client: *c}
)

func CommandsMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      commandsMessage,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      *welcomeMessage,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func ContainersHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      d.GetAllContainers(ctx, 10),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func ImagesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println()
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      d.GetImages(ctx),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func MemoryLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetMemoryLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func NetworkLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetNetworkLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func DiskLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetDiskLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func CpuLoadHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetCpuLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func SystemInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      system_info.GetSystemInfo(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}

func FullrepHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetCpuLoad() + "\n\n" + monitoring.GetMemoryLoad() + "\n\n" + monitoring.GetDiskLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Print(err)
	}
}
