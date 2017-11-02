package main

import (
	"flag"
	"fmt"
	"github.com/gpahal/go-meta"
)

var (
	url = flag.String("url", "http://tech.ifeng.com/a/20170910/44678332_0.shtml", "url to parse")
)

func main() {
	flag.Parse()
	object, _ := meta.ProcessURL(*url, nil)
	j, _ := object.JSON()
	fmt.Println(string(j))
}
