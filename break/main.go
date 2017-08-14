package main

import (
	"fmt"
	"time"
)

func main() {
DONE:
	for {
		fmt.Println("for")
		select {
		case <-time.After(5 * time.Second):
			//goto DONE
			break DONE
		}
	}
	fmt.Println("DONE")
}
