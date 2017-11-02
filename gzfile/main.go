package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func gzFile() {
	filename := flag.String("file", "./file.txt.gz", "file to read")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	gr, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(gr)
	i := 0
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		i++
		fmt.Println(string(line))
		item := new(interface{})
		err = json.Unmarshal(line, item)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	gzFile()
}
