package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/liuzl/dict"
	"github.com/liuzl/store"
)

var (
	dir = flag.String("dir", "data", "data dir")
)

func main() {
	flag.Parse()

	kvDir := filepath.Join(*dir, "kv")
	cedarDir := filepath.Join(*dir, "cedar")
	kv, err := store.NewLevelStore(kvDir)
	if err != nil {
		log.Fatal(err)
	}
	cedar := dict.New()
	if _, err = os.Stat(cedarDir); err == nil {
		err = cedar.LoadFromFile(cedarDir, "gob")
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println(kv)
}
