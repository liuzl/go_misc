package main

import (
	"fmt"
	"os"

	"github.com/liuzl/ling"
)

var nlp = ling.MustNLP(ling.Norm)

func cut(line string) []string {
	d := ling.NewDocument(line)
	if err := nlp.Annotate(d); err != nil {
		return nil
	}
	return d.XTokens(ling.Norm)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <str>\n", os.Args[0])
		return
	}
	ret := cut(os.Args[1])
	for _, r := range ret {
		fmt.Printf("%d,[%s]\n", len(r), r)
	}
}
