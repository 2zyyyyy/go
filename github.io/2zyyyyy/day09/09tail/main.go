package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件那个地方开始读取
		MustExist: false,                                //文件不存在不报错
		Poll:      true,                                 //
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%s", err)
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen. fileName:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
