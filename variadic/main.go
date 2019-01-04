package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(list(1, "abc", "hello", 1.8))
}

func list(items ...interface{}) []interface{} {
	return items
}
