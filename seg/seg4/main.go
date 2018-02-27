package main

import (
	"fmt"
	"github.com/liuzl/tokenizer"
)

func main() {
	c := `Life is like a box of chocolates. You never know what you're gonna get.`
	var ret []string = tokenizer.Tokenize(c)
	for _, term := range ret {
		fmt.Println(term)
	}
}
