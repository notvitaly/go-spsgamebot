package Updates

import "net/http"

type Update struct {
	UpdateID      int           `json:"update_id"`
	Message       Message       `json:"message"`
	InlineQuery   InlineQuery   `json:"inline_query"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

type Message struct {
	MessageID int32  `json:"message_id"`
	From      User   `json:"from"`
	Text      string `json:"text"`
}
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type InlineQuery struct {
	ID       string `json:"id"`
	From     User   `json:"from"`
	Query    string `json:"query"`
	ChatType string `json:"chat_type"`
}

type CallbackQuery struct {
	ID              string `json:"id"`
	From            User   `json:"from"`
	InlineMessageID string `json:"inline_message_id"`
	Data            string `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	return
}
