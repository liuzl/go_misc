package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func main() {
	//re := regexp.MustCompile(`^[\d\+\.]+$`)
	re := regexp.MustCompile(`\p{Han}+`)
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadBytes('\n')
		if c == io.EOF {
			break
		}
		str := strings.TrimSpace(string(line))
		fmt.Println(re.FindAllString(str, -1), str)
	}

	var arrs [][]string
	fmt.Println(reflect.TypeOf(arrs))
	//fmt.Println(arrs...)
}
