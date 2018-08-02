package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("cashcash.log")
	if err != nil {
		log.Fatal(err)
	}
	var m map[string]interface{}
	if err = json.Unmarshal(b, &m); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(strings.NewReader(m["request"].(string)))
	r, err := http.ReadRequest(reader)
	if err != nil {
		log.Fatal(err)
	}
	r.ParseForm()
	fmt.Println(r.FormValue("img"))
}
