package main

import "fmt"

var Funcs = make(map[string]func(...interface{}) interface{})

func init() {
	Funcs["nf.math.sum"] = func(a ...interface{}) interface{} {
		return 42
	}
}

func main() {
	f := Funcs["nf.math.sum"]
	if f != nil {
		fmt.Println(f(3, 4))
	}
}
