package main

import (
	"fmt"
	"os"

	"github.com/liuzl/motto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <jsfile>\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println("running", os.Args[1])
	_, v, err := motto.Run(os.Args[1])
	fmt.Println(v, err)
}
