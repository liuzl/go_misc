package main

import (
	"errors"
	"fmt"
	"github.com/crawlerclub/x/downloader"
	"github.com/crawlerclub/x/parser"
	"github.com/crawlerclub/x/types"
	"io/ioutil"
	"log"
)

func DouyinDL(url string) error {
	//url := "https://www.douyin.com/share/video/6418816467571051778"
	req := &types.HttpRequest{Url: url, Method: "GET", UseProxy: false, Platform: "pc"}
	res := downloader.Download(req)
	if res.Error != nil {
		return res.Error
	}
	items, err := parser.ParseRegex(res.Text, "video_id=(.+?)\\\\u")
	if err != nil {
		return err
	}
	if len(items) <= 0 {
		return errors.New("parse video_id error")
	}
	mp4Url := fmt.Sprintf("http://aweme.snssdk.com/aweme/v1/play/?video_id=%s", items[0])
	log.Println(mp4Url)
	mp4Req := &types.HttpRequest{Url: mp4Url, Method: "GET", UseProxy: false, Platform: "pc"}
	ret := downloader.Download(mp4Req)
	if ret.Error != nil {
		return ret.Error
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s.mp4", items[0]), ret.Content, 0666)
	return err
}

func main() {
	url := "https://www.douyin.com/share/video/6418816467571051778"
	DouyinDL(url)
}
