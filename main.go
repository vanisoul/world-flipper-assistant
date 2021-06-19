package main

import (
	"github.com/go-vgo/robotgo"
)

var yDifficulty = 0
var yBoss = 0
var imgBoss = "remaining.pngOrstroke.png"
var imgDifficulty = "updateList.pngOritemExchange.png"

func main() {

	//初始化設定 未來要判斷視窗位置
	infoScreen(-40, -40, 600, 1050)
	//初始化ID
	infoID()
	tmpRDifficulty := 0                   //重復打的難度
	tmpRNumber := 0                       //重復打的關卡
	tmpCNumber := 0                       //共鬥打的關卡
	tmpType := ""                         //哪一種模式
	tmpCMode := 0                         //共鬥的模式
	tmpPermanentPhysicalExertion := false //是否啟動常駐會消耗體力 自動判別有體力就去刷關

	tmpFreeRoom := "isNotFound" //判斷免費房間的標的圖片
	tmpAuto := ""               //紀錄auto狀態 如果是auto模式才有用到
	choseAuto := false          //紀錄這次有沒有被改變狀態
	choeseBossSeq := 0          //這次選擇的關卡
	for {
		//如果config有變動 需要重新回到主頁
		settingConfig, _ := LoadSettingConfig()
		if tmpRDifficulty != settingConfig.RDifficulty || tmpRNumber != settingConfig.RNumber || tmpType != settingConfig.Type || tmpCMode != settingConfig.CMode || tmpCNumber != settingConfig.CNumber || choseAuto || tmpPermanentPhysicalExertion != settingConfig.PermanentPhysicalExertion {
			tmpRDifficulty = settingConfig.RDifficulty
			tmpRNumber = settingConfig.RNumber
			tmpType = settingConfig.Type
			tmpCMode = settingConfig.CMode
			tmpCNumber = settingConfig.CNumber
			tmpPermanentPhysicalExertion = settingConfig.PermanentPhysicalExertion
			haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("ready.png"), getSystemImg("readyOK.png"), getSystemImg("readyA.png"), getSystemImg("readyAOK.png"), getSystemImg("goGame.png"), getSystemImg("stop.png")},
				func(x, y int) {
					haveOneImgsLeft(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsLeft(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsLeft(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsLeft(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsLeft(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsLeft(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsLeft(3, 0.05, true, getSystemImg("return.png"))
					haveOneImgsLeft(3, 0.05, true, getSystemImg("exitGreen.png"))
				},
				func(x, y int) {
					haveOneImgsLeft(3, 0.05, true, getSystemImg("return.png"))
				},
				func(x, y int) {
					leftMouseClick(x, y)
					haveOneImgsLeft(5, 0.01, false, getSystemImg("exit.png"))
					haveOneImgsLeft(5, 0.01, false, getSystemImg("exitYes.png"))
				})
			haveOneImgsLeft(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"), getSystemImg("main7.png"))
			//如果不等於重復關卡 就可以設定目前狀態 因為重復關卡可能是因為體力滿了而執行的AUTO
			if tmpAuto != "repalay" {
				tmpAuto = settingConfig.Type
			}
			choseAuto = false
		}

		//開啟遊戲
		haveOneImgsExecFunc(1, 0.03, false, []string{getSystemImg("gameLogo.png"), getSystemImg("joinMain.png"), getSystemImg("mainMission.png"), getSystemImg(imgBoss), getSystemImg(imgDifficulty), getSystemImg("YES.png"), getSystemImg("OK.png"), getSystemImg("dayGift.png"), getSystemImg("dayClose.png"), getSystemImg("updateList.png"), getSystemImg("ready.png"), getSystemImg("readyA.png"), getSystemImg("next1.png"), getSystemImg("next2.png"), getSystemImg("next3.png"), getSystemImg("next4.png"), getSystemImg("exitRoom.png"), getSystemImg("readyOK.png"), getSystemImg("readyAOK.png"), getSystemImg("exitHalfway.png"), getSystemImg("errorOK.png"), getSystemImg("goGame.png"), getSystemImg("gmaeOver.png"), getSystemImg("gmaeOverOK.png"), getSystemImg("LvUp.png"), getSystemImg("rePlay.png"), getSystemImg("lackOfEnergy.png"), getSystemImg("fiveStar.png"), getSystemImg("fourStar.png"), getSystemImg("threeStar.png"), getSystemImg("NEW.png"), getSystemImg("SKIP.png"), getSystemImg("dayGift2.png"), getSystemImg("closeRols.png")},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {

				now_type := ""
				if settingConfig.PermanentPhysicalExertion {
					now_type = tmpAuto
				} else {
					now_type = settingConfig.Type
				}

				if now_type == "freeBoss" {
					haveOneImgsLeft(10, 0.05, true, getSystemImg("joinBoss.png"))
					tmpFreeRoom = "freeRoom.png"
					imgBoss = "isNotFound"
					imgDifficulty = "isNotFound"
					yDifficulty = 0
					yBoss = 0
				} else if now_type == "freeActivity" {
					haveOneImgsLeft(10, 0.05, true, getSystemImg("joinActivity.png"))
					tmpFreeRoom = "toghterGo.png"
					imgBoss = "remaining.png"
					imgDifficulty = "isNotFound"
					yDifficulty = 310
					yBoss = 200
					choeseBossSeq = settingConfig.CNumber
				} else if now_type == "repalay" {
					haveOneImgsLeft(10, 0.05, true, getSystemImg("joinActivity.png"))
					imgBoss = "remaining.png"
					imgDifficulty = "gameProblem.png"
					yDifficulty = 200
					yBoss = 200
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
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg(tmpFreeRoom)}, func() {
					haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg(tmpFreeRoom)}, func(x, y int) {
						leftMouseClick(x-100, y)
					})
				}, func() {
					leftMouseClick(x, y)
				})
			},
			func(x, y int) {
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("fullOfEnergy.png")}, func() {
					tmpAuto = "repalay"
					choseAuto = true
				}, func() {
					leftMouseClick(x, y)
				})
			},
			func(x, y int) {
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("fullOfEnergy.png")}, func() {
					tmpAuto = "repalay"
					choseAuto = true
				}, func() {
					leftMouseClick(x, y)
				})
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				//如果是招募中 就看招募方式  以滿就直接開始
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("readyOK.png"), getSystemImg("recruiting.png")}, func() {
					if settingConfig.CMode == 1 {
						//直接開始
						haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
							leftMouseClick(x, y)
							leftMouseClick(x+80, y)
						})
					} else if settingConfig.CMode == 2 {
						//開放等人滿開始
						haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
							leftMouseClick(x, y)
							leftMouseClick(x+190, y)
						})
					}
				}, func() {
					haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
						leftMouseClick(x, y)
						leftMouseClick(x+80, y)
					})
				})
			},
			func(x, y int) {
				//如果是招募中 就看招募方式  以滿就直接開始
				haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("readyAOK.png"), getSystemImg("recruiting.png")}, func() {
					if settingConfig.CMode == 1 {
						//直接開始
						haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
							leftMouseClick(x, y)
							leftMouseClick(x+80, y)
						})
					} else if settingConfig.CMode == 2 {
						//開放等人滿開始
						haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
							leftMouseClick(x, y)
							leftMouseClick(x+190, y)
						})
					}
				}, func() {
					haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("dialogue.png")}, func(x, y int) {
						leftMouseClick(x, y)
						leftMouseClick(x+80, y)
					})
				})
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				haveOneImgsLeft(5, 0.05, false, getSystemImg("cancel.png"))
				tmpAuto = settingConfig.Type
				choseAuto = true
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			},
			func(x, y int) {
				leftMouseClick(x, y)
			})
	}
}

func choeseBoss(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yBoss + seq*110
	if seq < 6 {
		leftMouseClick(x, y)
	} else {
		ys := y_tmp + 310 + 4*110
		ye := y_tmp + 310 + 1*110
		robotgo.MoveMouse(x, ys)
		robotgo.MouseToggle(`down`, `left`)
		robotgo.Sleep(1)
		robotgo.MoveMouse(x, ye)
		robotgo.Sleep(1)
		robotgo.MouseToggle(`up`, `left`)
		choeseBoss(seq - 3)
		robotgo.Sleep(1)
	}
}

func choeseDifficulty(seq int) {
	_, _, x, y_tmp := findOneImages(1, 0.01, false, getSystemImg("stone.png"))
	y := y_tmp + yDifficulty + seq*110
	if seq < 6 {
		leftMouseClick(x, y)
	} else {
		ys := y_tmp + 310 + 2*110
		ye := y_tmp + 310 + 1*110
		robotgo.MoveMouse(x, ys)
		robotgo.MouseToggle(`down`, `left`)
		robotgo.Sleep(1)
		robotgo.MoveMouse(x, ye)
		robotgo.Sleep(1)
		robotgo.MouseToggle(`up`, `left`)
		choeseDifficulty(seq - 1)
		robotgo.Sleep(1)
	}
}
