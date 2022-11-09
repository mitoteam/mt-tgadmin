package app

import (
	"embed"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var WebAssets *embed.FS
var WebIndexHtml *string

func BuildWebRouter() *mux.Router {
	router := mux.NewRouter()

	//test
	router.HandleFunc("/api", WebApi)

	//serve assets
	router.PathPrefix("/assets").Handler(http.FileServer(http.FS(WebAssets))).Methods("GET")

	//serve root index.html
	router.HandleFunc("/", WebIndex).Methods("GET")

	return router
}

func WebIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(*WebIndexHtml))
}

func WebApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
