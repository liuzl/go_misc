package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"github.com/ttacon/libphonenumber"
	"io"
	"os"
	"strings"
)

type Contact struct {
	Phone string `json:"b2"`
	Name  string `json:"b1"`
}

type Item struct {
	IMEI        string    `json:"a1"`
	Phone       string    `json:"a2"`
	Loc         string    `json:"a5"`
	CountryCode string    `json:"a3"`
	Contacts    []Contact `json:"b"`
}

func main() {
	defer glog.Flush()
	filename := flag.String("file", "./file.txt.gz", "file to read")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()
	gr, err := gzip.NewReader(file)
	if err != nil {
		glog.Fatal(err)
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
			glog.Error(err)
			continue
		}
		cc := strings.ToUpper(item.CountryCode)
		glog.Info("loc:", item.Loc, "\tIMEI:", item.IMEI, "\tPhone:", item.Phone, "\tcc:", cc)
		for _, contact := range item.Contacts {
			number, err := libphonenumber.Parse(contact.Phone, cc)
			if err != nil {
				//		glog.Error(contact.Phone, err)
				continue
			}

			if !libphonenumber.IsValidNumber(number) {
				//		glog.Error(contact.Phone, "|", cc, "|", contact.Name, " invalid number")
				continue
			}

			//	glog.Info(contact.Phone, "|", cc, "|", contact.Name, " ", number)
		}
	}
}
