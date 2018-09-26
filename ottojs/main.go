package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/robertkrimen/otto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <jsfile>\n", os.Args[0])
		os.Exit(1)
	}

	vm := otto.New()
	vm.Set("require", require)
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = vm.Run(b)
	fmt.Println(err)
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
