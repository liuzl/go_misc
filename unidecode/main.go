package main

import (
	"fmt"
	"github.com/liuzl/unidecode"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(unidecode.Unidecode(`乾隆爷的乾儿子是谁？`))
	fmt.Println(unidecode.Unidecode("multiply by?"))
}
