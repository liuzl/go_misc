package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	i := new(big.Int)
	str := "1000000000000000000000000000000000"
	n, err := fmt.Sscan(str, i)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return
	}
	fmt.Println(n, len(str), i)
	fmt.Println(math.Pow(2.2, 7.8))
}
