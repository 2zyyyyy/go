package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数
func main() {

	// 创建一个标志位参数
	// name := flag.String("name", "万里", "用户名")
	// age := flag.Int("age", 18, "年龄")
	// married := flag.Bool("married", false, "婚姻状况")
	// cTime := flag.Duration("ct", time.Second, "结婚时间")

	// TypeVar
	var name string
	var age int
	var married bool
	var cTime time.Duration

	flag.StringVar(&name, "name", "万里", "用户名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚姻状况")
	flag.DurationVar(&cTime, "ct", time.Second, "结婚时间")

	// 使用flag
	flag.Parse()
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(cTime)

	// flag其他函数
	fmt.Println(flag.Args())  ////返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
