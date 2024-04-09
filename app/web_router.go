package app

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitoteam/goappbase"
)

var WebAssets *embed.FS
var WebIndexHtml *string

func BuildWebRouter(r *gin.Engine) {
	//API
	r.POST("/api/*any", WebApiRequestHandler)

	//serve assets
	r.StaticFS("/assets", http.FS(WebAssets))

	//serve root index.html
	r.GET("/", WebIndex)
}

type indexData struct {
	Global  interface{}
	AppInfo *goappbase.AppBase
	Auth    bool
}

func WebIndex(c *gin.Context) {
	session := sessions.Default(c)

	t := template.New("index")
	if _, err := t.Parse(*WebIndexHtml); err != nil {
		log.Fatalln(err)
	}

	data := &indexData{
		Global:  Global,
		AppInfo: App,
		Auth:    session.Get("auth") == true,
	}

	if err := t.Execute(c.Writer, data); err != nil {
		log.Fatalln(err)
	}
}
