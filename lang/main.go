package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abadojack/whatlanggo"
)

var (
	text = flag.String("text", "this is English", "text to detect")
)

func main() {
	flag.Parse()
	//fmt.Println(*text)
	info := whatlanggo.DetectLang(*text)
	s := whatlanggo.LangToString(info)
	data, _ := json.Marshal(s)
	fmt.Println(string(data))
	//fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])
}
