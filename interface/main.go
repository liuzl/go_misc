package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var void interface{}
	void = nil
	switch void.(type) {
	case []string:
		fmt.Println("void is []string")
	default:
		fmt.Println("default")
	}
}
