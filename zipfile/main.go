package main

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"github.com/liuzl/cedar-go"
	"github.com/liuzl/store"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := filepath.Glob("./data/lemmatization-*.zip")
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.NewLevelStore("leveldb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	trie := cedar.New()

	var gzipOut bytes.Buffer
	dict := make(map[string]map[string][]string)
	for _, file := range files {
		r, err := zip.OpenReader(file)
		if err != nil {
			log.Fatal(err)
		}
		two := strings.Split(file, "-")
		if len(two) != 2 {
			log.Fatal(file + " name error")
		}
		t := strings.Split(two[1], ".")
		if len(t) != 2 {
			log.Fatal(file + "name error2")
		}
		lang := t[0]
		m := make(map[string][]string)
		_, has := dict[lang]
		if has {
			m = dict[lang]
		} else {
			dict[lang] = m
		}

		for _, f := range r.File {
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			body, err := ioutil.ReadAll(rc)
			if err != nil {
				log.Fatal(err)
			}
			rc.Close()
			var compressed bytes.Buffer
			w := gzip.NewWriter(&compressed)
			w.Write(body)
			w.Close()
			c := base64.StdEncoding.EncodeToString(compressed.Bytes())
			gzipOut.WriteString(f.Name + "\t" + c + "\n")

			for _, line := range strings.Split(string(body), "\n") {
				if strings.TrimSpace(line) == "" {
					continue
				}
				items := strings.Split(line, "\t")
				if len(items) != 2 {
					fmt.Println(items)
					continue
				}
				_, has := m[items[1]]
				if has {
					m[items[1]] = append(m[items[1]], items[0])
				} else {
					m[items[1]] = []string{items[0]}
				}
				key := items[1] + "\t" + lang + "\t" + items[0]
				err = db.Put(key, nil)
				if err != nil {
					log.Fatal(err)
				}
				trie.Insert([]byte(items[1]), 0)
			}
		}
		r.Close()
	}
	ioutil.WriteFile("gzip.out", gzipOut.Bytes(), os.FileMode(0664))
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(dict); err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("gob.out", buffer.Bytes(), os.FileMode(0664))
	trie.SaveToFile("cedar.gob", "gob")
	trie.SaveToFile("cedar.json", "json")
}
