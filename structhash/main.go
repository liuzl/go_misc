package main

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

type pair struct {
	a string
	b string
}

type wordValue struct {
	typ string
	val interface{}
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

	shs := &wordValue{"loc_county", []string{"310201", "310221"}}
	sh := &wordValue{"loc_province", []string{"310000"}}
	k1, e1 := hashstructure.Hash(shs, nil)
	k2, e2 := hashstructure.Hash(sh, nil)
	k3, e3 := hashstructure.Hash(wordValue{}, nil)
	fmt.Println(k1, k2, k3, e1, e2, e3)
}
