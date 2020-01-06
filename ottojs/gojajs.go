package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <jsfile>\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println("running", os.Args[1])
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	registry := new(require.Registry) // this can be shared by multiple runtimes
	vm := goja.New()
	registry.Enable(vm)
	console.Enable(vm)

	if _, err = vm.RunString(string(b)); err != nil {
		fmt.Println(err)
	}

	//m, err := req.Require("m.js")
	//_, _ = m, err
}
