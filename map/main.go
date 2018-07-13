package main

import (
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/hashstructure"
)

type A struct {
	S string
}

type B struct {
	Arr []*A
}

func main() {
	a := &B{[]*A{&A{"a"}}}
	b := &B{[]*A{&A{"a"}}}

	fmt.Println(a, b)
	fmt.Printf("a==b: %+v\n", a == b)
	fmt.Printf("reflect.DeepEqual(a, b): %+v\n", reflect.DeepEqual(a, b))
	fmt.Printf("cmp.Equal(a, b): %+v\n", cmp.Equal(a, b))

	ha, _ := hashstructure.Hash(a, nil)
	hb, _ := hashstructure.Hash(b, nil)
	fmt.Println(ha, hb)
}
