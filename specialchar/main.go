package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

func main() {
	b, _ := ioutil.ReadFile("body.txt")
	for _, r := range []rune(string(b)) {
		if unicode.IsControl(r) && r != '\n' {
			fmt.Println("control", strconv.Quote(string(r)))
		}
	}
}
