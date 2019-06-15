package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/schollz/closestmatch"
)

var (
	dict  = flag.String("dict", "words.txt", "dict file")
	input = flag.String("i", "", "file of original text to read")
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	flag.Parse()

	content, err := ioutil.ReadFile(*dict)
	if err != nil {
		glog.Fatal(err)
	}
	words := strings.Split(string(content), "\n")
	bagSizes := []int{2, 3, 4}
	cm := closestmatch.New(words, bagSizes)
	fmt.Println("ready!")

	var in *os.File
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			glog.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			glog.Fatal(c)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Println(line, cm.Closest(line))
		fmt.Println()
	}
}
