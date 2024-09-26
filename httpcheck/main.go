package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func isImageURL(url string) (bool, error) {
	// 发起 HTTP HEAD 请求，只获取响应头部
	resp, err := http.Head(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 检查响应状态码是否为 200 OK
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// 获取响应头中的 Content-Type
	contentType := resp.Header.Get("Content-Type")

	// 检查 Content-Type 是否为图片类型
	if strings.HasPrefix(contentType, "image/") {
		return true, nil
	}

	return false, fmt.Errorf("URL does not point to an image, Content-Type: %s", contentType)
}

func main() {
	// 检查是否提供了命令行参数
	if len(os.Args) < 2 {
		fmt.Println("请提供一个URL作为命令行参数")
		os.Exit(1)
	}

	// 获取命令行参数中的URL
	url := os.Args[1]

	isImage, err := isImageURL(url)
	if err != nil {
		fmt.Println("错误:", err)
	} else if isImage {
		fmt.Println("URL 是一个可访问的图片。")
	} else {
		fmt.Println("URL 不是图片。")
	}
}
