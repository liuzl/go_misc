package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

func MakeAbsoluteUrl(href, baseurl string) (string, error) {
	u, err := url.Parse(href)
	if err != nil {
		return "", err
	}
	base, err := url.Parse(baseurl)
	if err != nil {
		return "", err
	}
	u = base.ResolveReference(u)
	return u.String(), nil
}

func main() {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := htmlquery.Parse(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	htmlquery.FindEach(doc, "//div[@class='all']//a", func(i int, node *html.Node) {
		fmt.Println(htmlquery.InnerText(node))
		fmt.Println(htmlquery.SelectAttr(node, "href"))
		fmt.Printf("%+v\n", node.Data)
	})

	expr := xpath.MustCompile("//div[@class='all']//a/@href")
	iter := expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(*xpath.NodeIterator)
	for iter.MoveNext() {
		u, _ := MakeAbsoluteUrl(iter.Current().Value(), "http://pm25.in/")
		fmt.Printf("%s \n", u)
	}
}
