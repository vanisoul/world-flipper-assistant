package main

import (
	"fmt"
	"path"

	"github.com/go-vgo/robotgo"
	"github.com/labstack/gommon/log"
)

var rigorousSec = 1

//搜尋多張圖片 只要符合一張,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 秒數可由rigorousSec更改,
//imgFullPaths : 傳入圖片相對全路徑 (多張),
//回傳 : 找到與否 - 找到的是第幾張 - 找到的中心點x - 找到的中心點y
func findOneImages(frequency int, matchNumber float64, rigorous bool, imgFullPaths ...string) (succ bool, index int, x int, y int) {
	log.Infof("findOneImages, frequency: %d, matchNumber: %e, rigorous: %s", frequency, matchNumber, rigorous)
	log.Infof("imgs", imgFullPaths)
	succ = false
	index = -1
	x = -1
	y = -1
	tmp_frequency := 0
	for tmp_frequency < frequency {
		robotgo.MoveMouse(0, 0)
		new_bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
		defer robotgo.FreeBitmap(new_bitmap)

		rigorous_bitmap := robotgo.CaptureScreen(0, 0, 0, 0)
		if rigorous {
			robotgo.Sleep(rigorousSec)
			rigorous_bitmap = robotgo.CaptureScreen(infox, infoy, infow, infoh)
		}
		defer robotgo.FreeBitmap(rigorous_bitmap)

		for idx, imgFullPath := range imgFullPaths {
			fileExist, imgw, imgh := imgFileExist(imgFullPath)
			if !fileExist {
				continue
			}
			dst_map := robotgo.OpenBitmap(imgFullPath)
			defer robotgo.FreeBitmap(dst_map)

			fx, fy := robotgo.FindBitmap(dst_map, new_bitmap, matchNumber)
			//如果沒找到就下一張圖
			if fx == -1 && fy == -1 {
				continue
			}
			//如果需要嚴謹 就在找一次 如果沒有 一樣進入下張圖
			if rigorous {
				fx, fy = robotgo.FindBitmap(dst_map, rigorous_bitmap, matchNumber)
				if fx == -1 && fy == -1 {
					continue
				}
			}
			succ = true
			index = idx
			x = fx + (imgw / 2) + infox
			y = fy + (imgh / 2) + infoy
			savescreen(path.Base(imgFullPath))
			log.Infof("findOneImages %s is found", imgFullPath)
			setLog("findOneImages", "尋找圖片成功", fmt.Sprintf("img: %s, x:%d, y:%d", imgFullPath, x, y))
			return
		}
		robotgo.Sleep(1)
		tmp_frequency = tmp_frequency + 1
	}
	return
}

//搜尋多張圖片 要尋找多張,
//frequency : 要搜尋幾次,
//matchNumber : 圖片符合的基準 越靠近1越容易匹配,
//rigorous : 是否需要持續1秒鐘以上圖片存在才算匹配成功 : 秒數可由rigorousSec修改,
//imgFullPaths : 傳入圖片相對全路徑 (多張),
//回傳 : 找到與否
func findAllImages(frequency int, matchNumber float64, rigorous bool, imgFullPaths ...string) (succ bool) {
	log.Infof("findAllImages, frequency: %d, matchNumber: %e, rigorous: %s", frequency, matchNumber, rigorous)
	log.Infof("imgs", imgFullPaths)
	succ = false

	filesExist := imgFilesExist(imgFullPaths...)
	if !filesExist {
		return
	}

	tmp_frequency := 0

	for tmp_frequency < frequency {
		robotgo.MoveMouse(0, 0)
		new_bitmap := robotgo.CaptureScreen(infox, infoy, infow, infoh)
		defer robotgo.FreeBitmap(new_bitmap)

		rigorous_bitmap := robotgo.CaptureScreen(0, 0, 0, 0)
		if rigorous {
			robotgo.Sleep(rigorousSec)
			rigorous_bitmap = robotgo.CaptureScreen(infox, infoy, infow, infoh)
		}
		defer robotgo.FreeBitmap(rigorous_bitmap)

		isbreak := false
		for _, imgFullPath := range imgFullPaths {
			dst_map := robotgo.OpenBitmap(imgFullPath)
			defer robotgo.FreeBitmap(dst_map)

			fx, fy := robotgo.FindBitmap(dst_map, new_bitmap, matchNumber)
			//如果沒找到就離開這次找圖 等待一秒下一次再一次掃圖加找圖
			if fx == -1 && fy == -1 {
				isbreak = true
				break
			}
			//如果需要嚴謹 就在找一次 如果沒找到就離開這次找圖 等待一秒下一次再一次掃圖加找圖
			if rigorous {
				fx, fy = robotgo.FindBitmap(dst_map, rigorous_bitmap, matchNumber)
				if fx == -1 && fy == -1 {
					isbreak = true
					break
				}
			}
			savescreen(path.Base(imgFullPath))
			log.Infof("findAllImages %s is found", imgFullPath)
			setLog("findAllImages", "尋找圖片成功", fmt.Sprintf("img: %s", imgFullPath))
		}
		tmp_frequency = tmp_frequency + 1
		if !isbreak {
			succ = true
			return
		}
	}
	return
}
