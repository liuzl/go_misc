package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		data, err := base64.StdEncoding.DecodeString(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
	}
}
