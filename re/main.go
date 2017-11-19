package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`^[\d\+\.]+$`)
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadBytes('\n')
		if c == io.EOF {
			break
		}
		str := strings.TrimSpace(string(line))
		fmt.Println(re.MatchString(str), str)
	}
}
