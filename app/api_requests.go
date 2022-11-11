package app

import (
	"fmt"
	"log"
	"strings"
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

			default:
				message := "Unknown API request: " + path
				log.Println(message)
				r.setStatus("danger", message)
			}
		}
	}

	if r.getInData("status") == "" {
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
