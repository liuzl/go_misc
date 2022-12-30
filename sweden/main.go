package main

import (
	personnummer "github.com/personnummer/go/v3"
)

func main() {
	print(personnummer.Valid("193408248981"))
	//=> true
}
