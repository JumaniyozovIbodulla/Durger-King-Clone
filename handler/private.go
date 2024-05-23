package handler

import (
	"context"
	"durger/keyboards"
	"durger/models"
	"durger/pkg/logger"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	// "github.com/spf13/cast"
)

func (h Handler) PrivateChat() {

	// h.BotHandler.HandleMessage(func(bot *telego.Bot, message telego.Message) {
	// 	fmt.Println(message.WebAppData.Data)
	// }, func(update telego.Update) bool {
	// 	if update.Message.Chat.Type == "private" {
	// 		return true
	// 	}
	// 	return false
	// })

	keyboard := keyboards.MainInlineButton()

	// Command start
	h.BotHandler.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		if message.Chat.Type == "private" {
			_, err := bot.SendMessage(&telego.SendMessageParams{
				ChatID: telego.ChatID{
					ID: message.From.ID,
				},
				Text: "O'zingiz uchun eng zo'r tushlikni tanlang 😊",
				ReplyMarkup: &telego.InlineKeyboardMarkup{
					InlineKeyboard: keyboard,
				},
			})

			if err != nil {
				h.Log.Error("failed start command", logger.Error(err))
			}
			h.Stora.UserStorage().Create(context.Background(), models.AddUser{
				Id:           message.From.ID,
				FirstName:    message.From.FirstName,
				LastName:     message.From.LastName,
				Username:     message.From.Username,
				IsPremium:    message.From.IsPremium,
				LanguageCode: message.From.LanguageCode,
			})
		} else {
			_, err := bot.SendMessage(&telego.SendMessageParams{
				ChatID: telego.ChatID{
					ID: message.From.ID,
				},
				Text: "Botdan faqat o'zining chat qismi orqali foydalanishingiz mumkin.",
			})

			if err != nil {
				h.Log.Error("failed start command else", logger.Error(err))
			}
		}

	}, th.CommandEqual("start"))

	// Command start

	h.BotHandler.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		if message.Chat.Type == "private" {
			_, err := bot.SendMessage(&telego.SendMessageParams{
				ChatID: telego.ChatID{
					ID: message.From.ID,
				},
				Text: "Bu qism hali tayyor emas 😞",
			})

			if err != nil {
				h.Log.Error("failed to help command", logger.Error(err))
			}

		} else {
			_, err := bot.SendMessage(&telego.SendMessageParams{
				ChatID: telego.ChatID{
					ID: message.From.ID,
				},
				Text: "Botdan faqat o'zining chat qismi orqali foydalanishingiz mumkin.",
			})

			if err != nil {
				h.Log.Error("failed to help command else", logger.Error(err))
			}
		}

	}, th.CommandEqual("help"))

}
