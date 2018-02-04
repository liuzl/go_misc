package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/liuzl/phonenumbers"
	"io/ioutil"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(phonenumbers.CarriersPb)
	data, err := base64.StdEncoding.DecodeString(phonenumbers.CarriersPb)
	fmt.Println(data, err)
	reader, err := gzip.NewReader(bytes.NewReader(data))
	fmt.Println(reader, err)
	pbBytes, err := ioutil.ReadAll(reader)
	fmt.Println(pbBytes, err)
}
