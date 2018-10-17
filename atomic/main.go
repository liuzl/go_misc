package main

import (
	"fmt"
	"sync/atomic"
)

var n uint64

func main() {
	fmt.Println("vim-go")
	for i := 0; i < 100; i++ {
		atomic.AddUint64(&n, 10)
	}
	fmt.Println(n)
}
