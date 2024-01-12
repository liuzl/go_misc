package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Category    string  `json:"category"`
	Freshness   int     `json:"freshness"`
	Quantity    int     `json:"quantity"`
	Value       float32 `json:"value"`
	Recyclable  bool    `json:"recyclable"`
	Description string  `json:"description"`
}

func main() {
	r := Response{
		Category:    "paper",
		Freshness:   2,
		Quantity:    1,
		Recyclable:  true,
		Description: "A4 white paper",
	}

	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	var item Response
	json.Unmarshal([]byte("{}"), &item)
	fmt.Println(item)
}
