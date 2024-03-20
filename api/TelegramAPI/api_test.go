package telegramAPI

import (
	"log"
	"testing"
)

func TestSendMessage(t *testing.T) {
	var bot TelegramBot = TelegramBot {
		TOKEN: "5960060789:AAEfJr9ssyG2VHh6b3yfaKox4i9Yi3siJN8",
	}
	bot.SendMessage(423215995, "hello", InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{},
	})
}

func TestConvertingID(t *testing.T) {
	log.Default().Println(int64(6930034973))
	// log.Default().Println(int32(6930034973))
}