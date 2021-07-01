package main

import "github.com/go-vgo/robotgo"

func textLocation(x int, y int, text string) (succ bool) {
	AdbShellInputTap(x, y)
	robotgo.Sleep(1)
	AdbShellInputTap(x, y)
	robotgo.Sleep(1)
	AdbShellInputText(text)
	succ = true
	return

}
