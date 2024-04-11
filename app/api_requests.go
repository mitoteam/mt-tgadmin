package app

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mitoteam/goappbase"
)

func apiCheckAuth(r *goappbase.ApiRequest) bool {
	if r.SessionGet("auth") == true {
		return true
	} else {
		r.SetErrorStatus("Auth Required")
		return false
	}
}

func Api_HealthCheck(r *goappbase.ApiRequest) error {
	r.SetOutData("auth", apiCheckAuth(r))
	r.SetOkStatus("API works: " + App.AppName)

	return nil
}

func Api_Password(r *goappbase.ApiRequest) error {
	if r.GetInData("password") == Settings.GuiPassword {
		// Set user as authenticated
		r.SessionSet("auth", true)

		r.Session().Save()
		r.SetOkStatus("You are authorized")
	} else {
		r.SetErrorStatus("Wrong password!")
	}
	return nil
}

func Api_Logout(r *goappbase.ApiRequest) error {
	if !apiCheckAuth(r) {
		return nil
	}

	r.SessionClear()
	r.SetOkStatus("Good bye!")

	return nil
}

func Api_Say(r *goappbase.ApiRequest) error {
	if !apiCheckAuth(r) {
		return nil
	}

	text := r.GetInData("message")
	text = PrepareTelegramHtml(text)
	msg := tgbotapi.NewMessage(Settings.BotChatID, text)
	msg.ParseMode = "HTML"

	if reply_to := r.GetInDataInt("reply_to", 0); reply_to > 0 {
		msg.ReplyToMessageID = reply_to
	}

	if r.GetInDataInt("silent", 0) != 0 {
		msg.DisableNotification = true
	}

	if _, err := tgBot.Send(msg); err != nil {
		r.SetErrorStatus("Error sending message: " + err.Error())
		return nil
	}

	//log.Println("Said:", text)
	return nil
}

func Api_ListMessages(r *goappbase.ApiRequest) error {
	if !apiCheckAuth(r) {
		return nil
	}

	updates_config := tgbotapi.NewUpdate(0)
	updates_config.Timeout = 1
	updates_config.Limit = 100
	updates_config.Offset = -100
	updates_config.AllowedUpdates = []string{"message"}

	updates_list, err := tgBot.GetUpdates(updates_config)
	if err != nil {
		r.SetErrorStatus(err.Error())
		return nil
	}

	list := make([]*apiMessage, 0, len(updates_list))

	//from end to beginning
	for i := len(updates_list) - 1; i >= 0; i-- {
		update := updates_list[i]

		if update.Message.Chat.ID != Settings.BotChatID {
			//skip messages from other channels or chats
			continue
		}

		m := &apiMessage{}
		m.Message = update.Message.Text
		m.MessageId = update.Message.MessageID
		m.User = update.Message.From.FirstName + " " + update.Message.From.LastName + " = @" + update.Message.From.UserName
		m.Date = time.Unix(int64(update.Message.Date), 0).Format("2006-01-02 15:04:05")

		list = append(list, m)
	}

	//log.Println("Updates count: ", len(list))

	r.SetOutData("list", list)
	return nil
}
