package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/corona10/goimagehash"
)

func main() {
	file1, _ := os.Open("sample1.jpg")
	file2, _ := os.Open("sample2.jpg")
	defer file1.Close()
	defer file2.Close()

	img1, _ := jpeg.Decode(file1)
	img2, _ := jpeg.Decode(file2)
	hash1, _ := goimagehash.AverageHash(img1)
	hash2, _ := goimagehash.AverageHash(img2)
	distance, _ := hash1.Distance(hash2)
	fmt.Printf("Distance between images: %v\n", distance)
	fmt.Printf("hash1 bit size: %v\n", hash1.Bits())
	fmt.Printf("hash2 bit size: %v\n", hash2.Bits())

	fmt.Printf("hash1: %v\n", hash1)
	fmt.Printf("hash2: %v\n", hash2)

	fmt.Printf("hash1: %s\n", hash1.ToString())
	fmt.Printf("hash2: %s\n", hash2.ToString())

	hash1, _ = goimagehash.DifferenceHash(img1)
	hash2, _ = goimagehash.DifferenceHash(img2)
	distance, _ = hash1.Distance(hash2)
	fmt.Printf("Distance between images: %v\n", distance)
	width, height := 8, 8
	hash3, _ := goimagehash.ExtAverageHash(img1, width, height)
	hash4, _ := goimagehash.ExtAverageHash(img2, width, height)
	distance, _ = hash3.Distance(hash4)
	fmt.Printf("Distance between images: %v\n", distance)
	fmt.Printf("hash3 bit size: %v\n", hash3.Bits())
	fmt.Printf("hash4 bit size: %v\n", hash4.Bits())

	hash5, _ := goimagehash.PerceptionHash(img1)
	hash6, _ := goimagehash.PerceptionHash(img2)
	distance, _ = hash5.Distance(hash6)
	fmt.Printf("Distance between images: %v\n", distance)
	fmt.Printf("hash5 bit size: %v\n", hash5.Bits())
	fmt.Printf("hash6 bit size: %v\n", hash6.Bits())

	fmt.Printf("hash5: %v\n", hash5.ToString())
	fmt.Printf("hash6: %v\n", hash6.ToString())
}
