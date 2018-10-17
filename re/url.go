package main

import (
	"fmt"
	"regexp"
)

const LinkPattern = `^(?i)(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w\.-]*)*\/?`

var linkRegex = regexp.MustCompile(LinkPattern)

func main() {
	url := "https://zliu.org/abc.html?a=b"
	fmt.Println(linkRegex.MatchString(url))
}
