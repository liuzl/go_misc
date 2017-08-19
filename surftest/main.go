package main

import (
	"flag"
	"fmt"
	"gopkg.in/headzoo/surf.v1"
)

var (
	link = flag.String("url", "http://m.newsmth.net", "url to get")
)

func main() {
	flag.Parse()
	bow := surf.NewBrowser()
	err := bow.Open(*link)
	if err != nil {
		panic(err)
	}
	fmt.Println(bow.Title())
	fmt.Println(bow.Body())
}
