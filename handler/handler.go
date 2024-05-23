package handler

import (
	"durger/config"
	"durger/pkg/logger"
	"durger/service"
	"durger/storage"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Handler struct {
	Bot        *telego.Bot
	Config     config.Config
	BotHandler *th.BotHandler
	Stora      storage.IStorage
	Service    service.IServiceManager
	Log        logger.ILogger
}

func NewBot(bot *telego.Bot, config config.Config, botHandler *th.BotHandler, store storage.IStorage, service service.IServiceManager, log logger.ILogger) Handler {
	return Handler{
		Bot:        bot,
		Config:     config,
		BotHandler: botHandler,
		Stora:      store,
		Service:    service,
		Log:        log,
	}
}