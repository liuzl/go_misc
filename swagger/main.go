package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	_ "swagger_test/docs"

	"github.com/go-chi/chi"
	"github.com/golang/glog"
	httpswagger "github.com/swaggo/http-swagger/v2"
	"zliu.org/goutil/rest"
)

// echo request godoc
// @Summary echo back the http raw request
// @Description echo
// @Tags ops
// @Accept  json
// @Produce  json
// @Param request body interface{} true "raw request"
// @Success 200 {string} string "Raw request"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal server error"
// @Router /echo [post] "echo"
// @Router /echo [get] "echo"
func echo(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		rest.ErrInternalServer(w, fmt.Sprintf("dump err, %v", err))
		return
	}
	glog.Info(string(dump))
	w.Write(dump)
}

// @title Go API
// @version 1.0
// @description This is a sample server golang server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email liang@zliu.org

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:1323
// @BasePath /
func main() {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpswagger.Handler(
		httpswagger.URL("http://127.0.0.1:1323/swagger/doc.json"), //The url pointing to API definition
	))

	r.Handle("/echo", rest.WithLog(echo))

	http.ListenAndServe(":1323", r)
}
