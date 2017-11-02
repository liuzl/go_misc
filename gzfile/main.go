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

type Contact struct {
	Phone string `json:"b2"`
	Name  string `json:"b1"`
}

type Item struct {
	CountryCode string    `json:"a3"`
	Contacts    []Contact `json:"b"`
}

func main() {
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
	var item Item
	for {
		line, c := br.ReadBytes('\n')
		if c == io.EOF {
			break
		}
		i++
		err = json.Unmarshal(line, &item)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(item)
		fmt.Println("\n\n\n")
	}
}
