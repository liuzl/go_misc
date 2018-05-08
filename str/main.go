package main

import (
	"fmt"
	"strings"
)

func ReverseOrderEncode(s string) string {
	var buf []rune
	for _, c := range s {
		if c >= '0' && c <= '9' {
			buf = append(buf, '9'-c+'0')
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}

func main() {
	fmt.Println("vim-go")
	dot := "\u00B7"
	fmt.Println(dot)
	fmt.Println("\u554a")
	str := "Jörgensen"
	l := strings.ToLower(str)
	u := strings.ToUpper(l)
	fmt.Println(str, l, u)

	s := "2018050716"
	fmt.Println(s)
	fmt.Println(ReverseOrderEncode(s))
	fmt.Println(ReverseOrderEncode("hello90"))

	s = "http://pm25.in/beijing"
	fmt.Println(strings.Trim(s, "http://pm25.in/"))
}
