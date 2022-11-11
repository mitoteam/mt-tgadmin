package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

const sessionName = "mtsession"

var apiSessionStore *sessions.CookieStore

func init() {
	apiSessionStore = sessions.NewCookieStore([]byte("mt-tgadmin super secret key"))
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

// region API object
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
		if err := json.Unmarshal(body, &r.outData); err != nil {
			return nil, err
		}
	}
	//log.Println(r.outData)

	return r, nil
}

func (r *apiRequest) getInData(name string) string {
	if value, ok := r.outData[name]; ok {
		return value.(string)
	} else {
		return ""
	}
}

func (r *apiRequest) getOutData(name string) string {
	if value, ok := r.outData[name]; ok {
		return value.(string)
	} else {
		return ""
	}
}

func (r *apiRequest) setOutData(name, value string) {
	r.outData[name] = value
}

func (r *apiRequest) setStatus(status, message string) {
	r.setOutData("status", status)
	r.setOutData("message", message)
}

func (r *apiRequest) apiCheckSession() bool {
	if auth, ok := r.session.Values["auth"].(bool); !ok || !auth {
		r.setStatus("danger", "Auth Required")
		http.Error(*r.responseWriter, "Auth Required", http.StatusForbidden)
		return false
	} else {
		return true
	}
}

//endregion
