package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/liuzl/goutil/rest"
)

var (
	addr = flag.String("addr", ":8080", "bind address")
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("addr=%s  method=%s host=%s uri=%s",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)

	r.ParseForm()
	var ret []interface{}
	ret = append(ret, r.Form["keys"])
	ret = append(ret, r.FormValue("keys"))
	rest.MustEncode(w, ret)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.Handle("/api", rest.WithLog(ArticleHandler))
	glog.Info("server listen on", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))
}
