package app

import (
	"errors"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var tgBot *tgbotapi.BotAPI

func InitTgApi() error {
	if len(Settings.BotToken) == 0 {
		return errors.New("bot_token required")
	}

	var err error
	tgBot, err = tgbotapi.NewBotAPI(Settings.BotToken)
	if err != nil {
		return err
	}

	App.Global["BotInfo"] = tgBot.Self.FirstName + " " + tgBot.Self.LastName + " (@" + tgBot.Self.UserName + ")"
	log.Printf("Authorized on telegram bot: %s\n", App.Global["BotInfo"].(string))

	if Settings.BotChatID == 0 {
		return errors.New("bot_chat_id required")
	}

	chat, err := tgBot.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: Settings.BotChatID}})
	if err != nil {
		return err
	}

	App.Global["ChatInfo"] = chat.Type + " \"" + chat.Title + "\", " + chat.InviteLink
	log.Printf("Chat info: %s\n", App.Global["ChatInfo"].(string))

	return nil
}

// See https://core.telegram.org/bots/api#formatting-options
func PrepareTelegramHtml(input string) (r string) {
	r = input

	r = strings.ReplaceAll(r, "&nbsp;", " ")

	r = strings.ReplaceAll(r, "<br>", "\n")
	r = strings.ReplaceAll(r, "<p>", "")
	r = strings.ReplaceAll(r, "</p>", "\n")

	r = strings.TrimSpace(r)

	return r
}
