package app

import (
	"fmt"
	"log"
	"strings"

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
	r.setOutData("auth", fmt.Sprintf("%t", r.session.Values["auth"]))
}

func (r *apiRequest) Password() {
	if r.getInData("password") == Global.Settings.GuiPassword {
		r.setStatus("ok", "You are authorized!")

		// Set user as authenticated
		r.session.Values["auth"] = true
		r.session.Options.MaxAge = 86400
		r.session.Save(r.request, *r.responseWriter)
	} else {
		r.setStatus("danger", "Wrong password")
	}
}

func (r *apiRequest) Logout() {
	r.session.Values["auth"] = false
	r.session.Options.MaxAge = 0
	r.session.Save(r.request, *r.responseWriter)

	r.setStatus("info", "Good bye!")
}

func (r *apiRequest) Say() {
	text := r.getInData("message")

	//text = tgbotapi.EscapeText("MarkdownV2", text)

	msg := tgbotapi.NewMessage(Global.Settings.BotChatID, text)
	//msg.ParseMode = "MarkdownV2"
	tgBot.Send(msg)

	log.Println("Said: ", text)
}
