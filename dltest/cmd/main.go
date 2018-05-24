package main

import (
	"crawler.club/dl"
	"fmt"
)

func main() {
	url := "http://127.0.0.1:8302/echo?a=asdf"
	resp := dl.DownloadUrl(url)
	fmt.Println(resp.Text)
}
