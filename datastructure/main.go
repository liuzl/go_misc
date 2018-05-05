package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var s1 []string = nil
	var s2 []string
	fmt.Printf("len(s1)=%d\n", len(s1))
	fmt.Printf("len(s2)=%d\n", len(s2))
	s1 = append(s1, "ok")
}
