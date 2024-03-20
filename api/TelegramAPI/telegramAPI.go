package telegramAPI

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (bot *TelegramBot) getMethodURL(methodName string) string {
	var base string = "https://api.telegram.org/bot" + bot.TOKEN + "/"
	switch methodName {
	case "sendMessage":
		fallthrough
	case "answerCallbackQuery":
		fallthrough
	case "answerInlineQuery":
		fallthrough
	case "editMessageText":
		return base + methodName
	}

	return ""
}

type TelegramBot struct {
	TOKEN string
}

type MessageToSend struct {
	ChatID      int64                `json:"chat_id"`
	Text        string               `json:"text"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
	ParseMode   string               `json:"parse_mode"`
}

type MessageToEdit struct {
	// ChatID          int64                `json:"chat_id"`
	InlineMessageID string               `json:"inline_message_id"`
	Text            string               `json:"text"`
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup"`
	ParseMode       string               `json:"parse_mode"`
}

type InlineQueryToAnswer struct {
	InlineQueryID string                     `json:"inline_query_id"`
	Results       []InlineQueryResultArticle `json:"results"`
	CacheTime     int                        `json:"cache_time"`
	IsPersonal    bool                       `json:"is_personal"`
}

type CallbackQueryToAnswer struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InputTextMessageContent struct {
	MessageText      string `json:"message_text"`
	ParseMode string `json:"parse_mode"`
}

type InlineQueryResultArticle struct {
	Type                string                  `json:"type"`
	ID                  string                  `json:"id"`
	Title               string                  `json:"title"`
	Description         string                  `json:"description"`
	ThumbnailURL        string                  `json:"thumbnail_url"`
	InputMessageContent InputTextMessageContent `json:"input_message_content"`
	ReplyMarkup         InlineKeyboardMarkup    `json:"reply_markup"`
}

func send(b []byte, url string) {
	log.Default().Println(string(b))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    bodyString := string(bodyBytes)
    log.Default().Println(bodyString)
}

func (bot *TelegramBot) SendMessage(chatID int64, messageText string, replyMarkup InlineKeyboardMarkup) {
	var url string = bot.getMethodURL("sendMessage")
	log.Default().Println(url)
	var message MessageToSend = MessageToSend{
		ChatID:      chatID,
		Text:        messageText,
		ReplyMarkup: replyMarkup,
		ParseMode:   "HTML",
	}

	var jsonBytes, _ = json.Marshal(message)
	send(jsonBytes, url)
}

func (bot *TelegramBot) AnswerInlineQuery(queryID string, results []InlineQueryResultArticle, cacheTime int, isPersonal bool) {
	var url string = bot.getMethodURL("answerInlineQuery")

	message := InlineQueryToAnswer{
		InlineQueryID: queryID,
		Results:       results,
		CacheTime:     cacheTime,
		IsPersonal:    isPersonal,
	}

	var jsonBytes, _ = json.Marshal(message)
	send(jsonBytes, url)
}

func (bot *TelegramBot) AnswerCallbackQuery(callbackQueryID, text string, showAlert bool) {
	var url string = bot.getMethodURL("answerCallbackQuery")

	message := CallbackQueryToAnswer{
		CallbackQueryID: callbackQueryID,
		Text:            text,
		ShowAlert:       showAlert,
	}

	var jsonBytes, _ = json.Marshal(message)
	send(jsonBytes, url)
}

func (bot *TelegramBot) EditMessageText(chatID int64, inlineMessageID string, textNew string, replyMarkup InlineKeyboardMarkup) {
	var url string = bot.getMethodURL("editMessageText")

	message := MessageToEdit{
		//ChatID:          int64(chatID),
		InlineMessageID: inlineMessageID,
		Text:            textNew,
		ReplyMarkup:     replyMarkup,
		ParseMode:       "HTML",
	}

	var jsonBytes, _ = json.Marshal(message)
	send(jsonBytes, url)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	return
}
