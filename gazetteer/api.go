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

func (d *Dictionary) Get(key string) (Values, error) {
	b, err := d.kv.Get(key)
	if err != nil {
		return nil, err
	}
	var v Values
	if err = store.BytesToObject(b, &v); err != nil {
		return nil, err
	}
	return v, nil
}

func (d *Dictionary) PrefixMatch(text string) (map[string]Values, error) {
	ret := make(map[string]Values)
	for _, id := range d.cedar.PrefixMatch([]byte(text), 0) {
		key, err := d.cedar.Key(id)
		if err != nil {
			return nil, err
		}
		word := string(key)
		v, err := d.Get(word)
		if err != nil {
			return nil, err
		}
		ret[word] = v
	}
	return ret, nil
}
