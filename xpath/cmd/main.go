package main

import (
	"fmt"
	"log"
	"strings"

	"crawler.club/dl"
	"github.com/antchfx/htmlquery"
)

func main() {
	fmt.Println("vim-go")
	resp := dl.DownloadUrl("http://www.newsmth.net/nForum/article/Sports/387504")
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	doc, err := htmlquery.Parse(strings.NewReader(resp.Text))
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(htmlquery.OutputHTML(doc, true))
	for _, n := range htmlquery.Find(doc, "//table[contains(concat(' ', @class, ' '), ' article ')]//td[contains(concat(' ', @class, ' '), ' a-content ')]") {
		//fmt.Printf("%s \n\n", htmlquery.InnerText(n))
		fmt.Println("=============================")
		fmt.Printf("%s\n", TextFromHTML(htmlquery.OutputHTML(n, true)))
		fmt.Println("=============================\n\n\n")
		//fmt.Printf("%s \n\n\n\n\n", htmlquery.OutputHTML(n, true))
	}
}
