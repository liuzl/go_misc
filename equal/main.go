package main

import (
	"fmt"

	"github.com/liuzl/fmr"
)

type TS struct {
	*fmr.TableState
}

func (t *TS) e(ts *TS) bool {
	if t == nil && ts == nil {
		return true
	}
	if (t != nil && ts == nil) || (t == nil && ts != nil) {
		return false
	}
	return false
}

func main() {
	var t *TS
	ret := t.e(t)
	fmt.Println(ret)

	var a, b interface{}
	fmt.Println(a == b)
}
