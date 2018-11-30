package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/araddon/dateparse"
	"github.com/liuzl/ling"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			log.Fatal(c)
		}
		line = strings.TrimSpace(line)
		d := ling.NewDocument(line)
		if err := ling.MustNLP().Annotate(d); err != nil {
			log.Fatal(err)
		}

		for _, token := range d.Tokens {
			if token.Type == ling.Space {
				continue
			}
			fmt.Println(line[token.StartByte:])
			t, err := dateparse.ParseAny(line[token.StartByte:])
			if err != nil {
				continue
			} else {
				fmt.Println(t)
				break
			}
		}
	}
}
