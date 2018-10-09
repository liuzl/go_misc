package main

import (
	"encoding/json"
	"fmt"

	"github.com/liuzl/goutil/rest"
)

func main() {
	jsonStr := `{"n":"13702032331","msg":"hello"}`
	jsonStr = `{"status":"ok","message":{"key":"20181009063703:4d819b10a0c780eeb2da8d6a2b2c921d","value":"eyJuIjogIjYyODE1MTkxMDM1MTYiLCAibXNnIjogInRoaXMgaXMganVzdCBhIHRlc3QifQ=="}}`
	var item map[string]string
	var ret rest.RestMessage
	err := json.Unmarshal([]byte(jsonStr), &ret)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(ret)
	}
	fmt.Println(ret.Message)
	b, _ := json.Marshal(ret.Message)
	err = json.Unmarshal(b, &item)
	fmt.Println(err, item)
}
