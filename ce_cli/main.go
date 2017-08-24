package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Obaied/RAKE.Go"
	"github.com/crawlerclub/ce"
	"github.com/crawlerclub/dl"
	"github.com/golang/glog"
	"github.com/urandom/text-summary/summarize"
	"strings"
)

var (
	url = flag.String("url", "http://news.qq.com/a/20170802/045735.htm", "url to get")
)

func run() {
	req := &dl.HttpRequest{Url: *url, Method: "GET", UseProxy: false, Platform: "mobile"}
	res := dl.Download(req)
	if res.Error != nil {
		glog.Error(res.Error)
		return
	}

	items := strings.Split(res.RemoteAddr, ":")
	ip := ""
	if len(items) > 0 {
		ip = items[0]
	}
	doc := ce.ParsePro(*url, res.Text, ip, true)
	s := summarize.NewFromString(doc.Title, doc.Text)
	doc.Debug["sum"] = s.KeyPoints()
	doc.Debug["rake"] = rake.RunRake(doc.Text)
	j, _ := json.Marshal(doc)
	fmt.Println(string(j))
}

func main() {
	flag.Parse()
	defer glog.Flush()
	run()
}
