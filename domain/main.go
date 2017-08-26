package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/joeguo/tldextract"
)

var (
	url = flag.String("url", "http://zliu.org", "url to parse")
)

func main() {
	flag.Parse()
	tldExtractor, _ := tldextract.New("./tld.cache", false)
	result := tldExtractor.Extract(*url)
	j, _ := json.Marshal(result)
	fmt.Println(string(j))
}
