package app

import (
	"html/template"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitoteam/goappbase"
)

func BuildWebRouter(r *gin.Engine) {
	//serve assets
	r.StaticFS("/assets", http.FS(webAssetsFS))

	//serve HTML from templates (just index.html for now)
	t := template.Must(template.New("").ParseFS(templatesFS, "index.html"))
	r.SetHTMLTemplate(t)
	r.GET("/", WebIndex)
}

func WebIndex(c *gin.Context) {
	session := sessions.Default(c)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"AppInfo": App,
		"Auth":    session.Get("auth") == true,
	})
}

// builds API routing and handlers for goappbase
func BuildWebApiRouter(app *goappbase.AppBase) {
	app.WebApiPathPrefix = "/api"
	app.WebApiEnableGet = !Settings.Production // in production mode only

	app.
		ApiHandler("/ping", Api_HealthCheck).
		ApiHandler("/password", Api_Password).
		ApiHandler("/logout", Api_Logout).
		ApiHandler("/say", Api_Say).
		ApiHandler("/list_messages", Api_ListMessages)
}
