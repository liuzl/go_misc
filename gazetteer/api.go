package main

import (
	"os"
	"path/filepath"

	"github.com/liuzl/dict"
	"github.com/liuzl/store"
)

func NewDictionary(dir string) (*Dictionary, error) {
	kvDir := filepath.Join(dir, "kv")
	cedarDir := filepath.Join(dir, "cedar")
	kv, err := store.NewLevelStore(kvDir)
	if err != nil {
		return nil, err
	}
	cedar := dict.New()
	if _, err = os.Stat(cedarDir); err == nil {
		err = cedar.LoadFromFile(cedarDir, "gob")
		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}
	d := &Dictionary{kv, cedar}
	return d, nil
}
