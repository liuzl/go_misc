package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"log"
)

func main() {
	box, err := rice.FindBox("data")
	if err != nil {
		log.Fatal(err)
	}
	s, err := box.String("aes.js")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
