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
	tmpStaminaRecoveryTime := 0           //Auto時回體時間
	tmpPermanentPhysicalExertion := false //是否啟動常駐會消耗體力 自動判別有體力就去刷關

	tmpAuto := "repalay" //紀錄auto狀態 如果是auto模式才有用到
	secAuto := 0         //紀錄回體時間
	choseAuto := false   //紀錄這次有沒有被改變狀態
	choeseBossSeq := 0   //這次選擇的關卡
	for {
		//如果config有變動 需要重新回到主頁
		settingConfig, _ := LoadSettingConfig()
		if tmpRDifficulty != settingConfig.RDifficulty || tmpRNumber != settingConfig.RNumber || tmpType != settingConfig.Type || tmpCMode != settingConfig.CMode || tmpStaminaRecoveryTime != settingConfig.StaminaRecoveryTime || tmpCNumber != settingConfig.CNumber || choseAuto || tmpPermanentPhysicalExertion != settingConfig.PermanentPhysicalExertion {
			tmpRDifficulty = settingConfig.RDifficulty
			tmpRNumber = settingConfig.RNumber
			tmpType = settingConfig.Type
			tmpCMode = settingConfig.CMode
			tmpCNumber = settingConfig.CNumber
			tmpStaminaRecoveryTime = settingConfig.StaminaRecoveryTime
			tmpPermanentPhysicalExertion = settingConfig.PermanentPhysicalExertion
			haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("startRaising.png")},
				func(x, y int) {
					haveOneImgsLeft(1, 0.05, true, getSystemImg("return.png"))
					haveOneImgsLeft(1, 0.05, true, getSystemImg("disband.png"))
				})
			haveOneImgsLeft(20, 0.1, false, getSystemImg("main1.png"), getSystemImg("main2.png"), getSystemImg("main3.png"), getSystemImg("main4.png"), getSystemImg("main5.png"), getSystemImg("main6.png"))
			choseAuto = false
		}

		//開啟遊戲
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("gameLogo.png")}, func(x, y int) {
			leftMouseClick(x, y)
		})
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("joinMain.png")}, func(x, y int) {
			leftMouseClick(x, y)
		})

		//進入bose戰鬥or活動戰鬥
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("mainMission.png")}, func(x, y int) {

			now_type := ""
			if settingConfig.PermanentPhysicalExertion {
				now_type = tmpAuto
			} else {
				now_type = settingConfig.Type
			}

			if now_type == "freeBoss" {
				haveOneImgsLeft(10, 0.05, true, getSystemImg("joinBoss.png"))
				imgBoss = "isNotFound"
				imgDifficulty = "isNotFound"
				yDifficulty = 0
				yBoss = 0
			} else if now_type == "freeActivity" {
				haveOneImgsLeft(10, 0.05, true, getSystemImg("joinActivity.png"))
				imgBoss = "remaining.png"
				imgDifficulty = "isNotFound"
				yDifficulty = 310
				yBoss = 200
				choeseBossSeq = settingConfig.CNumber
			} else if now_type == "repalay" {
				haveOneImgsLeft(10, 0.05, true, getSystemImg("joinActivity.png"))
				imgBoss = "remaining.png"
				imgDifficulty = "gameProblem.png"
				yDifficulty = 310
				yBoss = 200
				choeseBossSeq = settingConfig.RNumber
			}
		})

		//處理選擇boss關卡
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg(imgBoss)}, func(x, y int) { choeseBoss(choeseBossSeq) })

		//處理選擇難度
		haveOneImgsExecFunc(1, 0.01, false, []string{getSystemImg(imgDifficulty)}, func(x, y int) { choeseDifficulty(settingConfig.RDifficulty) })

		//使用強化點數
		haveOneImgsLeft(1, 0.05, true, getSystemImg("YES.png"))

		//閒置
		haveOneImgsLeft(1, 0.05, true, getSystemImg("OK.png"))

		//如果有需要領取獎勵 關閉活動通知
		haveOneImgsLeft(1, 0.05, true, getSystemImg("dayGift.png"))
		haveOneImgsLeft(1, 0.05, true, getSystemImg("dayClose.png"))

		//進入free到房間
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("updateList.png")}, func(x, y int) {
			haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("freeRoom.png")}, func() {
				haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("freeRoom.png")}, func(x, y int) {
					leftMouseClick(x-100, y)
				})
			}, func() {
				leftMouseClick(x, y)
			})
		})

		//尚未準備 開始執行準備
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("ready.png")}, func(x, y int) {
			leftMouseClick(x, y)
		})

		//打完的下一步加離開房間
		haveOneImgsLeft(1, 0.05, false, getSystemImg("next1.png"))
		haveOneImgsLeft(1, 0.05, false, getSystemImg("next2.png"))
		haveOneImgsLeft(1, 0.05, false, getSystemImg("next3.png"))
		haveOneImgsLeft(1, 0.05, false, getSystemImg("next4.png"))
		haveOneImgsLeft(1, 0.05, false, getSystemImg("exitRoom.png"))

		// 如果以準備 確認人數有沒有滿 沒滿要使用何種招募方式 如果滿就開始
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("readyOK.png")}, func(x, y int) {
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
				leftMouseClick(x, y)
				leftMouseClick(x+50, y)
			})
		})

		//有錯誤問題
		haveOneImgsLeft(1, 0.01, false, getSystemImg("exitHalfway.png"))
		haveOneImgsLeft(1, 0.01, false, getSystemImg("errorOK.png"))

		//有挑戰點挑戰
		haveOneImgsLeft(1, 0.01, false, getSystemImg("goGame.png"))

		//戰鬥失敗
		haveOneImgsLeft(1, 0.01, false, getSystemImg("gmaeOver.png"))
		haveOneImgsLeft(1, 0.01, false, getSystemImg("gmaeOverOK.png"))
		//LvUp
		haveOneImgsLeft(1, 0.01, false, getSystemImg("LvUp.png"))
		//rePlay
		haveOneImgsLeft(1, 0.01, false, getSystemImg("rePlay.png"))
		//auto相關判斷 如果是repalay狀態 則不計時
		//如果不是 就判斷大於等待時間後 重新去刷體力關卡
		secAuto = secAuto + 1
		if secAuto > tmpStaminaRecoveryTime {
			tmpAuto = "repalay"
			choseAuto = true
			secAuto = 0
		} else if tmpAuto == "repalay" {
			secAuto = 0
		}
		//如果體力不夠改變暫時的動作
		haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg("lackOfEnergy.png")}, func(x, y int) {
			haveOneImgsLeft(5, 0.05, false, getSystemImg("cancel.png"))
			tmpAuto = settingConfig.Type
			choseAuto = true
			secAuto = 0
		})

		robotgo.Sleep(3)
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
