package main

import (
	"fmt"
	"github.com/gpahal/go-meta"
)

func main() {
	fmt.Println("vim-go")
	object, err := meta.ProcessURL("http://ogp.me", nil)
	fmt.Println(object)
	fmt.Println(err)
}
