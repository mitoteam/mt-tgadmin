package app

type apiMessage struct {
	Message   string `json:"message"`
	MessageId int    `json:"message_id"`
	User      string `json:"user"`
	Date      string `json:"date"`
}
