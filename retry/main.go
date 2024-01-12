package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/avast/retry-go/v4"
)

func main() {
	data, err := retry.DoWithData(
		func() (string, error) {
			if rand.Intn(100) < 50 {
				return "Hello World!", nil
			} else {
				return "", errors.New("ERROR")
			}
		},
		retry.Attempts(5),
		retry.Delay(1*time.Second),
	)

	fmt.Println(data)
	fmt.Println(err)
}
