package main

import (
	"fmt"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func main() {
	fmt.Println("vim-go")
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	text := "drop me a line in next wednesday at 2:25 p.m"
	text = "December 19th, 2018"
	r, err := w.Parse(text, time.Now())
	if err != nil {
		// an error has occurred
	}
	if r == nil {
		// no matches found
	}

	fmt.Println(
		"the time",
		r.Time.String(),
		"mentioned in",
		text[r.Index:r.Index+len(r.Text)],
	)

}
