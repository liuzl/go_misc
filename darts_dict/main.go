package main

import (
	"fmt"
	"github.com/motemen/go-darts"
)

func main() {
	d, err := darts.Import("darts.txt", "darts.lib", false)
	if err == nil {
		key := []rune("考察队员")
		if d.ExactMatchSearch(key, 0) {
			fmt.Println("考察队员 is in dictionary")
		}
		key = []rune("lzl是一个考察队员")
		r := d.CommonPrefixSearch(key, 0)
		for i := 0; i < len(r); i++ {
			fmt.Println(string(key[:r[i].PrefixLen]))
		}
	}
}
