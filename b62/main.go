package main

import (
	"fmt"

	"github.com/liuzl/goutil"
)

func main() {
	var i uint64 = 1736308120714183564
	fmt.Println(goutil.ToB62(i))
}
