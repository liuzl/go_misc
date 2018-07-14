package main

import (
	"encoding/json"
	"fmt"
)

type TermType byte

const (
	EOF TermType = iota
	Nonterminal
	Terminal
	Any
)

func main() {
	fmt.Println("vim-go", Any)
	b, err := json.Marshal(Any)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	str := "\"Any\""
	var t TermType
	fmt.Println(t)
	err = json.Unmarshal([]byte(str), &t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
}
