package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

/*
runtime/pprof：采集工具型应用运行数据进行分析
net/http/pprof：采集服务型应用运行时数据进行分析
*/

// 一段有问题的代码
func logicCode() {
	var c chan int // c 未初始化默为：nil
	for {
		select {
		case v := <-c: // 阻塞
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)

	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}
