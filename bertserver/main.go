package main

import (
	"context"
	"embed"
	"flag"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/text2text"
	"zliu.org/goutil/rest"
)

var (
	addr      = flag.String("addr", ":8300", "bind address")
	modelsDir = flag.String("dir", "models", "the models dir")
	modelName = flag.String("name", "Helsinki-NLP/opus-mt-en-zh", "the translation model name")
)

type TranslationResult struct {
	Input  string  `json:"input"`
	Output string  `json:"output"`
	Time   float64 `json:"time"`
}

//go:embed static
var staticFS embed.FS

func main() {
	flag.Parse()
	m, err := tasks.Load[text2text.Interface](&tasks.Config{ModelsDir: *modelsDir, ModelName: *modelName})
	if err != nil {
		glog.Fatal(err)
	}
	defer tasks.Finalize(m)

	opts := text2text.DefaultOptions()

	http.Handle("/translate", rest.WithLog(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		text := strings.TrimSpace(r.FormValue("text"))
		if text == "" {
			rest.MustEncode(w, &rest.RestMessage{Status: "fail", Message: "input text is empty"})
			return
		}
		start := time.Now()
		result, err := m.Generate(context.Background(), text, opts)
		if err != nil {
			rest.MustEncode(w, &rest.RestMessage{Status: "fail", Message: err.Error()})
			glog.Error(err)
			return
		}
		ret := &TranslationResult{
			text,
			result.Texts[0],
			time.Since(start).Seconds(),
		}
		rest.MustEncode(w, &rest.RestMessage{Status: "ok", Message: ret})
	}))

	htmlContent, err := fs.Sub(fs.FS(staticFS), "static")
	if err != nil {
		glog.Error(err)
		return
	}
	http.Handle("/", http.FileServer(http.FS(htmlContent)))

	glog.Info("server listen on http://0.0.0.0", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))
}
