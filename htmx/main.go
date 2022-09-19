package main

import (
	"embed"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var templateFS embed.FS

func timeNow() string {
	return time.Now().Format(time.Kitchen)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		t := template.Must(template.ParseFS(templateFS,
			"templates/index.html",
			"templates/htmx_time.html"))
		t.Execute(c.Writer, gin.H{"ts": timeNow()})
	})

	r.GET("/htmx/time.html", func(c *gin.Context) {
		t := template.Must(template.ParseFS(templateFS,
			"templates/htmx_time.html"))
		t.Execute(c.Writer, gin.H{"ts": timeNow()})
	})

	r.Run()
}
