package main

import (
	"errors"
	"fmt"
	"reflect"
)

func foo() {
	fmt.Println("in function foo")
}

func bar(i1, i2, i3 int) {
	fmt.Printf("in func bar, i1:%d, i2:%d, i3:%d\n", i1, i2, i3)
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	nf, ok := m[name]
	if !ok {
		err = errors.New("func " + name + " not found")
		return
	}
	f := reflect.ValueOf(nf)
	fmt.Printf("len(params)=%d, f.Type().NumIn()=%d\n", len(params), f.Type().NumIn())
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func main() {
	funcs := map[string]interface{}{
		"foo": foo,
		"bar": bar,
		"sum": sum,
	}
	ret, err := Call(funcs, "sum", []int{1, 2, 3})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}
