package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("active")
	if err != nil {
		panic(err)
	}
	ret := strings.Fields(string(b))
	fmt.Println(ret)
}
