package main

import (
	"github.com/go-vgo/robotgo"
)

var yDifficulty = 0
var yBoss = 0
var yEvery = 0
var imgBoss = "remaining.pngOrstroke.png"
var imgDifficulty = "updateList.pngOritemExchange.png"

func main() {

	//初始化ID
	infoID()

	tmpFreeRoom := "isNotFound" //判斷免費房間的標的圖片
	tmpAuto := ""               //紀錄auto狀態 如果是auto模式才有用到
	choseAuto := false          //紀錄這次有沒有被改變狀態
	choeseBossSeq := 0          //這次選擇的關卡

	tmpSettingConfig, _ := LoadSettingConfig()
	tmpSettingConfig.Type = ""
	adbinit(tmpSettingConfig.Nox)

	for {
		//如果config有變動 需要重新回到主頁
		settingConfig, _ := LoadSettingConfig()
		if tmpSettingConfig != settingConfig || choseAuto {
			tmpSettingConfig = settingConfig
			haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("ready.png"), getSystemImg("readyOK.png"), getSystemImg("readyA.png"), getSystemImg("readyAOK.png"), getSystemImg("goGame.png"), getSystemImg("stop.png")},
				func(x, y int) {
					haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsClick(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsClick(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsClick(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsClick(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
				},
				func(x, y int) {
					mouseClick(x, y)
					haveOneImgsClick(5, 0.01, false, getSystemImg("exit.png"))
					haveOneImgsClick(5, 0.01, false, getSystemImg("exitYes.png"))
				})
			haveOneImgsClick(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"), getSystemImg("main7.png"))
			robotgo.Sleep(8)
			//如果不等於重復關卡 就可以設定目前狀態 因為重復關卡可能是因為體力滿了而執行的AUTO
			if tmpAuto != "repalay" {
				tmpAuto = settingConfig.Type
			}
			choseAuto = false
		}

		//開啟遊戲
		haveOneImgsExecFunc(1, 0.03, false, []string{getSystemImg("gameLogo.png"), getSystemImg("joinMain.png"), getSystemImg("mainMission.png"), getSystemImg(imgBoss), getSystemImg(imgDifficulty), getSystemImg("OK.png"), getSystemImg("YES.png"), getSystemImg("dayGift.png"), getSystemImg("dayClose.png"), getSystemImg("updateList.png"), getSystemImg("ready.png"), getSystemImg("next1.png"), getSystemImg("exitRoom.png"), getSystemImg("readyOK.png"), getSystemImg("goGame.png"), getSystemImg("rePlay.png")},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {

				now_type := ""
				if settingConfig.PermanentPhysicalExertion {
					now_type = tmpAuto
				} else {
					now_type = settingConfig.Type
				}

				if now_type == "freeBoss" {
					haveOneImgsClick(10, 0.05, true, getSystemImg("joinBoss.png"))
					tmpFreeRoom = "freeRoom.png"
					imgBoss = "isNotFound"
					imgDifficulty = "isNotFound"
					yDifficulty = 0
					yBoss = 0
					yEvery = 0
				} else if now_type == "freeActivity" {
					haveOneImgsClick(10, 0.05, true, getSystemImg("joinActivity.png"))
					tmpFreeRoom = "toghterGo.png"
					imgBoss = "remaining.png"
					imgDifficulty = "isNotFound"
					yBoss = 310
					yEvery = 178
					yDifficulty = 460
					choeseBossSeq = settingConfig.CNumber
				} else if now_type == "repalay" {
					haveOneImgsClick(10, 0.05, true, getSystemImg("joinActivity.png"))
					imgBoss = "remaining.png"
					imgDifficulty = "gameProblem.png"
					yBoss = 310
					yEvery = 178
					yDifficulty = 310
					choeseBossSeq = settingConfig.RNumber
				}
			},
			func(x, y int) {
				choeseBoss(choeseBossSeq)
			},
			func(x, y int) {
				choeseDifficulty(settingConfig.RDifficulty)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg(tmpFreeRoom)}, func() {
					haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg(tmpFreeRoom)}, func(x, y int) {
						mouseClick(x-100, y)
					})
				}, func() {
					mouseClick(x, y)
				})
			},
			func(x, y int) {
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("fullOfEnergy.png")}, func() {
					if settingConfig.PermanentPhysicalExertion {
						tmpAuto = "repalay"
						choseAuto = true
					} else {
						mouseClick(x, y)
					}
				}, func() {
					mouseClick(x, y)
				})
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				//如果是招募中 就看招募方式  以滿就直接開始
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("readyOK.png"), getSystemImg("recruiting.png")}, func() {
					if settingConfig.CMode == 1 {
						//直接開始
						haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
							mouseClick(x, y)
							mouseClick(x+140, y)
						})
					} else if settingConfig.CMode == 2 {
						//開放等人滿開始
						haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
							mouseClick(x, y)
							mouseClick(x+300, y)
						})
					}
				}, func() {
					haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
						mouseClick(x, y)
						mouseClick(x+140, y)
					})
				})
			},
			func(x, y int) {
				mouseClick(x, y)
			},
			func(x, y int) {
				mouseClick(x, y)
			})
	}
}

func choeseBoss(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yBoss + seq*yEvery
	if seq < 6 {
		mouseClick(x, y)
		robotgo.Sleep(3)
	} else {
		ys := y_tmp + yBoss + 4*yEvery
		ye := y_tmp + yBoss + 3*yEvery
		AdbShellInputSwipe(x, ys, x, ye)
		choeseBoss(seq - 1)
		robotgo.Sleep(1)
	}
}

func choeseDifficulty(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yDifficulty + seq*175
	if seq < 6 {
		mouseClick(x, y)
		robotgo.Sleep(3)
	} else {
		ys := y_tmp + yDifficulty + 2*175
		ye := y_tmp + yDifficulty + 1*175
		AdbShellInputSwipe(x, ys, x, ye)
		choeseDifficulty(seq - 1)
		robotgo.Sleep(1)
	}
}
