package main

import (
	"flag"
	"fmt"
)

var (
	input = flag.String("i", "default input", "input text")
)

func main() {
	flag.Parse()
	r := []rune(*input)
	fmt.Println(*input)
	fmt.Println(r)
}
