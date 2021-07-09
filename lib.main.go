package main

func addJoinMain(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinMain.png", funcs)
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

func addGameOverOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "GameOverOK.png", funcs)
	return
}

func addmainOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("OK.png"))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
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
func addJoinBoss2(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "joinBoss2.png", funcs)
	return
}

func addYES(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "YES.png", funcs)
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

func addThreeStar(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "threeStar.png", funcs)
	return
}
func addFiveStar(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "fiveStar.png", funcs)
	return
}

func addFourStar(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "fourStar.png", funcs)
	return
}
func addLvUP(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "LvUP.png", funcs)
	return
}

func addNoAuto(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "noAuto.png", funcs)
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
		findOneImgFuncStop(func() {
			mouseClick(x, y)
			mouseClick(x+140, y)
		}, 1, 0.05, false, getSystemImg("stop.png"), getSystemImg("OK.png"))
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
			findOneImgFuncLoop(func() {
				mouseClickNoTime(x, y)
				mouseClickNoTime(x+300, y)
			}, 1, 0.05, false, getSystemImg("recruiting.png"))
		}, func() {
			findOneImgFuncStop(func() {
				mouseClickNoTime(x, y)
				mouseClickNoTime(x+140, y)
			}, 1, 0.05, false, getSystemImg("stop.png"), getSystemImg("OK.png"))
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

func addLackOfEnergy(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg("lackOfEnergy.png"))
	funcs = append(funcs, func(x, y int) {
		status = 0
	})
	resStrs = strs
	resFuncs = funcs
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
func addDownloadOK(strs []string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	resStrs, resFuncs = clickBase(strs, "downloadOK.png", funcs)
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
