package bot

import (
	tgbotapi "github.com/amsjavan/telegram-bot-api"
	"log"
	"taskulu/internal/postgres"
	"taskulu/pkg"
)

type TaskuluBot struct {
	log       *pkg.Logger
	db        *postgres.Database
	token     string
	updates   tgbotapi.UpdatesChannel
	botApi    *tgbotapi.BotAPI
	userState map[int]string
}

func New(log *pkg.Logger, db *postgres.Database, token string) *TaskuluBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Bot::", err)
	}
	bot.SetAPIEndpoint("https://tapi.bale.ai/bot%s/%s")
	bot.Debug = false // Has the library display every request and response.
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}
	return &TaskuluBot{
		log:       log,
		db:        db,
		token:     token,
		botApi:    bot,
		updates:   updates,
		userState: make(map[int]string),
	}
}

func (t *TaskuluBot) Run() {

	for update := range t.updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s, %v", update.Message.From.UserName, update.Message.Text, update.Message.From.ID)

		chatId := int(update.Message.Chat.ID)
		senderUserId := update.Message.From.ID

		if t.getState(chatId) == "PASSWORD" {
			pass := update.Message.Text
			err := t.db.UpsertTaskuluByPassword(chatId, pass)
			if err != nil {
				t.log.Error("Bot::", err)
				t.sendMessage(chatId, "خطا در ذخیره سازی در دیتابیس")
				t.entry(senderUserId, chatId)
				continue
			}
			t.setState(senderUserId, "START")
			t.sendMessage(chatId, "اطلاعات کاربری شما با موفقیت دریافت شد")
			continue
		}

		if t.getState(int(chatId)) == "USERNAME" {
			username := update.Message.Text
			err := t.db.UpsertTaskuluByUsername(int(update.Message.Chat.ID), username)
			if err != nil {
				t.log.Error("Bot::", err)
				t.sendMessage(chatId, "خطا در ذخیره سازی در دیتابیس")
				t.entry(senderUserId, chatId)
				continue
			}
			t.setState(update.Message.From.ID, "PASSWORD")
			t.sendMessage(chatId, "رمز عبور خود را وارد کنید")
			continue
		}

		if update.Message.Text == START {
			t.entry(senderUserId, chatId)
			continue
		}

		if update.Message.Text == LOGIN {
			t.setState(senderUserId, "USERNAME")
			t.sendMessage(chatId, "نام کاربری خود را وارد کنید")
			continue
		}

		if update.Message.Text == RETURN {
			t.entry(senderUserId, chatId)
			continue
		}
	}

}

const (
	START  = "/start"
	LOGIN  = "لاگین"
	HELP   = "راهنما"
	RETURN = "بازگشت"
)

func (t *TaskuluBot) sendMessage(id int, text string) {
	msg := tgbotapi.NewMessage(int64(id), text)
	msg.ReplyMarkup = t.returnMenu()
	_, err := t.botApi.Send(msg)
	if err != nil {
		t.log.Error("Bot::", err)
	}
}

func (t *TaskuluBot) setState(user int, state string) {
	t.userState[user] = state
}

func (t *TaskuluBot) getState(user int) string {
	return t.userState[user]
}

func (t *TaskuluBot) entry(senderUserId, chatId int) {
	t.setState(chatId, START)
	msg := tgbotapi.NewMessage(int64(chatId), "گزینه مورد نظر خود را انتخاب کنید‌")
	msg.ReplyMarkup = t.entryMenu()
	_, err := t.botApi.Send(msg)
	if err != nil {
		t.log.Error("Bot::", err)
	}
}

func (t *TaskuluBot) entryMenu() tgbotapi.ReplyKeyboardMarkup {
	loginItem := tgbotapi.KeyboardButton{Text: LOGIN}
	helpItem := tgbotapi.KeyboardButton{Text: HELP}
	menu := tgbotapi.NewKeyboardButtonRow(loginItem, helpItem)
	return tgbotapi.NewReplyKeyboard(menu)
}
func (t *TaskuluBot) returnMenu() tgbotapi.ReplyKeyboardMarkup {
	returnItem := tgbotapi.KeyboardButton{Text: RETURN}
	menu := tgbotapi.NewKeyboardButtonRow(returnItem)
	return tgbotapi.NewReplyKeyboard(menu)
}
