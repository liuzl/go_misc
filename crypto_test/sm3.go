package main

import (
	"fmt"

	"github.com/tjfoc/gmsm/sm3"
)

func main() {
	data := "test"
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	fmt.Printf("digest value is: %x\n", sum)
}
