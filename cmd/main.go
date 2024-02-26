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
	configPath string = "../config.json"
)

func main() {
	config.LoadConfig(configPath)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	token, isFinded := os.LookupEnv("TOKEN")
	if !isFinded {
		panic("Token is not defined")
	}

	options := []bot.Option{
		bot.WithMiddlewares(auth.AuthMiddleware),
	}

	b, err := bot.New(token, options...)
	if err != nil {
		log.Fatal(err)
	}

	registerHandlers(b)

	b.Start(ctx)
}

func registerHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", 0, commands.HelpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/info", 0, commands.InfoHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/mem", 0, commands.MemoryHandler)
}
