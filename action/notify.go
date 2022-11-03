package action

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// const (
//
//	ChatID  = -1001358735880
//	TgToken = "781105443:AAHVIJ0e3YLip2pnfqpbkXNc-ibTQnvhAFU"
//
// )
func getBotEnv() (chatid int64, tgtoken string) {
	chatid_s, exist := os.LookupEnv("TG_CHAT_ID")
	if !exist {
		log.Panicln("Can not get ENV TG_CHAT_ID")
	}
	chatid, _ = strconv.ParseInt(chatid_s, 10, 64)
	tgtoken, exist = os.LookupEnv("TG_TOKEN")
	if !exist {
		log.Panicln("Can not get ENV TG_TOKEN")
	}
	return
}

func PushTgNotification(text string) (err error) {
	ChatID, TgToken := getBotEnv()
	bot, err := tgbotapi.NewBotAPI(TgToken)
	if err != nil {
		log.Panic(err)
	}
	// bot.Debug = true

	//log.Printf("Telegram Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(ChatID, text)
	_, err = bot.Send(msg)

	return
	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message != nil { // If we got a message
	// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 		msg.ReplyToMessageID = update.Message.MessageID

	// 		bot.Send(msg)
	// 	}
	// }
}
