package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"github.com/golang/glog"
	"io/ioutil"
)

var (
	file = flag.String("file", "input.txt", "the input file")
	name = flag.String("name", "str", "the var name to store the string")
)

func EncodeFileToString(fname string) (string, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", err
	}
	var compressed bytes.Buffer
	w := gzip.NewWriter(&compressed)
	w.Write(data)
	w.Close()
	return base64.StdEncoding.EncodeToString(compressed.Bytes()), nil
}

func DecodeString(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

func main() {
	flag.Parse()
	str, err := EncodeFileToString(*file)
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info(str)
	data, err := DecodeString(str)
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info(string(data))
}
