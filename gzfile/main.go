package main

import (
	"./pb"
	"bufio"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/liuzl/store"
	"github.com/ttacon/libphonenumber"
	"io"
	"os"
	"strings"
	"time"
)

const (
	kTimeFormat = "02/Jan/2006:15:04:05 -0700"
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
	Ts          string    `json:"ts"`
}

func main() {
	defer glog.Flush()

	filename := flag.String("file", "./file.txt.gz", "file to read")
	storename := flag.String("dir", "./data", "leveldb dir")
	flag.Parse()

	db, err := store.NewLevelStore(*storename)
	if err != nil {
		glog.Fatal(err)
	}
	defer db.Close()

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
		//glog.Info(string(line))
		cc := strings.ToUpper(item.CountryCode)
		ab := &pb.AddressBook{Imei: item.IMEI, Cc: cc}
		number, err := libphonenumber.Parse(item.Phone, cc)
		if err != nil || !libphonenumber.IsValidNumber(number) {
			ab.Number = item.Phone
		} else {
			ab.Number = libphonenumber.Format(number, libphonenumber.E164)
		}
		t, err := time.Parse(kTimeFormat, item.Ts)
		if err != nil {
			glog.Error(err)
		} else {
			ab.Ts = t.Unix()
		}
		if item.Loc != "" {
			loc := &pb.Location{}
			fmt.Sscanf(item.Loc, "%f_%f", &loc.Latitude, &loc.Longitude)
			ab.Loc = loc
			glog.Info(ab)
		}
		for _, contact := range item.Contacts {
			person := &pb.Person{Name: contact.Name}
			number, err := libphonenumber.Parse(contact.Phone, cc)
			if err != nil || !libphonenumber.IsValidNumber(number) {
				person.Number = contact.Phone
			} else {
				person.Number = libphonenumber.Format(number, libphonenumber.E164)
			}
			ab.People = append(ab.People, person)
		}
		if has, _ := db.Has(ab.Imei); has {
			glog.Info("dup for ", ab.Imei)
			continue
		}
		out, err := proto.Marshal(ab)
		if err != nil {
			glog.Fatal(err)
		}
		err = db.Put(ab.Imei, out)
		if err != nil {
			glog.Fatal(err)
		}
	}
}
