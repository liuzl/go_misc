package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"zliu.org/goutil"
)

var (
	url = flag.String("url", "http://www.qq.com/", "url to fetch and parse")
)

func main() {
	doc, err := htmlquery.LoadURL(*url)
	if err != nil {
		panic(err)
	}
	var links []string
	htmlquery.FindEach(doc, "//a", func(i int, node *html.Node) {
		link := htmlquery.SelectAttr(node, "href")
		if u, err := goutil.MakeAbsoluteUrl(link, *url); err == nil {
			if strings.HasPrefix(u, "http") && !strings.HasSuffix(u, ".exe") {
				links = append(links, u)
			}
		}
	})
	for i, link := range links {
		fmt.Println(i, link)
	}
}
