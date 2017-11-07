package main

import (
	"./pb"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/liuzl/filestore"
	"github.com/liuzl/store"
	"github.com/satori/go.uuid"
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

func MD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	defer glog.Flush()

	filename := flag.String("file", "./file.txt.gz", "file to read")
	storename := flag.String("dir", "./data", "leveldb dir")
	contactsfile := flag.String("contacts", "./contacts", "contacts dir")
	flag.Parse()

	fss := make(map[string]*filestore.FileStore)

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
	m := make(map[string]int)
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
		if len(item.Contacts) == 0 {
			continue
		}
		if item.IMEI == "" {
			glog.Error("IMEI is null")
			item.IMEI = uuid.NewV4().String()
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
			//glog.Info(ab)
		}
		set := make(map[string]bool)
		for _, contact := range item.Contacts {
			person := &pb.Person{Name: contact.Name}
			number, err := libphonenumber.Parse(contact.Phone, cc)
			if err != nil || !libphonenumber.IsValidNumber(number) {
				person.Number = contact.Phone
			} else {
				person.Number = libphonenumber.Format(number, libphonenumber.E164)
			}

			// dedup
			if _, has := set[person.Number+"\t"+person.Name]; has {
				continue
			} else {
				set[person.Number+"\t"+person.Name] = true
			}

			ab.People = append(ab.People, person)
		}

		// deal with duplications
		save := false
		l := len(ab.People)
		if cnt, has := m[ab.Imei]; has {
			if cnt < l {
				save = true
				m[ab.Imei] = l
			}
		} else {
			if has, _ = db.Has(ab.Imei); has {
				// we have this imei twice at least
				oldAb := &pb.AddressBook{}
				in, _ := db.Get(ab.Imei)
				if err = proto.Unmarshal(in, oldAb); err != nil {
					glog.Error(err)
					save = true
				} else {
					m[ab.Imei] = len(oldAb.People)
					if l > len(oldAb.People) {
						save = true
						m[ab.Imei] = l
					}
				}
			} else {
				save = true
			}
		}

		if save {
			out, err := proto.Marshal(ab)
			if err != nil {
				glog.Fatal(err)
			}
			err = db.Put(ab.Imei, out)
			if err != nil {
				glog.Fatal(err)
			}
			for _, person := range ab.People {
				key := MD5(person.Number)
				fn := fmt.Sprintf("%s/%s", *contactsfile, key[:2])
				fs, has := fss[fn]
				if !has {
					fs, err = filestore.NewFileStore(fn)
					if err != nil {
						glog.Fatal(err)
					}
					fss[fn] = fs
					defer fs.Close()
				}

				name := strings.Replace(person.Name, "\t", " ", -1)
				name = strings.Replace(name, "\n", " ", -1)
				one := fmt.Sprintf("%s\t%s\t%s\t%s", person.Number, cc, ab.Imei, name)
				fs.WriteLine([]byte(one))
			}
		}
	}
}
