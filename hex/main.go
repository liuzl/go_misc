package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)
	s := "48656c6c6f20476f7068657221"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", decoded)

	//src = []byte("48656c6c6f20476f7068657221")
	src = []byte("48656c6c6f20576f726c64")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])
}
