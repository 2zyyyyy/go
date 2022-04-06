package main

import (
	"fmt"
	"time"
)

func timeNow() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间戳
func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

// 使用time.Unix()函数可以将时间戳转为时间格式。
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 求2个小时之后的时间：
func timeDifference() {
	now := time.Now()
	later := now.Add(time.Hour * 2) // 当前时间加 2 个小时后的时间
	fmt.Println(later)
}

// 定时器
func timer() {
	ticker := time.Tick(time.Second) // 定义一个间隔 1s 的定时器
	for j := range ticker {
		fmt.Println(j) // 每秒都会执行的任务
	}
}

// 时间格式化
func formatTime() {
	now := time.Now()
	// 格式化的模板为 Go 的出生时间（2006 1/2 15:04）
	// 24 小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12 小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func main() {
	//timeNow()
	//timestampDemo()
	//timestampDemo2(1649247038)
	//timeDifference()
	//timer()
	formatTime()
}
