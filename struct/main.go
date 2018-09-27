package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name string `json:"name"`
}

type Dog struct {
	*Animal
	Owner string `json:"owner"`
}

func main() {
	a := &Animal{Name: "this is an animal"}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))

	d := &Dog{a, "this is the owner name of this dog"}
	b, _ = json.Marshal(d)
	fmt.Println(string(b))

	b, _ = json.Marshal(nil)
	fmt.Println(string(b))
}
