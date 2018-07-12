package main

import (
	"flag"
	"log"
)

var (
	dir = flag.String("dir", "data", "data dir")
)

func main() {
	flag.Parse()
	d, err := Load(*dir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(d)
}
