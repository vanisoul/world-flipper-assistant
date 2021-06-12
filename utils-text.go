package main

import "github.com/go-vgo/robotgo"

func textLocation(x int, y int, text string) (succ bool) {
	robotgo.MoveMouse(x, y)
	robotgo.MouseClick()
	robotgo.Sleep(1)
	robotgo.MouseClick()
	robotgo.Sleep(1)
	robotgo.TypeStr(text, 10)
	succ = true
	return

}
