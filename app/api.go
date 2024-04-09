package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var tgBot *tgbotapi.BotAPI

func InitApi() error {
	if len(Settings.GuiCookieSecretKey) < 32 {
		return errors.New("gui_cookie_secret_key should be at least 32 characters long")
	}

	if len(Settings.BotToken) == 0 {
		return errors.New("bot_token required")
	}

	var err error
	tgBot, err = tgbotapi.NewBotAPI(Settings.BotToken)
	if err != nil {
		return err
	}

	Global.BotInfo = tgBot.Self.FirstName + " " + tgBot.Self.LastName + " (@" + tgBot.Self.UserName + ")"
	log.Printf("Authorized on telegram bot: %s\n", Global.BotInfo)

	if Settings.BotChatID == 0 {
		return errors.New("bot_chat_id required")
	}

	chat, err := tgBot.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: Settings.BotChatID}})
	if err != nil {
		return err
	}

	Global.ChatInfo = chat.Type + " \"" + chat.Title + "\", " + chat.InviteLink
	log.Printf("Chat info: %s\n", Global.ChatInfo)

	return nil
}

func WebApiRequestHandler(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.URL.Path, "/api")
	//log.Println(path)

	api_request, err := NewApiRequest(c)
	if err != nil {
		log.Println(err)
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	//run request handler
	api_request.Run(path)

	//prepare reply
	c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(c.Writer).Encode(api_request.outData)
}

// #region API object
type apiRequest struct {
	inData  map[string]interface{}
	outData map[string]interface{}
	session sessions.Session

	context *gin.Context
}

func NewApiRequest(c *gin.Context) (*apiRequest, error) {
	r := &apiRequest{
		inData:  make(map[string]interface{}),
		outData: make(map[string]interface{}),
		context: c,
	}

	//prepare session
	r.session = sessions.Default(c)

	//prepare input data
	body, err := io.ReadAll(io.LimitReader(c.Request.Body, 1048576))
	if err != nil {
		return nil, err
	}
	//log.Println(string(body))

	if json.Valid(body) {
		if err := json.Unmarshal(body, &r.inData); err != nil {
			return nil, err
		}
	}
	//log.Println(r.inData)

	return r, nil
}

func (r *apiRequest) getInData(name string) string {
	if value, ok := r.inData[name]; ok {
		if _, ok := value.(string); ok {
			return value.(string)
		} else {
			return fmt.Sprintf("%v", value)
		}
	} else {
		return ""
	}
}

func (r *apiRequest) getInDataInt(name string, default_value int) int {
	if value, ok := r.inData[name]; ok {
		//log.Println(reflect.TypeOf(value))

		if _, ok := value.(int); ok {
			return value.(int)
		}

		if _, ok := value.(float64); ok {
			return int(value.(float64))
		}
	}

	return default_value
}

func (r *apiRequest) getOutData(name string) string {
	if value, ok := r.outData[name]; ok {
		return value.(string)
	} else {
		return ""
	}
}

func (r *apiRequest) setOutData(name string, value interface{}) {
	r.outData[name] = value
}

func (r *apiRequest) setStatus(status, message string) {
	r.setOutData("status", status)
	r.setOutData("message", message)
}

func (r *apiRequest) apiCheckSession() bool {
	if r.session.Get("auth") == nil {
		r.setError("Auth Required")
		return false
	} else {
		return true
	}
}

func (r *apiRequest) setError(message string) {
	r.setStatus("danger", message)
	http.Error(r.context.Writer, message, http.StatusInternalServerError)
}

//#endregion

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
