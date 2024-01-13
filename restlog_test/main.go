package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/golang/glog"
	"zliu.org/goutil/rest"
)

var (
	addr = flag.String("addr", ":9001", "bind address")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")

	http.Handle("/echo", rest.WithLog(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			rest.ErrInternalServer(w, fmt.Sprintf("dump err, %v", err))
			return
		}
		glog.Info(string(dump))
		w.Write(dump)
	}))

	go func() {
		for {
			glog.Info("hello")
			rest.Log().Info().Msg("GO FUNC")
			time.Sleep(time.Second * 5)
		}
	}()

	glog.Info("server listen on ", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))
}
