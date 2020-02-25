package main

import (
	"flag"
	"net/http"
	"net/http/httputil"

	"github.com/golang/glog"
	"github.com/liuzl/goutil/rest"
)

var (
	serverAddr = flag.String("addr", ":8080", "bind address")
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		rest.MustEncode(w, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "error", Message: err.Error()})
		return
	}
	w.Write(dump)
}

func IpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.RemoteAddr))
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.HandleFunc("/echo/", EchoHandler)
	http.HandleFunc("/ip/", IpHandler)
	glog.Info("server listen on", *serverAddr)
	glog.Error(http.ListenAndServe(*serverAddr, nil))
}
