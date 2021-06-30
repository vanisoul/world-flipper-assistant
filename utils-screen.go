package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

//一直執行func 當遇到圖片其一
//funcs : 被執行的方法,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 : 秒數可由rigorousSec修改,
//imgFullPaths : 傳入圖片相對全路徑 (多張)
func findOneImgFuncStop(funcs func(), frequency int, matchNumber float64, rigorous bool, imgFullPaths ...string) {
	for {
		funcs()
		findSucc, _, _, _ := findOneImages(frequency, matchNumber, rigorous, imgFullPaths...)
		if findSucc {
			return
		}
	}
}

//一直執行func 當遇到圖片全部
//funcs : 被執行的方法,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 : 秒數可由rigorousSec修改,
//imgFullPaths : 傳入圖片相對全路徑 (多張)
func findAllImgFuncStop(funcs func(), frequency int, matchNumber float64, rigorous bool, imgFullPaths ...string) {
	for {
		funcs()
		findSucc := findAllImages(frequency, matchNumber, rigorous, imgFullPaths...)
		if findSucc {
			return
		}
	}
}

//如果符合其中一張圖 就執行匹配到的第幾個func => func 有x y用為找到圖片的中心點,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 : 秒數可由rigorousSec修改,
//imgFullPaths : 傳入圖片相對全路徑 (多張),
//funcs : 可能被執行的方法
func haveOneImgsExecFunc(frequency int, matchNumber float64, rigorous bool, imgFullPaths []string, funcs ...func(x int, y int)) {
	findSucc, index, x, y := findOneImages(frequency, matchNumber, rigorous, imgFullPaths...)
	if findSucc {
		funcs[index](x, y)
	}
}

//如果符合全部圖片 就執行匹配到的func
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 : 秒數可由rigorousSec修改,
//imgFullPaths : 傳入圖片相對全路徑 (多張),
//funcs : 可能被執行的方法
func haveAllImgsExecFunc(frequency int, matchNumber float64, rigorous bool, imgFullPaths []string, funcY func(), funcN func()) {
	findSucc := findAllImages(frequency, matchNumber, rigorous, imgFullPaths...)
	if findSucc {
		funcY()
	} else {
		funcN()
	}
}

// 儲存定義的的視窗 ,
// paths 儲存至 log/${paths[0]}/${paths[1]}.....
func savescreen(name string) {

	imgName := fmt.Sprintf("log/%s", name)

	AdbShellScreencapPullRm()
	bitmap := robotgo.OpenBitmap("screen.png")
	defer robotgo.FreeBitmap(bitmap)
	robotgo.SaveBitmap(bitmap, imgName)
}
