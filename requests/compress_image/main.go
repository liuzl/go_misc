package main

import (
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		fmt.Println("请提供输入图片文件名")
		os.Exit(1)
	}

	// 获取输入文件名
	inputFileName := os.Args[1]

	// 生成输出文件名
	ext := filepath.Ext(inputFileName)
	outputFileName := strings.TrimSuffix(inputFileName, ext) + "output" + ext

	// 打开原始图片
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// 获取原始图片的尺寸
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	// 获取原始文件大小
	originalFileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	originalSize := originalFileInfo.Size()

	// 调整图片大小，如果宽度或高度超过400像素，同时保持宽高比
	var resized image.Image
	if originalWidth > 400 || originalHeight > 400 {
		if originalWidth > originalHeight {
			resized = resize.Resize(400, 0, img, resize.Lanczos3)
		} else {
			resized = resize.Resize(0, 400, img, resize.Lanczos3)
		}
	} else {
		resized = img
	}

	// 创建输出文件
	out, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// 将调整大小后的图片编码为 JPEG 并保存
	jpeg.Encode(out, resized, &jpeg.Options{Quality: 85})

	// 获取转换后的图片尺寸
	newWidth := resized.Bounds().Dx()
	newHeight := resized.Bounds().Dy()

	// 获取转换后的文件大小
	newFileInfo, err := out.Stat()
	if err != nil {
		log.Fatal(err)
	}
	newSize := newFileInfo.Size()

	// 打印原始和转换后的图片信息
	fmt.Printf("原始图片: %s\n", inputFileName)
	fmt.Printf("宽度 %d, 高度 %d, 文件大小 %d 字节\n", originalWidth, originalHeight, originalSize)
	fmt.Printf("转换后图片: %s\n", outputFileName)
	fmt.Printf("宽度 %d, 高度 %d, 文件大小 %d 字节\n", newWidth, newHeight, newSize)
}
