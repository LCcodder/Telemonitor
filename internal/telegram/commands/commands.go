package commands

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"telemonitor/internal/monitoring"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	welcomeMessage string = "<b>All Commands:</b>\n /all_metrics - will show all of the available system metrics\n /system_info - shows basic machine hardware data\n /mem_load - shows RAM and swap memory load\n /disk_load - shows '/' disk partition load and write/read data\n /cpu_load - shows cpu load stats\n /netstat - shows basic network statistics (only for the 1st interface)\n /connections - shows all tcp and upd connections"
)

var (
	pid         int    = os.Getpid()
	ipv4Address []byte = getOutboundIP()
	infoMessage string = fmt.Sprintf("Your <i>Telebot</i> instance is running on host: <b>%x</b> with PID: <b>%d</b>", ipv4Address, pid)
)

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      welcomeMessage,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func InfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      infoMessage,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func MemoryHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      monitoring.GetMemoryLoad(),
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Fatal(err)
	}
}
