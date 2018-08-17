package main

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

func main() {
	fmt.Println("vim-go")

	var item1, item2 interface{}
	item1 = "hello world"
	item2 = []string{"hello world"}
	h1, _ := hashstructure.Hash(item1, nil)
	h2, _ := hashstructure.Hash("hello world", nil)
	fmt.Println(h1, h2, item2)
}
