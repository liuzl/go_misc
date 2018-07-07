package main

import (
	"fmt"
	"regexp"
)

var (
	LinkPattern = `(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w\.-]*)*\/?`

	text = `https://crawler.club是爬虫的主页哈哈`
)

func main() {
	r := regexp.MustCompile(LinkPattern)
	fmt.Println(r.FindAllString(text, -1))
}
