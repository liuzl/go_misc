package main

import (
	"fmt"

	"github.com/araddon/dateparse"
)

func main() {
	fmt.Println("vim-go")
	t, err := dateparse.ParseAny("6/10/2021, 8:26:03 AM")
	fmt.Printf("%+v, %+v\n", t, err)
}
