package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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

	// 获取文件大小
	fileSize := buf.Len()
	fmt.Printf("Original file size: %d bytes\n", fileSize)
	fmt.Printf("Content-Type: %s\n", contentType)

	// 尝试获取图片格式
	_, imgFormat, err := image.DecodeConfig(bytes.NewReader(buf.Bytes()))
	if err != nil {
		fmt.Printf("Warning: Unable to determine image format: %v\n", err)
	} else {
		fmt.Printf("Image format: %s\n", imgFormat)
	}

	// 获取图片尺寸
	img, _, err := image.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	fmt.Printf("Original image dimensions: %dx%d pixels\n", width, height)

	// 获取图片类型
	imageType := contentType[6:] // 去掉 "image/" 前缀

	// 转换为 base64 编码
	base64Data := base64.StdEncoding.EncodeToString(buf.Bytes())

	// 生成内联图片数据
	inlineImageData := fmt.Sprintf("data:%s;base64,%s", contentType, base64Data)

	fmt.Printf("Image type: %s\n", imageType)

	return inlineImageData, nil
}
