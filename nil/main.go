package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var x interface{}
	fmt.Println(x)
	v, ok := x.(string)
	fmt.Println(v, ok)
}
