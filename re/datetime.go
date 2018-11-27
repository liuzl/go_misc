package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	MonthStr = `(?:(?:jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)[a-z]*)`
	ReDate   = regexp.MustCompile(`(?is)((?:` + MonthStr + `[\.,\-\s]*\d{1,2}(?:st|nd|rd|th)*[\.,\-\s]*(\d{4}))|` +
		`(?:\d{1,2}(?:st|nd|rd|th)*[\.,\-\s]*` + MonthStr + `[\.,\-\s]*(\d{4}))|` +
		MonthStr + `.\d{1,2}|` +
		`(?:(19|20)\d{2}[^0-9]\d{1,2}[^0-9]\d{1,2})|` +
		`(?:\d{1,2}[^0-9]\d{1,2}[^0-9](19|20)\d{2})|` +
		`(?:(\d{4}年){0,1}\d{1,2}月\d{1,2}日))`)

	ReTime = regexp.MustCompile(`(?is)((?:0?|[12])\d\s*:+\s*[0-5]\d(?:\s*:+\s*[0-5]\d)?(?:\s*[,:.]*\s*(?:am|pm))?|` +
		`(?:0?|[12])\d\s*[.\s]+\s*[0-5]\d(?:\s*[,:.]*\s*(?:am|pm))+)`)
)

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		ret := ReDate.FindAllString(line, -1)
		for _, item := range ret {
			fmt.Println(item)
		}
	}
}
