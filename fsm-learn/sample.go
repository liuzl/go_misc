package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	f := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(f.Current())
	fmt.Println(fsm.Visualize(f))
}
