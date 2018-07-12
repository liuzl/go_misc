package main

import (
	"fmt"
	"reflect"
)

type A struct {
	S string
}

type B struct {
	Arr []*A
}

func main() {
	a := &B{[]*A{&A{"a"}}}
	b := &B{[]*A{&A{"b"}}}

	fmt.Println(a, b)
	fmt.Printf("a==b: %+v\n", a == b)
	fmt.Printf("reflect.DeepEqual(a, b): %+v\n", reflect.DeepEqual(a, b))
}
