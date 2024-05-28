package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"

	"github.com/go-telegram/bot"

	"telemonitor/config"
	"telemonitor/internal/system_info"
	"telemonitor/internal/telegram/auth"
	"telemonitor/internal/telegram/commands"
)

// use relative path whether you starting instance
// for example ./ prefix if via "cmd/telemonitor/main"
const (
	configPath string = "./config.json"
)

func main() {
	fmt.Printf(
		"[info] Starting Telemonitor instance at host %x with PID: %d\n",
		system_info.Ipv4Address,
		system_info.Pid,
	)

	if !isRunningWithRoot() {
		fmt.Println("[Warning] Telemonitor must run in sudo mode to use all functions properly")
	}

	// initializing config via "configPath" directory and extracting token
	config.PathFromEntrypoint = configPath
	config.LoadConfig()

	token, isFinded := os.LookupEnv("TOKEN")
	if !isFinded {
		panic("[token] Token is not defined")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

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
	// binding help command handler to all possible start commands
	for _, helpCommand := range []string{"/help", "/start", "/bot", "/hello"} {
		b.RegisterHandler(bot.HandlerTypeMessageText, helpCommand, 0, commands.HelpHandler)
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/commands", 0, commands.CommandsMessageHandler)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/containers", 0, commands.ContainersHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/images", 0, commands.ImagesHandler)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/mem", 0, commands.MemoryLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/net", 0, commands.NetworkLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/cpu", 0, commands.CpuLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/disk", 0, commands.DiskLoadHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/rep", 0, commands.FullrepHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/sysinfo", 0, commands.SystemInfoHandler)
}

func isRunningWithRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}
	return currentUser.Username == "root"
}
