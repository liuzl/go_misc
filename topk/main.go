package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cheggaaa/pb"
	"github.com/dgryski/go-topk"
	"github.com/liuzl/goutil"
	"github.com/liuzl/ling"
)

var (
	k = flag.Int("n", 100, "k")
	i = flag.String("i", "input.txt.gz", "input file")
)

var nlp = ling.MustNLP(ling.Norm)

func main() {
	flag.Parse()
	count, err := goutil.FileLineCount(*i)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(*i)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(gr)
	tk := topk.New(*k)
	bar := pb.StartNew(count)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		d := ling.NewDocument(line)
		if err := nlp.Annotate(d); err != nil {
			log.Fatal(err)
		}
		tokens := d.XRealTokens(ling.Norm)
		for i := 1; i < len(tokens); i++ {
			//fmt.Println(strings.Join(tokens[i:], ""))
			tk.Insert(strings.Join(tokens[i:], ""), 1)
		}
		bar.Increment()
	}
	bar.FinishPrint("done!")
	fmt.Println(len(tk.Keys()))
	for _, v := range tk.Keys() {
		fmt.Println(v.Key, v.Count, v.Error)
	}
}
