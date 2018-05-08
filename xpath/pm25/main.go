package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"github.com/liuzl/goutil"
	"github.com/tkuchiki/parsetime"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

var keys = []string{"position_name", "aqi", "quality", "primary_pollutant",
	"pm2_5", "pm10", "co", "no2", "o3", "o3_8h", "so2"}

func main() {
	detail()
}

func detail() {
	data, err := ioutil.ReadFile("beijing.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := htmlquery.Parse(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	node := htmlquery.CreateXPathNavigator(doc)
	expr := xpath.MustCompile(`//div[@class="live_data_time"]/p/text()`)
	iter := expr.Evaluate(node).(*xpath.NodeIterator)
	var arr []rune
	for iter.MoveNext() {
		for _, c := range iter.Current().Value() {
			if unicode.IsDigit(c) {
				arr = append(arr, c)
			}
		}
	}
	tstr := string(arr)
	p, err := parsetime.NewParseTime("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	t, err := p.Parse(tstr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tstr, t)

	city := xpath.MustCompile(`//h2/text()`)
	iter = city.Evaluate(node).(*xpath.NodeIterator)
	for iter.MoveNext() {
		fmt.Println(iter.Current().Value())
	}

	rows := xpath.MustCompile(`//table[@id='detail-data']//tr`)
	td := xpath.MustCompile(`td`)
	iter = rows.Evaluate(node).(*xpath.NodeIterator)
	for i := 0; iter.MoveNext(); i++ {
		if i == 0 {
			continue
		}
		it := td.Evaluate(iter.Current()).(*xpath.NodeIterator)
		for j := 0; it.MoveNext(); j++ {
			if j < len(keys) {
				fmt.Printf("%s:%s\n", keys[j], it.Current().Value())
			}
		}
	}
}

func list() {
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
		u, _ := goutil.MakeAbsoluteUrl(iter.Current().Value(), "http://pm25.in/")
		fmt.Printf("%s \n", u)
	}
	detail()
}
