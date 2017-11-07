package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/liuzl/store"
)

var (
	dbfile = flag.String("dir", "./data", "leveldb dir")
)

func main() {
	defer glog.Flush()
	flag.Parse()
	db, err := store.NewLevelStore(*dbfile)
	if err != nil {
		glog.Fatal(err)
	}
	defer db.Close()
	cnt := 0
	err = db.ForEach(nil, func(key, value []byte) (bool, error) {
		//glog.Info(string(key))
		cnt += 1
		return true, nil
	})
	glog.Info(cnt)
}
