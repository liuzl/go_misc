package main

import (
	"fmt"

	es "github.com/elastic/go-elasticsearch/v7"
)

func main() {
	client, _ := es.NewDefaultClient()
	fmt.Println(es.Version)
	fmt.Println(client.Info())
}
