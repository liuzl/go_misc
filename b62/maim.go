package main

import (
	"fmt"

	"github.com/wcalandro/base62"
)

func main() {
	fmt.Println("vim-go")
	var i uint64 = 123455678
	fmt.Println(base62.ToB62(i))
}
