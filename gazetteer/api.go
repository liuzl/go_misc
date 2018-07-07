package main

import (
	"github.com/liuzl/store"
)

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

func (d *Dictionary) Insert(k string, v Values) error {
	return nil
}
