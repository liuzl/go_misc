package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var s1 []string = nil
	var s2 []string
	fmt.Printf("len(s1)=%d\n", len(s1))
	fmt.Printf("len(s2)=%d\n", len(s2))
	s1 = append(s1, "ok")

	var m1 map[string]string = nil
	m2 := map[string]string{}
	fmt.Printf("len(m1)=%d\n", len(m1))
	fmt.Printf("len(m2)=%d\n", len(m2))
}
