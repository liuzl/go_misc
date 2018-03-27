package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := new(big.Int)
	_, err := fmt.Sscan("1000000000000000000000000000000000", i)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return
	}
	fmt.Println(i)
}
