package main

import (
	"github.com/amsjavan/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("512522922:3a615125d2d088f1768c9fb29acb9d7ff5127ad8")
	bot.SetAPIEndpoint("https://tapi.bale.ai/bot%s/%s")
	if err != nil {
		panic(err) // You should add better error handling than this!
	}

	bot.Debug = false // Has the library display every request and response.
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s, %v", update.Message.From.UserName, update.Message.Text, update.Message.From.ID)

		if update.Message.Text == START {
			setState(update.Message.From.ID,START)
			entry(bot, update)
		}

		if update.Message.Text == LOGIN {
			setState(update.Message.From.ID,LOGIN)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "نام کاربری خود را وارد کنید")
			msg.ReplyMarkup = returnMenu()

			bot.Send(msg)
		}

		if update.Message.Text == RETURN {
			entry(bot,update)
		}



		}


}

var userState = make(map[int]string)

const (
	START = "/start"
	LOGIN = "لاگین"
	HELP = "راهنما"
	RETURN = "بازگشت"
)

func setState(user int, state string){
	userState[user] = state
}

func entry(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "گزینه مورد نظر خود را انتخاب کنید‌")
	msg.ReplyMarkup = entryMenu()
	bot.Send(msg)
}

func entryMenu() tgbotapi.ReplyKeyboardMarkup {
	loginItem := tgbotapi.KeyboardButton{ Text: LOGIN}
	helpItem := tgbotapi.KeyboardButton{ Text: HELP}
	menu := tgbotapi.NewKeyboardButtonRow(loginItem, helpItem)
	return tgbotapi.NewReplyKeyboard(menu)
}
func returnMenu() tgbotapi.ReplyKeyboardMarkup {
	returnItem := tgbotapi.KeyboardButton{ Text: RETURN}
	menu := tgbotapi.NewKeyboardButtonRow(returnItem)
	return tgbotapi.NewReplyKeyboard(menu)
}