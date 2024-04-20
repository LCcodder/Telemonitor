package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"

	"telemonitor/config"
	"telemonitor/internal/telegram/auth"
	"telemonitor/internal/telegram/commands"
)

const (
	configPath string = "../../config.json"
)

func main() {
	config.PathFromEntrypoint = configPath
	config.LoadConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	token, isFinded := os.LookupEnv("TOKEN")
	if !isFinded {
		panic("Token is not defined")
	}

	options := []bot.Option{
		bot.WithMiddlewares(auth.WhitelistAuthMiddleware),
	}

	bot, err := bot.New(token, options...)
	if err != nil {
		log.Fatal(err)
	}

	registerHandlers(bot)
	bot.Start(ctx)
}

func registerHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", 0, commands.HelpHandler)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/mem", 0, commands.MemoryLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/net", 0, commands.NetworkLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/cpu", 0, commands.CpuLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/disk", 0, commands.DiskLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/sysinfo", 0, commands.SystemInfoHandler)
}
