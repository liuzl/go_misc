package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func FileGuard(f string) bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return false
	}
	fname := filepath.Join(dir, f)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		if err = ioutil.WriteFile(fname,
			[]byte(time.Now().UTC().Format(time.RFC3339)), 0644); err != nil {
			return false
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(FileGuard("first.lock"))
}
