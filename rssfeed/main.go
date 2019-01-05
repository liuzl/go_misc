package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	url := "https://www.gamblinginsider.com/feed/index.xml"
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	fmt.Println(len(feed.Items))
	for _, i := range feed.Items {
		fmt.Println(i.Title, i.Published, i.PublishedParsed)
		fmt.Println(i.Link)
	}
}
