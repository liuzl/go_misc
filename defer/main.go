package main

import "fmt"

var s = "1"

func main() {
	v := "one"
	defer deferRun(v)
	defer deferRun(s)
	v = "two"
	v = "three"
	s = "2"
}

func deferRun(v string) {
	fmt.Println(v)
}
