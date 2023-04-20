package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	container := launcher.MustResolveURL("127.0.0.1:9222")
	// Launch a browser with default options
	browser := rod.New().ControlURL(container).MustConnect().MustPage("https://www.bilibili.com/read/home")

	fmt.Println(browser.MustEval("()=>document.title").String())

	client := browser.MustWaitLoad().MustElement(".article-item")

	result := client.MustEval(`() => Array.from(document.querySelectorAll('.article-list .article-item .article-title')).map(article=>article.innerText).join("\n")`).String()

	fmt.Println(result)

	// Close the browser
	browser.MustClose()
}
