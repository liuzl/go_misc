package main

import (
	"fmt"
	"time"
)

func main() {
	timeoutchan := make(chan bool)

	go func() {
		<-time.After(2 * time.Second)
		timeoutchan <- true
	}()

	select {
	case <-timeoutchan:
		fmt.Println("timeoutChan")
		break
	case <-time.After(122 * time.Second):
		fmt.Println("time.After 10")
		break
	}

	fmt.Println("Hello, playground")
}
