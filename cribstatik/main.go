package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	file = flag.String("file", "file.txt", "input file")
)

func FprintZipData(dest *bytes.Buffer, zipData []byte) {
	for _, b := range zipData {
		if b == '\n' {
			dest.WriteString(`\n`)
			continue
		}
		if b == '\\' {
			dest.WriteString(`\\`)
			continue
		}
		if b == '"' {
			dest.WriteString(`\"`)
			continue
		}
		if (b >= 32 && b <= 126) || b == '\t' {
			dest.WriteByte(b)
			continue
		}
		fmt.Fprintf(dest, "\\x%02x", b)
	}
}

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		return
	}
	var qb bytes.Buffer
	FprintZipData(&qb, data)
	fmt.Println(qb.String())
}
