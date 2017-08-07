package main

import (
	"flag"
	"fmt"
	"github.com/huichen/sego"
)

var (
	str = flag.String("str", "百度在线网络技术（北京）有限公司是一家好公司", "input str")
)

func main() {
	flag.Parse()
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("dictionary.txt")

	text := []byte(*str)
	segments := segmenter.Segment(text)
	fmt.Println(sego.SegmentsToString(segments, false))
}
