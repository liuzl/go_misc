package main

import (
	"fmt"
	"github.com/rs/xid"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(xid.New().String())
	fmt.Println(xid.New().String())
	fmt.Println(xid.New().String())
}
