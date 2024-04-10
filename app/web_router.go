package app

import (
	"html/template"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitoteam/goappbase"
)

func BuildWebRouter(r *gin.Engine) {
	//API
	r.POST("/api/*any", WebApiRequestHandler)

	//serve assets
	r.StaticFS("/assets", http.FS(webAssetsFS))

	//serve HTML from templates (just index.html for now)
	t := template.Must(template.New("index").ParseFS(templatesFS, "index.html"))
	r.SetHTMLTemplate(t)
	r.GET("/", WebIndex)
}

func WebIndex(c *gin.Context) {
	session := sessions.Default(c)

	//c.HTML(http.StatusOK, "index.html", data)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"AppInfo": App,
		"Auth":    session.Get("auth") == true,
	})
}
