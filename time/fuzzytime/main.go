package main

import (
	"fmt"

	"github.com/bcampbell/fuzzytime"
)

func main() {
	inputs := []string{
		"Wed Apr 16 17:32:51 NZST 2014",
		"2010-02-01T13:14:43Z", // an iso 8601 form
		"no date or time info here",
		"Published on March 10th, 1999 by Brian Credability",
		"2:51pm",
		"6/10/2021, 8:26:03 AM",
	}

	for _, inp := range inputs {
		dt, _, _ := fuzzytime.Extract(inp)
		fmt.Println(dt.ISOFormat())
	}
}
