package main

import (
	"fmt"
	"strings"

	"github.com/bytedance/gopkg/lang/stringx"
)

func main() {
	s1 := "2020-01-21 00:00:00"
	s2 := "2020-01-21"
	s3 := "2020-01-22"
	fmt.Println(s1 > s2)
	fmt.Println(s1 > s3)

	var x []string
	fmt.Println(x, len(x))
	x = strings.Split("", ",")
	fmt.Println(x, len(x))
	x = strings.Split("a", ",")
	fmt.Println(x, len(x))
	x = strings.Split("a,b", ",")
	fmt.Println(x, len(x))

	fmt.Printf("PadLeft=[%s]\n", stringx.PadLeftSpace("abc", 7))
}
