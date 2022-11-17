package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gorilla/sessions"
)

const sessionName = "mtsession"

var apiSessionStore *sessions.CookieStore
var tgBot *tgbotapi.BotAPI

func InitApi() bool {
	if len(Global.Settings.GuiCookieSecretKey) < 32 {
		log.Println("gui_cookie_secret_key should be at least 32 characters long")
		return false
	}

	apiSessionStore = sessions.NewCookieStore([]byte(Global.Settings.GuiCookieSecretKey))

	if len(Global.Settings.BotToken) == 0 {
		log.Println("bot_token required")
		return false
	}

	var err error
	tgBot, err = tgbotapi.NewBotAPI(Global.Settings.BotToken)
	if err != nil {
		log.Println(err)
		return false
	}

	Global.BotInfo = tgBot.Self.FirstName + " " + tgBot.Self.LastName + " (@" + tgBot.Self.UserName + ")"
	log.Printf("Authorized on telegram bot: %s\n", Global.BotInfo)

	if Global.Settings.BotChatID == 0 {
		log.Println("bot_chat_id required")
		return false
	}

	chat, err := tgBot.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: Global.Settings.BotChatID}})
	if err != nil {
		log.Println(err)
		return false
	}

	Global.ChatInfo = chat.Type + " \"" + chat.Title + "\", " + chat.InviteLink
	log.Printf("Chat info: %s\n", Global.ChatInfo)

	return true
}

func WebApiRequestHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")
	//log.Println(path)

	api_request, err := NewApiRequest(&w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//run request handler
	api_request.Run(path)

	//prepare reply
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(api_request.outData)
}

// #region API object
type apiRequest struct {
	inData  map[string]interface{}
	outData map[string]interface{}
	session *sessions.Session

	responseWriter *http.ResponseWriter
	request        *http.Request
}

func NewApiRequest(responseWriter *http.ResponseWriter, request *http.Request) (*apiRequest, error) {
	r := &apiRequest{
		inData:         make(map[string]interface{}),
		outData:        make(map[string]interface{}),
		responseWriter: responseWriter,
		request:        request,
	}

	//prepare session
	r.session, _ = apiSessionStore.Get(request, sessionName)

	//prepare input data
	body, err := io.ReadAll(io.LimitReader(r.request.Body, 1048576))
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
	if auth, ok := r.session.Values["auth"].(bool); !ok || !auth {
		r.setError("Auth Required")
		return false
	} else {
		return true
	}
}

func (r *apiRequest) setError(message string) {
	r.setStatus("danger", message)
	http.Error(*r.responseWriter, message, http.StatusInternalServerError)
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
