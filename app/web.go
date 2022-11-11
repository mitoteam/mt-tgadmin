package app

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var WebAssets *embed.FS
var WebIndexHtml *string

func BuildWebRouter() *mux.Router {
	router := mux.NewRouter()

	//API
	router.PathPrefix("/api").HandlerFunc(WebApiRequestHandler)

	//serve assets
	router.PathPrefix("/assets").Handler(http.FileServer(http.FS(WebAssets))).Methods("GET")

	//serve root index.html
	router.HandleFunc("/", WebIndex).Methods("GET")

	return router
}

type indexData struct {
	Global interface{}
	Auth   bool
}

func WebIndex(w http.ResponseWriter, r *http.Request) {
	session, _ := apiSessionStore.Get(r, sessionName)

	t := template.New("index")
	if _, err := t.Parse(*WebIndexHtml); err != nil {
		log.Fatalln(err)
	}

	data := &indexData{
		Global: Global,
		Auth:   session.Values["auth"] == true,
	}

	if err := t.Execute(w, data); err != nil {
		log.Fatalln(err)
	}
}
