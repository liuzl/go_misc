package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/carlmjohnson/requests"
)

func main() {
	url := "https://profile.line-scdn.net/0hAzoguxL4HlVSHg3Kc2FhAm5bEDglMBgdKitSOiUaEjJ3JlwFPitZMHMWQWwqKw5XZ34GM3MfRGB2"
	fmt.Println(url)

	var buf bytes.Buffer
	var contentType string
	err := requests.URL(url).ToBytesBuffer(&buf).
		AddValidator(func(r *http.Response) error {
			contentType = r.Header.Get("Content-Type")
			return nil
		}).
		Fetch(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取图片类型
	imageType := contentType[6:] // 去掉 "image/" 前缀

	// 转换为 base64 编码
	base64Data := base64.StdEncoding.EncodeToString(buf.Bytes())

	// 生成内联图片数据
	inlineImageData := fmt.Sprintf("data:%s;base64,%s", contentType, base64Data)

	fmt.Printf("Image type: %s\n", imageType)
	fmt.Printf("Inline image data length: %d\n", len(inlineImageData))
	fmt.Printf("Inline image data (first 100 characters): %s...\n", inlineImageData[:100])
}
