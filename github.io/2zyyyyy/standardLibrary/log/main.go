package main

import (
	"fmt"
	"log"
	"os"
)

func logOutput() {
	logFile, err := os.OpenFile("./mylog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[测试日志前缀]")
	log.Println("这是一条很普通的日志。")
}

func logDemo() {
	log.Println("普通的 log~~")
	msg := "普通的"
	log.Printf("这是一条%s 的日志", msg)
	log.Fatalln("这是一条会触发 fatal 的日志")
	log.Panic("这是一条会触发 panic 的日志")
}

// 设置一下标准logger的输出选项
func logFlagSet() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条普通的日志")
}

// 配置日志前缀
func setPrefixDemo() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条普通的日志")
	// 设置输出前缀
	log.SetPrefix("[测试日志前缀]")
	log.Println("这是一条普通的日志")
}

func main() {
	//logDemo()
	//logFlagSet()
	//setPrefixDemo()
	logOutput()
}
