package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
)

func main() {
	//https://news.google.com/news/headlines?ned=us&hl=en
	//url := "https://www.yahoo.com/news/rss/mostviewed"
	url := "https://news.google.com/news/rss/headlines/section/topic/NATION/?ned=us&hl=en"
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed.Title)
	for _, item := range feed.Items {
		log.Println(item.Title, item.Link)
	}
}
