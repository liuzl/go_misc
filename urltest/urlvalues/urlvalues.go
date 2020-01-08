package main

import (
	"fmt"
	"net/url"
)

func main() {
	// url encode
	v := url.Values{}
	v.Add("msg", "此订单不存在或已经提交")
	body := v.Encode()
	fmt.Println(v)
	fmt.Println(body)
	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println(m)
}
