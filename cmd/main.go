package main

import (
	"context"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"

	"durger/config"
	"durger/handler"
	"durger/service"

	"durger/pkg/logger"
	"durger/storage/postgres"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)

	store, err := postgres.New(context.Background(), cfg)

	if err != nil {
		log.Error("failed to connect postgres", logger.Error(err))
	}

	defer store.CloseDB()

	bot, err := telego.NewBot(cfg.BotToken, telego.WithDefaultDebugLogger())

	if err != nil {
		log.Error("failed to build the bot", logger.Error(err))
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	defer bot.Close()

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Error("failed to new handler", logger.Error(err))
	}

	defer bh.Stop()

	service := service.New(store, log)

	hand := handler.Handler{
		Bot:        bot,
		Config:     cfg,
		BotHandler: bh,
		Stora:      store,
		Service:    service,
		Log:        log,
	}
	hand.PrivateChat()
	bh.Start()
}