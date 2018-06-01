package main

import (
	//"gopkg.in/cheggaaa/pb.v1"
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/liuzl/goutil"
	"time"
)

func main() {
	fmt.Println(goutil.FileLineCount("main.go"))
	count := 100000
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
}
