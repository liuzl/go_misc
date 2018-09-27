package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/liuzl/goutil/rest"
	"github.com/skip2/go-qrcode"
)

var (
	addr = flag.String("addr", ":8080", "bind address")
)

func QrHandler(w http.ResponseWriter, r *http.Request) {
	ts := time.Now().Format(time.RFC3339)
	var png []byte
	png, err := qrcode.Encode(ts, qrcode.Medium, 256)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	img := base64.StdEncoding.EncodeToString(png)
	w.Write([]byte(fmt.Sprintf(
		"<html><img src=\"data:image/png;base64,%s\">", img)))
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.Handle("/qr.html", rest.WithLog(QrHandler))
	glog.Info("server listen on", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))

}
