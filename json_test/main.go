package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"n":"13702032331","msg":"hello"}`
	var item map[string]string
	err := json.Unmarshal([]byte(jsonStr), &item)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item)
	}
}
