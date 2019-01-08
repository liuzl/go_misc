package main

import (
	"fmt"

	"github.com/liuzl/goutil"
	//"github.com/m3ng9i/feedreader"
	//"github.com/mmcdole/gofeed"
	"github.com/xgolib/gofeed"
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

	//feed1, _ := feedreader.Fetch(url)
	b, _ := goutil.JSONMarshal(feed)
	fmt.Println(string(b))
}
