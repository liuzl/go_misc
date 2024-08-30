package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	// 启动无头浏览器
	l := launcher.New().Headless(true)
	url := l.MustLaunch()

	// 创建浏览器实例
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// 打开一个页面
	page := browser.MustPage("https://www.google.com")

	// 等待页面加载完成
	page.MustWaitLoad()

	// 截取页面截图并保存
	screenshot := page.MustScreenshot()
	utils.OutputFile("screenshot.png", screenshot)

	println("Screenshot saved as screenshot.png")
}
