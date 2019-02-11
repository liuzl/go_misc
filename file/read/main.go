package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		return
	}
	in, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	br := bufio.NewReader(in)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			line = strings.TrimSuffix(line, "\n")
			if line != "" {
				fmt.Println(line)
			}
			break
		}
		if c != nil {
			log.Fatal(c)
		}
		fmt.Println(strings.TrimSuffix(line, "\n"))
	}
}
