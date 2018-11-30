package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/araddon/dateparse"
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

		t, err := dateparse.ParseAny(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(t)
	}
}
