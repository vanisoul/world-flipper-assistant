package main

import "github.com/go-vgo/robotgo"

//如果符合其中一張圖 就執行滑鼠右鍵,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 : 秒數可由rigorousSec修改,
//imgFullPaths : 傳入圖片相對全路徑 (可多張)
func haveOneImgsClick(frequency int, matchNumber float64, rigorous bool, imgFullPaths ...string) {
	findSucc, _, x, y := findOneImages(frequency, matchNumber, rigorous, imgFullPaths...)
	if findSucc {
		mouseClick(x, y)
	}
}

func mouseClick(x int, y int) {
	AdbShellInputTap(x, y)
	notthink = 0
	robotgo.Sleep(1)
}
