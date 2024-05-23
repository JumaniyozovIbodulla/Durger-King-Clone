package keyboards

import "github.com/mymmrac/telego"

func MainInlineButton() [][]telego.InlineKeyboardButton {
	keyboard := [][]telego.InlineKeyboardButton{
		[]telego.InlineKeyboardButton{
			telego.InlineKeyboardButton{
				Text: "Ovqat buyurtma qilish",
				WebApp: &telego.WebAppInfo{
					URL: "https://durgerking-bot-git-main-mominovdev.vercel.app/",
				},
			},
		},
	}
	return keyboard
}