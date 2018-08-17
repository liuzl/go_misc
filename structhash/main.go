package main

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

type pair struct {
	a string
	b string
}

func main() {
	fmt.Println("vim-go")

	var item1, item2 interface{}
	var item3 map[string]string
	var item4 map[uint64]string
	item1 = "hello world"
	h1, _ := hashstructure.Hash(item1, nil)
	h2, _ := hashstructure.Hash("hello world", nil)
	h3, _ := hashstructure.Hash(item2, nil)
	h4, _ := hashstructure.Hash(item3, nil)
	h5, _ := hashstructure.Hash(item4, nil)
	fmt.Println(h1, h2, item2, h3, h4, h5)

	x := &pair{"a", "b"}
	y := pair{"a", "b"}
	hx, _ := hashstructure.Hash(x, nil)
	hy, _ := hashstructure.Hash(y, nil)
	fmt.Println(hx, hy)
}
