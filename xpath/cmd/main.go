package main

import (
	"crawler.club/dl"
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	resp := dl.DownloadUrl("http://www.newsmth.net/nForum/article/Taiwan/50328")
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	doc, err := htmlquery.Parse(strings.NewReader(resp.Text))
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(htmlquery.OutputHTML(doc, true))
	//for _, n := range htmlquery.Find(doc, "//table[contains(concat(' ', @class, ' '), ' article ')]//td[contains(concat(' ', @class, ' '), ' a-content ')]") {
	for _, n := range htmlquery.Find(doc, "//table[contains(concat(' ', @class, ' '), ' article ')]") { //span[contains(@class, 'a-pos')]") {
		fmt.Printf("%s \n\n\n\n\n", htmlquery.OutputHTML(n, true))
	}
}
