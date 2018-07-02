package main

import (
	"fmt"

	"github.com/liuzl/ling"
)

var nlp = ling.MustNLP(ling.Norm)

func main() {
	text := "北京ทันทุกเหตุการ有限公司"
	d := ling.NewDocument(text)
	if err := nlp.Annotate(d); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
	for i, token := range d.Tokens {
		fmt.Println(i, token)
	}
}
