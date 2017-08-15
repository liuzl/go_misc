package main

import (
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
)

var (
	imgFile = flag.String("img", "test.jpg", "image file")
)

func main() {
	// open "test.jpg"
	flag.Parse()
	file, err := os.Open(*imgFile)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(150, 0, img, resize.Lanczos3)

	out, err := os.Create(fmt.Sprintf("%s_resize.jpg", *imgFile))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
