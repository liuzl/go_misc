package main

import (
	"bufio"
	"fmt"
	"github.com/liuzl/segment"
	"io"
	"os"
	"strings"
)

// Word Types
const (
	None = iota
	Number
	Letter
	Kana
	Ideo
)

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadBytes('\n')
		if c == io.EOF {
			break
		}
		fmt.Println(string(line[:len(line)-1]))
		seg1 := segment.NewSegmenterDirect(line)
		for seg1.Segment() {
			text := seg1.Text()
			if strings.TrimSpace(text) == "" {
				continue
			}
			label := seg1.Type()
			fmt.Print(text, "\\", label, " ")
		}
		fmt.Println("\n")
	}
	fmt.Println(None, Number, Letter, Kana, Ideo)
	fmt.Println("None, Number, Letter, Kana, Ideo")
}
