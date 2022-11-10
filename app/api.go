package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const sessionName = "mtsession"

var API struct {
	SessionStore *sessions.CookieStore
}

func init() {
	API.SessionStore = sessions.NewCookieStore([]byte("mt-tgadmin super secret key"))
}

func apiSetReply(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func apiCheckSession(w http.ResponseWriter, r *http.Request) bool {
	session, _ := API.SessionStore.Get(r, sessionName)

	if auth, ok := session.Values["auth"].(bool); !ok || !auth {
		http.Error(w, "Auth Required", http.StatusForbidden)
		return false
	} else {
		return true
	}
}

func ApiHealthCheck(w http.ResponseWriter, r *http.Request) {
	session, _ := API.SessionStore.Get(r, sessionName)

	apiSetReply(w, map[string]interface{}{
		"status":  "OK",
		"message": "API working",
		"auth":    session.Values["auth"],
	})
}

func ApiPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	//log.Println(data)

	var status string
	var message string

	if data["password"].(string) == Global.Settings.GuiPassword {
		status = "OK"

		session, _ := API.SessionStore.Get(r, sessionName)

		// Set user as authenticated
		session.Values["auth"] = true
		session.Options.MaxAge = 86400
		session.Save(r, w)
	} else {
		status = "Error"
		message = "Wrong password"
	}

	apiSetReply(w, map[string]interface{}{
		"status":  status,
		"message": message,
	})
}

func ApiLogout(w http.ResponseWriter, r *http.Request) {
	if !apiCheckSession(w, r) {
		return
	}

	session, _ := API.SessionStore.Get(r, sessionName)
	session.Values["auth"] = false
	session.Options.MaxAge = 0
	session.Save(r, w)

	apiSetReply(w, map[string]interface{}{"status": "OK"})
}
