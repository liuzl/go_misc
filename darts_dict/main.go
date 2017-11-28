package main

import (
	"fmt"
	"github.com/motemen/go-darts"
)

func blanks(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "  "
	}
	return s
}

func main() {
	//d, err := darts.Import("darts.txt", "darts.lib", true)
	d, err := darts.Load("darts.lib")
	if err == nil {
		key := []rune("考察队员")
		if d.ExactMatchSearch(key, 0) {
			fmt.Println("考察队员 is in dictionary")
		}
		key = []rune("中国的李乐是一名考察队员")
		for index, _ := range key {
			sub := key[index:]
			r := d.CommonPrefixSearch(sub, 0)
			for i := 0; i < len(r); i++ {
				fmt.Println(blanks(index) + string(sub[:r[i].PrefixLen]))
			}
		}
	}
}
