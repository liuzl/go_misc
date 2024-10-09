package main

import (
	"fmt"

	"github.com/liuzl/store"
)

func main() {
	fmt.Println("vim-go")
	if db, err := store.NewLevelStore("./data/db"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", db)
	}
}
