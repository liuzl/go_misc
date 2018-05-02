package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := new(big.Int)
	str := "10000000000000000000000000000000001"
	str = "19"
	n, err := fmt.Sscan(str, i)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return
	}
	fmt.Println(n, len(str), i)
	fmt.Println(i.ProbablyPrime(10))
}
