package main

import (
	"fmt"
)

func main() {
	sliceInt := []int{11, 2, 19, 220, 31, 5, 65, 70, 100}
	find := 65
	fmt.Println(sliceInt)
	if len((sliceInt)) == 0 || (sliceInt)[0] == find {
		fmt.Println(sliceInt)
		return
	}
	if (sliceInt)[len(sliceInt)-1] == find {
		(sliceInt) = append([]int{find}, (sliceInt)[:len(sliceInt)-1]...)
		fmt.Println(sliceInt)
		return
	}
	for p, x := range sliceInt {
		if x == find {
			(sliceInt) = append([]int{find}, append((sliceInt)[:p], (sliceInt)[p+1:]...)...)
			break
		}
	}
	fmt.Println(sliceInt)
}
