package main

import (
	"flag"
	"fmt"
	"net/url"
)

var (
	u = flag.String("url", "https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=7&rsv_sug1=6", "url")
)

func main() {
	flag.Parse()
	fmt.Println("vim-go")
	fmt.Println(*u)

	values, err := url.ParseRequestURI(*u)
	fmt.Println(values, err)
}
