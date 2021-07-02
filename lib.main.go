package main

func addGameLogo(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "gameLogo.png", funcs)
	return
}

func addJoinMain(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinMain.png", funcs)
	return
}

func addMainMission(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("mainMission.png"))
	funcs = append(funcs, func(x, y int) {
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
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addMainToActivity(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("mainMission.png"))
	funcs = append(funcs, func(x, y int) {
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
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addImgBoss(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg(imgBoss))
	funcs = append(funcs, func(x, y int) {
		choeseBoss(choeseBossSeq)
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addImgDifficulty(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg(imgDifficulty))
	funcs = append(funcs, func(x, y int) {
		choeseDifficulty(settingConfig.RDifficulty)
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("OK.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
		status = 6
	})
	resStrs = strs
	resFuncs = funcs
	return

}

func addJoinActivity(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinActivity.png", funcs)
	return
}

func addJoinBoss(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinBoss.png", funcs)
	return
}

func addYES(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "YES.png", funcs)
	return
}

func addDayGift(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "dayGift.png", funcs)
	return
}

func addDayClose(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "dayClose.png", funcs)
	return
}

func addJoinfreeRoom(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("updateList.png"))
	funcs = append(funcs, func(x, y int) {
		haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg(tmpFreeRoom)}, func() {
			haveOneImgsExecFunc(1, 0.05, false, []string{getSystemImg(tmpFreeRoom)}, func(x, y int) {
				mouseClick(x-100, y)
			})
		}, func() {
			mouseClick(x, y)
		})
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addReady(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("ready.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
		status = 66
	})
	resStrs = strs
	resFuncs = funcs
	return

}

func addFullOfEnergy(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("fullOfEnergy.png"))
	funcs = append(funcs, func(x, y int) {
		status = 0
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addNotFullOfEnergyMain(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("mainMission.png"))
	funcs = append(funcs, func(x, y int) {
		if settingConfig.Type == "freeActivity" {
			tmpFreeRoom = "toghterGo.png"
			imgBoss = "remaining.png"
			imgDifficulty = "isNotFound"
			yBoss = 310
			yEvery = 178
			yDifficulty = 460
			choeseBossSeq = settingConfig.CNumber
			status = 6
		} else if settingConfig.Type == "freeBoss" {
			tmpFreeRoom = "freeRoom.png"
			imgBoss = "isNotFound"
			imgDifficulty = "isNotFound"
			yDifficulty = 0
			yBoss = 0
			yEvery = 0
			status = 6
		} else if settingConfig.Type == "repalay" {
			imgBoss = "remaining.png"
			imgDifficulty = "gameProblem.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 310
			choeseBossSeq = settingConfig.RNumber
			status = 5
		} else {
			imgBoss = "isNotFound"
			imgDifficulty = "isNotFound"
			yBoss = -1
			yEvery = -1
			yDifficulty = -1
			choeseBossSeq = -1
			status = 0
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addFullOfEnergyMain(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("fullOfEnergy.png"))
	funcs = append(funcs, func(x, y int) {
		if settingConfig.PermanentPhysicalExertion {
			imgBoss = "remaining.png"
			imgDifficulty = "gameProblem.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 310
			choeseBossSeq = settingConfig.RNumber
			status = 5
		} else if settingConfig.Type == "freeActivity" {
			tmpFreeRoom = "toghterGo.png"
			imgBoss = "remaining.png"
			imgDifficulty = "isNotFound"
			yBoss = 310
			yEvery = 178
			yDifficulty = 460
			choeseBossSeq = settingConfig.CNumber
			status = 6
		} else if settingConfig.Type == "freeBoss" {
			tmpFreeRoom = "freeRoom.png"
			imgBoss = "isNotFound"
			imgDifficulty = "isNotFound"
			yDifficulty = 0
			yBoss = 0
			yEvery = 0
			status = 7
		} else if settingConfig.Type == "repalay" {
			imgBoss = "remaining.png"
			imgDifficulty = "gameProblem.png"
			yBoss = 310
			yEvery = 178
			yDifficulty = 310
			choeseBossSeq = settingConfig.RNumber
			status = 5
		} else {
			imgBoss = "isNotFound"
			imgDifficulty = "isNotFound"
			yBoss = -1
			yEvery = -1
			yDifficulty = -1
			choeseBossSeq = -1
			status = 0
		}
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addNext1(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "next1.png", funcs)
	return
}

func addExitRoom(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("exitRoom.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
		status = 6
	})
	resStrs = strs
	resFuncs = funcs
	return

}

func addModeGo(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("dialogue.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
		mouseClick(x+140, y)
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addModeWait(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("dialogue.png"))
	funcs = append(funcs, func(x, y int) {
		//如果是招募中 就看招募方式  以滿就直接開始
		haveAllImgsExecFunc(1, 0.05, false, []string{getSystemImg("recruiting.png")}, func() {
			//開放等人滿開始
			mouseClick(x, y)
			mouseClick(x+300, y)
		}, func() {
			mouseClick(x, y)
			mouseClick(x+140, y)
		})
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addReadyOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("readyOK.png"))
	funcs = append(funcs, func(x, y int) {
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
	})
	resStrs = strs
	resFuncs = funcs
	return
}
func addGoGame(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "goGame.png", funcs)
	return
}
func addGoGameMaze(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "goGameMaze.png", funcs)
	return
}
func addRePlay(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "rePlay.png", funcs)
	return
}
func addNetworkerrorOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "networkerrorOK.png", funcs)
	return
}

func addRready(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("ready.png"))
	funcs = append(funcs, func(x, y int) {
		haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
		haveOneImgsClick(3, 0.05, true, getSystemImg("exitGreen.png"))
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addLackOfEnergy(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("lackOfEnergy.png"))
	funcs = append(funcs, func(x, y int) {
		status = 0
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addRreadyOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("readyOK.png"))
	funcs = append(funcs, func(x, y int) {
		haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
		haveOneImgsClick(3, 0.05, true, getSystemImg("exitGreen.png"))
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addRgoGame(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("goGame.png"))
	funcs = append(funcs, func(x, y int) {
		haveOneImgsClick(3, 0.05, true, getSystemImg("return.png"))
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addRstop(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("stop.png"))
	funcs = append(funcs, func(x, y int) {
		haveOneImgsClick(5, 0.01, false, getSystemImg("exit.png"))
		haveOneImgsClick(5, 0.01, false, getSystemImg("exitYes.png"))
	})
	resStrs = strs
	resFuncs = funcs
	return
}

func addMain2(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "main2.png", funcs)
	return
}

func addGmaeOver(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "gmaeOver.png", funcs)
	return
}
func addExitHalfway(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "exitHalfway.png", funcs)
	return
}

func addMainMissionMainOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("mainMission.png"))
	funcs = append(funcs, func(x, y int) {
		status = 2
	})
	resStrs = strs
	resFuncs = funcs
	return
}
