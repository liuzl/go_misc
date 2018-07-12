package main

import "fmt"

type Values map[string]interface{}

func main() {
	var m interface{}
	m = map[string]interface{}{"m1": "vm1"}

	var v interface{}
	v = Values{"m1": "vm1"}

	fmt.Println(m, v)

	//x := v.(map[string]interface{})
	x := v.(Values)
	fmt.Println(x)
}
