package main

import (
	"fmt"
	"github.com/mozillazg/go-unidecode"
)

func main() {
	s := "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc

	s = "北京"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing

	s = "kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: kozuscek
}
