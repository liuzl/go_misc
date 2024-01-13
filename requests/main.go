package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/carlmjohnson/requests"
)

func Record(rt http.RoundTripper, rawReq *string, rawRes *string) requests.Transport {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return requests.RoundTripFunc(func(req *http.Request) (res *http.Response, err error) {
		defer func() {
			if err != nil {
				err = fmt.Errorf("problem while recording transport: %w", err)
			}
		}()

		b, err := httputil.DumpRequest(req, true)
		if err != nil {
			return nil, err
		}
		*rawReq = string(b)
		if res, err = rt.RoundTrip(req); err != nil {
			return
		}
		b, err = httputil.DumpResponse(res, true)
		if err != nil {
			return nil, err
		}
		*rawRes = string(b)
		return
	})
}

func main() {
	var responseString string
	var req string
	var res string
	err := requests.URL("https://www.baidu.com").Transport(Record(nil, &req, &res)).
		ToString(&responseString).Fetch(context.Background())
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
	//fmt.Println(responseString)

	fmt.Println("Request:\n" + req)
	fmt.Println("Response:\n" + res)

}
