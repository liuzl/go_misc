package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\d+(?:st|nd|rd|th)`)
	src := "December 19th, 2018"
	dst := re.ReplaceAllStringFunc(src, func(in string) string {
		return in[:len(in)-2]
	})
	fmt.Println(dst)
}
