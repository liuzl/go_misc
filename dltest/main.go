package main

import (
	"fmt"
	"github.com/liuzl/dl"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFile writes data to file at filePath
func WriteFile(filePath, fileName string, data []byte) {
	path := filepath.Join(filePath, fileName)
	fmt.Println(path)
	err := ioutil.WriteFile(path, data, os.FileMode(0664))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	resp := dl.DownloadUrl("https://instagram.fcgk6-1.fna.fbcdn.net/vp/6af9dfe77a72a523f5990c74a8c537fb/5AC4DA8F/t50.2886-16/29911547_158978848119097_7202585723608760320_n.mp4")
	if resp.Error != nil {
		fmt.Printf("error fetching: %+v\n", resp.Error)
		return
	}
	WriteFile(".", "a.mp4", resp.Content)
}
