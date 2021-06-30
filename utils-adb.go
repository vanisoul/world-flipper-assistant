package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"regexp"

	"os/exec"
	"strconv"
)

var adbHost string = ""

func adbinit(dev int) {
	MyCmd := exec.Command("adb", "-s", adbHost, "devices")
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	stdo := string(MyBytes)
	res := strings.Split(stdo, "device")

	for i, v := range res {
		if i == 0 {
			continue
		}
		if i == len(res)-1 {
			continue
		}
		if i == dev {
			port := strings.Split(v, ":")
			adbHost = fmt.Sprintf("127.0.0.1:%s", strings.Trim(port[1], "	"))
			break
		}
	}
	fmt.Println(adbHost)
}

//類比按鍵，如按下home鍵，鍵值參考；https://blog.csdn.net/shililang/article/details/14449527
//adb shell input keyevent 3
func AdbShellInputKeyEvent(s string) {
	exec.Command("adb", "-s", adbHost, "shell", "input", "keyevent", s).Run()
}

//類比螢幕點選
//有的控件死活抓不到，隻能直接點選
//adb shell input tap  900 800
func AdbShellInputTap(x, y int) {
	x2 := strconv.Itoa(x)
	y2 := strconv.Itoa(y)
	exec.Command("adb", "-s", adbHost, "shell", "input", "tap", x2, y2).Run()
}

//類比滑動
//adb shell input swipe  0 0  600 600
func AdbShellInputSwipe(x1, y1, x2, y2 int) {
	xx1 := strconv.Itoa(x1)
	yy1 := strconv.Itoa(y1)
	xx2 := strconv.Itoa(x2)
	yy2 := strconv.Itoa(y2)
	exec.Command("adb", "-s", adbHost, "shell", "input", "swipe", xx1, yy1, xx2, yy2).Run()
}

//類比長按 最後一個參數1000錶示1秒，可將下麵某個參數由500改為501，即允許坐標點有很小的變化。
//adb shell input swipe  500 500  500 500 1000
func AdbShellInputSwipeL(x1, y1, x2, y2, t int) {
	xx1 := strconv.Itoa(x1)
	yy1 := strconv.Itoa(y1)
	xx2 := strconv.Itoa(x2)
	yy2 := strconv.Itoa(y2)
	exec.Command("adb", "-s", adbHost, "shell", "swipe", "tap", xx1, yy1, xx2, yy2).Run()
}

//類比輸入"字符"
//adb shell input text "abc"
//若需輸入中文，可參考：https://blog.csdn.net/slimboy123/article/details/54140029
func AdbShellInputText(s string) {
	exec.Command("adb", "-s", adbHost, "shell", "input", "text", s).Run()
}

//截屏並保存到當前目錄下。
//由於需在手機和電腦上複製文件，必要時可增加延時或用下麵的PathExists()判斷文件是否存在，如：
//time.Sleep(time.Duration(2) * time.Second)
func AdbShellScreencapPullRm() {
	println("adb", "-s", adbHost, "shell", "screencap", "-p", "/sdcard/screen.png")
	exec.Command("adb", "-s", adbHost, "shell", "screencap", "-p", "/sdcard/screen.png").Run()
	exec.Command("adb", "-s", adbHost, "pull", "/sdcard/screen.png", ".").Run()
	exec.Command("adb", "-s", adbHost, "shell", "rm", "/sdcard/screen.png").Run()
}

//檢視手機上應用的packageName
//adb shell pm list packages
func AdbShellPmListPackages() string {
	MyCmd := exec.Command("adb", "-s", adbHost, "shell", "pm", "list", "packages")
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	s := string(MyBytes)
	return s
}

//通過adb 檢視最上層activity名字：
//adb shell dumpsys activity | findstr "mFocusedActivity"
//代碼中不能直接執行findstr過濾,改正則匹配
func AdbShellDumpsysActivityF() string {
	MyCmd := exec.Command("cmd.exe", "/c", "adb", "-s", adbHost, "shell", "dumpsys", "activity")
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	s := string(MyBytes)
	//正則匹配mFocusedActivity
	r := regexp.MustCompile(`mFocusedActivity.+?\}`)
	match := r.FindString(s)
	fmt.Println(match)
	return match
}
