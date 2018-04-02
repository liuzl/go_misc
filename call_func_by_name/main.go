package main

import (
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

func Call(m map[string]interface{},
	name string, params ...interface{}) ([]reflect.Value, error) {
	var nf interface{}
	if nf = m[name]; nf == nil {
		return nil, fmt.Errorf("func %s not found", name)
	}
	f := reflect.ValueOf(nf)
	if f.Kind() != reflect.Func {
		return nil, fmt.Errorf("%s is not a function", name)
	}
	t := f.Type()
	if len(params) != t.NumIn() {
		return nil, fmt.Errorf("(len(params)=%d) != (t.NumIn()=%d)",
			len(params), t.NumIn())
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
		if in[k].Type() != t.In(k) {
			return nil, fmt.Errorf("(in[%d].Type()=%s) != (t.In(%d)=%s)",
				k, in[k].Type(), k, t.In(k))
		}
	}
	return f.Call(in), nil
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
		for i, v := range ret {
			fmt.Printf("%d: %+v\n", i, v)
		}
	}

	_, err = Call(funcs, "sum", []string{"string"})
	if err != nil {
		fmt.Println(err)
	}
}
