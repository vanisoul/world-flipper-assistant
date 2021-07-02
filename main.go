package main

import (
	"github.com/go-vgo/robotgo"
)

var yDifficulty = 0
var yBoss = 0
var yEvery = 0
var imgBoss = "remaining.pngOrstroke.png"
var imgDifficulty = "updateList.pngOritemExchange.png"
var tmpFreeRoom = "isNotFound" //判斷免費房間的標的圖片
var tmpAuto = ""               //紀錄auto狀態 如果是auto模式才有用到
var choeseBossSeq = 0          //這次選擇的關卡
var settingConfig SettingConfig
var status = 0   //0為重開 1為到首頁 2為確認已到首頁 3為bossfree 4為自動耗體
var notthink = 0 //多少次迴圈沒動作 1000次沒動作 就關閉重啟

func main() {

	//初始化ID
	infoID()

	tmpSettingConfig, _ := LoadSettingConfig()
	tmpSettingConfig.Type = ""

	for {
		//如果config有變動 遊戲重開
		settingConfig, _ = LoadSettingConfig()
		if tmpSettingConfig != settingConfig {
			tmpSettingConfig = settingConfig
			adbinit(settingConfig.Nox)
			status = 0
		}

		//重開遊戲
		if status == 0 {
			//關閉遊戲
			closeApp()
			robotgo.Sleep(2)
			//啟動遊戲
			openApp()

			status = 1
		}

		// 到首頁
		if status == 1 {
			toMainImg := []string{}
			toMainFunc := []func(x int, y int){}
			toMainImg, toMainFunc = addJoinMain(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addmainOK(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addExitHalfway(toMainImg, toMainFunc)
			toMainImg, toMainFunc = addMainMissionMainOK(toMainImg, toMainFunc)
			haveOneImgsExecFunc(1, 0.05, false, toMainImg, toMainFunc...)
		}

		//確認已到首頁
		if status == 2 {
			//判斷體力是否滿
			dowhatImg := []string{}
			dowhatFunc := []func(x int, y int){}
			//如果體力滿就狀態5 消耗體力
			if settingConfig.PermanentPhysicalExertion {
				dowhatImg, dowhatFunc = addFullOfEnergyMain(dowhatImg, dowhatFunc)
			}
			dowhatImg, dowhatFunc = addNotFullOfEnergyMain(dowhatImg, dowhatFunc)

			haveOneImgsExecFunc(1, 0.05, false, dowhatImg, dowhatFunc...)
		}

		//消耗體力
		if status == 5 {
			runActivityImgImg := []string{}
			runActivityImgFunc := []func(x int, y int){}
			runActivityImgImg, runActivityImgFunc = addJoinActivity(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addImgBoss(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addImgDifficulty(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addYES(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addGoGame(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addGoGameMaze(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addGmaeOver(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addOK(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addNext1(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addRePlay(runActivityImgImg, runActivityImgFunc)
			runActivityImgImg, runActivityImgFunc = addLackOfEnergy(runActivityImgImg, runActivityImgFunc)
			haveOneImgsExecFunc(1, 0.05, false, runActivityImgImg, runActivityImgFunc...)
		}

		//活動共鬥freeActivity
		if status == 6 {
			runActivityFreeImg := []string{}
			runActivityFreeFunc := []func(x int, y int){}
			if settingConfig.PermanentPhysicalExertion {
				runActivityFreeImg, runActivityFreeFunc = addFullOfEnergy(runActivityFreeImg, runActivityFreeFunc)
			}
			if settingConfig.Type == "freeActivity" {
				runActivityFreeImg, runActivityFreeFunc = addJoinActivity(runActivityFreeImg, runActivityFreeFunc)
			} else if settingConfig.Type == "freeBoss" {
				runActivityFreeImg, runActivityFreeFunc = addJoinBoss(runActivityFreeImg, runActivityFreeFunc)
			}
			runActivityFreeImg, runActivityFreeFunc = addImgBoss(runActivityFreeImg, runActivityFreeFunc)
			runActivityFreeImg, runActivityFreeFunc = addImgDifficulty(runActivityFreeImg, runActivityFreeFunc)
			runActivityFreeImg, runActivityFreeFunc = addJoinfreeRoom(runActivityFreeImg, runActivityFreeFunc)
			runActivityFreeImg, runActivityFreeFunc = addYES(runActivityFreeImg, runActivityFreeFunc)
			runActivityFreeImg, runActivityFreeFunc = addReady(runActivityFreeImg, runActivityFreeFunc)

			haveOneImgsExecFunc(1, 0.05, false, runActivityFreeImg, runActivityFreeFunc...)
		}

		//在房間內
		if status == 66 {
			runActivityFreeImg := []string{}
			runActivityFreeFunc := []func(x int, y int){}

			if settingConfig.CMode == 1 {
				runActivityFreeImg, runActivityFreeFunc = addModeGo(runActivityFreeImg, runActivityFreeFunc)
			} else if settingConfig.CMode == 2 {
				runActivityFreeImg, runActivityFreeFunc = addModeWait(runActivityFreeImg, runActivityFreeFunc)
			}

			runActivityFreeImg, runActivityFreeFunc = addNext1(runActivityFreeImg, runActivityFreeFunc)

			runActivityFreeImg, runActivityFreeFunc = addOK(runActivityFreeImg, runActivityFreeFunc)
			runActivityFreeImg, runActivityFreeFunc = addExitRoom(runActivityFreeImg, runActivityFreeFunc)

			haveOneImgsExecFunc(1, 0.05, false, runActivityFreeImg, runActivityFreeFunc...)
		}

		if notthink > 1000 {
			status = 0
			notthink = 0
		}
		notthink++
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
