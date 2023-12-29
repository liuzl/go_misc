package main

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

func main() {
	var responseString string
	err := requests.URL("https://www.baidu.com").ToString(&responseString).Fetch(context.Background())
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
	fmt.Printf(responseString)

	data := map[string]string{
		"a": "b",
	}
	err = requests.URL("http://127.0.0.1:9001/hdd/async").
		BodyJSON(&data).ToString(&responseString).
		Fetch(context.Background())
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
	fmt.Printf(responseString)

}
