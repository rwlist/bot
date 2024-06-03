package main

import (
	"log/slog"
	"net/http"

	"github.com/rwlist/bot/internal/conf"
	"github.com/rwlist/bot/internal/mlog"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := conf.ParseEnv()
	if err != nil {
		mlog.Fatal("failed to parse config from env", mlog.Err(err))
	}

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(cfg.PrometheusBind, mux)
		if err != nil && err != http.ErrServerClosed {
			mlog.Fatal("prometheus server error", mlog.Err(err))
		}
	}()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		mlog.Fatal("failed to create bot", mlog.Err(err))
	}
	bot.Debug = true
	slog.Info("Authorized on account", slog.String("username", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			slog.Info("message from user", slog.String("username", update.Message.From.UserName), slog.String("message", update.Message.Text))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
