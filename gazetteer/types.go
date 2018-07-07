package main

import (
	"github.com/liuzl/dict"
	"github.com/liuzl/store"
)

type DictValue struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Dictionary struct {
	kv    *store.LevelStore
	cedar *dict.Cedar
}
