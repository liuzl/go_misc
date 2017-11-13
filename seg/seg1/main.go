package main

import (
	"bufio"
	"fmt"
	"github.com/blevesearch/segment"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(segment.SplitWords)
	for scanner.Scan() {
		tokenBytes := scanner.Bytes()
		word := strings.TrimSpace(string(tokenBytes))
		if word == "" {
			continue
		}
		fmt.Println(word)
	}
}
