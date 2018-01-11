package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func hello(s int) {
	fmt.Println(s)
	once.Do(func() {
		fmt.Println("once.Do()")
	})
}

func main() {
	fmt.Println("vim-go")
	for i := 0; i < 10; i++ {
		hello(i)
	}
}
