package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(file_line())
}

func file_line() string {
	_, fileName, fileLine, ok := runtime.Caller(1)
	var s string
	if ok {
		s = fmt.Sprintf("%s:%d", fileName, fileLine)
	} else {
		s = ""
	}
	return s
}
