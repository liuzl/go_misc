package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	dot := "\u00B7"
	fmt.Println(dot)
	fmt.Println("\u554a")
	str := "Jörgensen"
	l := strings.ToLower(str)
	u := strings.ToUpper(l)
	fmt.Println(str, l, u)
}
