package commands

import (
	"log"

	"github.com/mymmrac/telego"
)

func SetAllPrivateCommands(bt *telego.Bot) {

	commands := []telego.BotCommand{}

	startCommand := telego.BotCommand{
		Command:     "start",
		Description: "Botni ishga tushirish",
	}

	helpCommand := telego.BotCommand{
		Command:     "help",
		Description: "Yo'riqnoma bilan tanishish",
	}

	commands = append(commands, startCommand, helpCommand)

	params := telego.SetMyCommandsParams{
		Commands: commands,
		Scope: &telego.BotCommandScopeAllPrivateChats{
			Type: "all_private_chats",
		},
	}

	if err := bt.SetMyCommands(&params); err != nil {
		log.Fatal("failed to set all private commands")
	}
}
