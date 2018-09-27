package main

import (
	"fmt"
	"regexp"
)

// see: https://en.wikipedia.org/wiki/Emoji for emoji ranges
var wordTokenRegex = regexp.MustCompile("((\\pL|\\pN|[\u20A0-\u20CF]|[\u2600-\u27BF])+|[\U0001F170-\U0001F9CF])")

// TokenizeString returns the words in the passed in string, split by non word characters including emojis
func TokenizeString(str string) []string {
	return wordTokenRegex.FindAllString(str, -1)
}

func main() {
	str := "新华社消息： 2007.9.18 Emerging after two hours of talks with Chinese President Xi Jinping, Trump said he doesn't fault China for taking advantage of differences between the way the two countries do business. 3 hours later"
	ret := TokenizeString(str)
	fmt.Printf("%+v\n", ret)
}
