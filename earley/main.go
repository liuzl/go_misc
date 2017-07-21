package main

import "fmt"

func main() {
	items := []int{0, 1, 2, 1}
	fmt.Println(items)
	for i := 0; i < len(items); i++ {
		if items[i] == 1 {
			items = append(items, 10)
		}
	}
	fmt.Println(items)
}
