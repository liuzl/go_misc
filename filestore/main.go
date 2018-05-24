package main

import (
	"fmt"
	"log"
	"zliu.org/filestore"
)

func main() {
	fmt.Println("vim-go")
	f, err := filestore.NewFileStore("store")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		f.WriteLine([]byte(fmt.Sprintf("hello %d", i)))
	}
}
