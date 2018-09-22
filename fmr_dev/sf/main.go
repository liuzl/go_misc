package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/liuzl/fmr"
	"github.com/liuzl/goutil"
)

var (
	grammar = flag.String("g", "sf.grammar", "grammar file")
	input   = flag.String("i", "input.txt", "file of original text to read")
	debug   = flag.Bool("debug", false, "debug mode")
	start   = flag.String("start", "flight", "start rule")
)

func main() {
	flag.Parse()
	if *debug {
		fmr.Debug = true
	}
	g, err := fmr.GrammarFromFile(*grammar)
	if err != nil {
		glog.Fatal(err)
	}
	if *debug {
		b, err := goutil.JsonMarshalIndent(g, "", "    ")
		if err != nil {
			glog.Fatal(err)
		}
		fmt.Printf("%s\n", string(b))
	}

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
		fmt.Println(line)

		ps, err := g.EarleyParseAll(line, *start)
		if err != nil {
			glog.Fatal(err)
		}
		for i, p := range ps {
			for _, f := range p.GetFinalStates() {
				trees := p.GetTrees(f)
				//fmt.Printf("%+v\n", p)
				fmt.Printf("p%d tree number:%d\n", i, len(trees))
				for _, tree := range trees {
					tree.Print(os.Stdout)
					sem, err := tree.Semantic()
					fmt.Println(sem)
					if err != nil {
						glog.Fatal(err)
					}
					if *debug {
						fmt.Printf("%s = ?\n", sem)
					}
				}
			}
		}
		fmt.Println()
	}
}
