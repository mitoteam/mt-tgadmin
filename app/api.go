package app

import (
	"encoding/json"
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

func ApiHealthCheck(w http.ResponseWriter, r *http.Request) {
	reply := map[string]interface{}{
		"status":  "OK",
		"message": "API working",
	}

	json.NewEncoder(w).Encode(reply)
}

func ApiPassword(w http.ResponseWriter, r *http.Request) {
	status := "OK"

	session, _ := API.SessionStore.Get(r, sessionName)

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["auth"] = true
	session.Options.MaxAge = 86400
	session.Save(r, w)

	reply := map[string]interface{}{
		"status":  status,
		"session": session.Name(),
	}

	json.NewEncoder(w).Encode(reply)
}
