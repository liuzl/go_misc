package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

func main() {
	var num int64 = 6285777273663

	// convert int64 to []byte
	buf := make([]byte, binary.MaxVarintLen64)
	_ = binary.PutVarint(buf, num)
	// convert []byte to int64
	x, n := binary.Varint(buf)
	fmt.Printf("x is: %v, %s, n is: %v\n", x, base64.StdEncoding.EncodeToString(buf), n)
}
