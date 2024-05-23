package settings

import (
	"log"

	"github.com/mymmrac/telego"
)

func SetBotName(bt *telego.Bot, botName string) {
	err := bt.SetMyName(&telego.SetMyNameParams{
		Name:         botName,
		LanguageCode: "en",
	})

	if err != nil {
		log.Fatal("failed set bot name: ", err)
	}
}

func SetMyDescription(bt *telego.Bot, description string) {
	// botga start bosyotganda ko'rinadigan qismi
	err := bt.SetMyDescription(&telego.SetMyDescriptionParams{
		Description:  description,
		LanguageCode: "en",
	})

	if err != nil {
		log.Fatal("error with set my full description: ", err)
	}
}

func SetMyAbout(bt *telego.Bot, aboutText string) {

	if len(aboutText) <= 120 {

		err := bt.SetMyShortDescription(&telego.SetMyShortDescriptionParams{
			ShortDescription: aboutText,
			LanguageCode:     "en",
		})

		if err != nil {
			log.Fatal("failed set about text: ", err)
		}
	} else {
		log.Fatal("aboutText length is higher than 120: ")
	}
}