package main

import (
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Set("require", require)
	vm.Run(`
        require("test.js");
        test();
		abc();
    `)
}

func require(call otto.FunctionCall) otto.Value {
	file := call.Argument(0).String()
	fmt.Printf("requiring: %s\n", file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = call.Otto.Run(string(data))
	fmt.Println(string(data))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return otto.TrueValue()
}
