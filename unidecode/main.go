package main

import (
	"fmt"
	"github.com/liuzl/unidecode"
)

func main() {
	fmt.Println("vim-go")
	str := `乾隆爷的乾儿子是谁？`
	fmt.Println(unidecode.Unidecode(str))
}
