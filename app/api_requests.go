package app

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *apiRequest) Run(path string) {
	if path == "" {
		//call without session check
		r.HealthCheck()
	} else if path == "/password" {
		//call without session check
		r.Password()
	} else {
		if r.apiCheckSession() {
			switch strings.TrimPrefix(path, "/") {
			case "logout":
				r.Logout()

			case "say":
				r.Say()

			case "list_messages":
				r.ListMessages()

			default:
				message := "Unknown API request: " + path
				log.Println(message)
				r.setStatus("danger", message)
			}
		}
	}

	if r.getOutData("status") == "" {
		r.setStatus("ok", r.getOutData("message"))
	}
}

func (r *apiRequest) HealthCheck() {
	r.setStatus("ok", "API working")
	r.setOutData("auth", fmt.Sprintf("%t", r.session.Get("auth")))
}

func (r *apiRequest) Password() {
	if r.getInData("password") == Settings.GuiPassword {
		// Set user as authenticated
		r.session.Set("auth", true)

		r.session.Options(sessions.Options{
			MaxAge: 24 * 3600,
			Path:   "/",
		})

		r.session.Save()

		r.setStatus("ok", "You are authorized!")
	} else {
		r.setStatus("danger", "Wrong password")
	}
}

func (r *apiRequest) Logout() {
	r.session.Clear()

	r.session.Options(sessions.Options{
		MaxAge: -1, //remove
		Path:   "/",
	})

	r.session.Save()

	r.setStatus("info", "Good bye!")
}

func (r *apiRequest) Say() {
	text := r.getInData("message")
	text = PrepareTelegramHtml(text)
	msg := tgbotapi.NewMessage(Settings.BotChatID, text)
	msg.ParseMode = "HTML"

	if reply_to := r.getInDataInt("reply_to", 0); reply_to > 0 {
		msg.ReplyToMessageID = reply_to
	}

	if r.getInDataInt("silent", 0) != 0 {
		msg.DisableNotification = true
	}

	tgBot.Send(msg)

	//log.Println("Said:", text)
}

func (r *apiRequest) ListMessages() {
	updates_config := tgbotapi.NewUpdate(0)
	updates_config.Timeout = 1
	updates_config.Limit = 100
	updates_config.Offset = -100
	updates_config.AllowedUpdates = []string{"message"}

	updates_list, err := tgBot.GetUpdates(updates_config)
	if err != nil {
		r.setError(err.Error())
		return
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

	r.setOutData("list", list)
}
