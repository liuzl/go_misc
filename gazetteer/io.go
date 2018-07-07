package main

import (
	"encoding/json"
	"os"

	"github.com/liuzl/store"
)

func (d *Dictionary) Dump(path string) error {
	wf, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer wf.Close()
	err = d.kv.ForEach(nil, func(key, value []byte) (bool, error) {
		var v Values
		if err := store.BytesToObject(value, &v); err != nil {
			return false, err
		}
		m := map[string]interface{}{"key": string(key), "values": v}
		b, err := json.Marshal(m)
		if err != nil {
			return false, err
		}
		if _, err = wf.Write(append(b, '\n')); err != nil {
			return false, err
		}
		return true, nil
	})
	if err != nil {
		return err
	}
	return nil
}
