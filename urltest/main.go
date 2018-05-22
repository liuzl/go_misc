package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
)

var (
	u = flag.String("url", "./", "url")
)

func main() {
	flag.Parse()
	ret, err := url.Parse(*u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret.IsAbs())
}
