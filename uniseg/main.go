package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/golang/glog"
	"github.com/rivo/uniseg"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			glog.Fatal(c)
		}
		fmt.Println(line)
		gr := uniseg.NewGraphemes(line)
		for gr.Next() {
			fmt.Println(gr.Str())
		}
	}
}
