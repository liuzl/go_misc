package main

import (
	"flag"
	"fmt"
)

var (
	input = flag.String("i", "default input, 你好", "input text")
)

func main() {
	flag.Parse()
	fmt.Printf("str: %+v, len:%d\n", *input, len(*input))
	fmt.Printf("rune: %+v, len:%d\n", []rune(*input), len([]rune(*input)))
	fmt.Printf("byte: %+v, len:%d\n", []byte(*input), len([]byte(*input)))
	/*
		str: 如果是汉字呢haha, len:22
		rune: [22914 26524 26159 27721 23383 21602 104 97 104 97], len:10
		byte: [229 166 130 230 158 156 230 152 175 230 177 137 229 173 151 229 145 162 104 97 104 97], len:22
	*/
	str := "哈哈this is good"
	fmt.Println(len(str[0:4]))

	for i, c := range str {
		fmt.Println(i, string(c), str[i:])
	}
}
