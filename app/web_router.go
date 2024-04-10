package app

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitoteam/goappbase"
)

// embedded web assets
//
//go:embed assets/*.min.js assets/*.css assets/favicon.ico assets/index.html
var embedFS embed.FS

var webAssetsFS fs.FS

func init() {
	//prepare FS for subdirectory "/assets"
	webAssetsFS, _ = fs.Sub(embedFS, "assets")
}

func BuildWebRouter(r *gin.Engine) {
	//API
	r.POST("/api/*any", WebApiRequestHandler)

	//serve assets
	r.StaticFS("/assets", http.FS(webAssetsFS))

	//serve HTML from templates (just index.html for now)
	t := template.Must(template.New("index").ParseFS(webAssetsFS, "index.html"))
	r.SetHTMLTemplate(t)
	r.GET("/", WebIndex)
}

type indexData struct {
	Global  interface{}
	AppInfo *goappbase.AppBase
	Auth    bool
}

func WebIndex(c *gin.Context) {
	session := sessions.Default(c)

	data := &indexData{
		AppInfo: App,
		Auth:    session.Get("auth") == true,
	}

	c.HTML(http.StatusOK, "index.html", data)
}
