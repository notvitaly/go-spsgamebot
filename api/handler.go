package handler

import (
	"encoding/json"
	"fmt"
	telegramAPI "gospsgamemod/TelegramAPI"
	"gospsgamemod/Updates"
	"gospsgamemod/game"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var bot telegramAPI.TelegramBot = telegramAPI.TelegramBot{
	TOKEN: os.Getenv("TELEGRAM_BOT_TOKEN"),
}

func Handler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Default().Println(string(bodyBytes))
	var update Updates.Update
	json.Unmarshal(bodyBytes, &update)

	if (update.Message != Updates.Message{}) {
		handleMessage(update.Message)
	} else if (update.CallbackQuery != Updates.CallbackQuery{}) {
		handleCallbackQuery(update.CallbackQuery)
	} else if (update.InlineQuery != Updates.InlineQuery{}) {
		handleInlineQuery(update.InlineQuery)
	}

	w.WriteHeader(http.StatusOK)
}

func handleMessage(message Updates.Message) {
	if message.Text == "/start" {
		bot.SendMessage(int64(message.From.ID), "Hello! To play Scissor-Paper-Stones, call me inline in a chat with your friend", telegramAPI.InlineKeyboardMarkup{InlineKeyboard: [][]telegramAPI.InlineKeyboardButton{}})
	}
}
func handleInlineQuery(inlineQuery Updates.InlineQuery) {
	imagesURLs := []string{
		"https://em-content.zobj.net/source/microsoft/379/keycap-digit-one_31-fe0f-20e3.png",
		"https://em-content.zobj.net/source/microsoft/379/keycap-digit-two_32-fe0f-20e3.png",
		"https://em-content.zobj.net/source/microsoft/379/keycap-digit-three_33-fe0f-20e3.png",
	}

	imagesURLsForPreparedTurns := []string{
		"https://em-content.zobj.net/source/microsoft/379/moai_1f5ff.png",
		"https://em-content.zobj.net/source/microsoft/379/scissors_2702-fe0f.png",
		"https://em-content.zobj.net/source/microsoft/379/roll-of-paper_1f9fb.png",
	}

	numbersStrings := []string{
		"One Round", "Two Rounds", "Three Rounds",
	}

	turns := []string{
		"🗿", "✂️", "🧻",
	}

	var results []telegramAPI.InlineQueryResultArticle

	for i := range numbersStrings {
		g := game.Game{
			InitiatorID: -1,
			RoundsTotal: i + 1,
		}
		var buttons [][]telegramAPI.InlineKeyboardButton = [][]telegramAPI.InlineKeyboardButton{}
		var row []telegramAPI.InlineKeyboardButton

		for j, turnEmoji := range turns {
			g.Turn = j
			row = append(row, telegramAPI.InlineKeyboardButton{
				Text:         turnEmoji,
				CallbackData: g.EncodeGame(),
			})
		}

		buttons = append(buttons, row)

		inlineQueryResult := telegramAPI.InlineQueryResultArticle{
			Type:         "article",
			ID:           strconv.Itoa(i),
			Title:        numbersStrings[i],
			ThumbnailURL: imagesURLs[i],
			InputMessageContent: telegramAPI.InputTextMessageContent{
				MessageText: fmt.Sprintf("%s wants to play Scissor-Paper-Stones!", inlineQuery.From.FirstName),
			},
			ReplyMarkup: telegramAPI.InlineKeyboardMarkup{
				InlineKeyboard: buttons,
			},
		}
		results = append(results, inlineQueryResult)
	}

	for i, emojiTurn := range turns {
		g := game.Game{
			InitiatorID: -1,
			RoundsTotal: 1,
		}
		g.MakeNewTurn(inlineQuery.From.ID, inlineQuery.From.FirstName, i)

		var buttons [][]telegramAPI.InlineKeyboardButton = [][]telegramAPI.InlineKeyboardButton{}
		var row []telegramAPI.InlineKeyboardButton

		for j, turnEmoji := range turns {
			g.Turn = j
			row = append(row, telegramAPI.InlineKeyboardButton{
				Text:         turnEmoji,
				CallbackData: g.EncodeGame(),
			})
		}

		buttons = append(buttons, row)

		inlineQueryResult := telegramAPI.InlineQueryResultArticle{
			Type:         "article",
			ID:           strconv.Itoa(i + 3),
			Title:        fmt.Sprintf("Make a new turn with %s", emojiTurn),
			ThumbnailURL: imagesURLsForPreparedTurns[i],
			InputMessageContent: telegramAPI.InputTextMessageContent{
				MessageText: g.ToString(),
				ParseMode:   "HTML",
			},
			ReplyMarkup: telegramAPI.InlineKeyboardMarkup{
				InlineKeyboard: buttons,
			},
		}
		results = append(results, inlineQueryResult)
	}

	bot.AnswerInlineQuery(inlineQuery.ID, results, 0, true)
}
func handleCallbackQuery(callback Updates.CallbackQuery) {
	var game game.Game = game.DecodeString(callback.Data)
	makingTurnResult := game.MakeNewTurn(int64(callback.From.ID), callback.From.FirstName, game.Turn)
	if !makingTurnResult {
		bot.AnswerCallbackQuery(callback.ID, "Please wait for your turn", true)
	} else {
		emojis := []string{
			"🗿", "✂️", "🧻",
		}
		var InlineKeyboard telegramAPI.InlineKeyboardMarkup
		var buttons [][]telegramAPI.InlineKeyboardButton = [][]telegramAPI.InlineKeyboardButton{}
		var row []telegramAPI.InlineKeyboardButton

		if !game.Finished {
			for i, emoji := range emojis {
				game.Turn = i
				row = append(row, telegramAPI.InlineKeyboardButton{
					Text:         emoji,
					CallbackData: game.EncodeGame(),
				})
			}

			buttons = append(buttons, row)
		}

		InlineKeyboard.InlineKeyboard = buttons
		bot.EditMessageText(callback.From.ID, callback.InlineMessageID, game.ToString(), InlineKeyboard)
	}

}
