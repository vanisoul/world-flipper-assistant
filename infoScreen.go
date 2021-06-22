package main

import "fmt"

var infox int = 0
var infoy int = 0
var infow int = 20000
var infoh int = 20000

//手機板 900*1600

//初始化本程式未來要找的座標範圍 (不找全圖)
//需要有initImg.png  x:掃描範圍x與initImg.png位置偏移量 y:掃描範圍y與initImg.png位置偏移量 w:寬度 h:高度
func infoScreen(x int, y int, w int, h int) bool {
	succ, _, xx, yy := findOneImages(5, 0.01, false, getSystemImg("initImg.png"))
	if succ {
		infox = xx + x
		infoy = yy + y
		infow = w
		infoh = h
		setLog("infoScreen", "設定視窗位置成功", fmt.Sprintf("%d, %d, %d, %d", infox, infoy, infow, infoh))
		return true
	}
	setLog("infoScreen", "設定視窗位置失敗", fmt.Sprintf("%d, %d, %d, %d", infox, infoy, infow, infoh))
	return false
}
