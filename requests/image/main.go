package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/carlmjohnson/requests"
)

func main() {
	url := "https://profile.line-scdn.net/0hAzoguxL4HlVSHg3Kc2FhAm5bEDglMBgdKitSOiUaEjJ3JlwFPitZMHMWQWwqKw5XZ34GM3MfRGB2"
	fmt.Println("Downloading image from:", url)

	inlineImageData, err := getInlineImageData(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Inline image data length: %d\n", len(inlineImageData))
	fmt.Printf("Inline image data (first 100 characters): %s...\n", inlineImageData[:100])
}

func getInlineImageData(url string) (string, error) {
	var buf bytes.Buffer
	var contentType string

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := requests.URL(url).ToBytesBuffer(&buf).
		AddValidator(func(r *http.Response) error {
			contentType = r.Header.Get("Content-Type")
			if !strings.HasPrefix(contentType, "image/") {
				return fmt.Errorf("invalid content type: %s", contentType)
			}
			return nil
		}).Fetch(ctx)

	if err != nil {
		if err == context.DeadlineExceeded {
			return "", fmt.Errorf("timeout while downloading the image")
		}
		return "", fmt.Errorf("failed to download image: %w", err)
	}

	if buf.Len() == 0 {
		return "", fmt.Errorf("downloaded image is empty")
	}

	// 获取图片类型
	imageType := contentType[6:] // 去掉 "image/" 前缀

	// 转换为 base64 编码
	base64Data := base64.StdEncoding.EncodeToString(buf.Bytes())

	// 生成内联图片数据
	inlineImageData := fmt.Sprintf("data:%s;base64,%s", contentType, base64Data)

	fmt.Printf("Image type: %s\n", imageType)

	return inlineImageData, nil
}
