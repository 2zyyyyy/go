package main

import "fmt"

// goto

func main() {

	// 常规操作，跳出多层for循环
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			fmt.Println("标记flag=true，跳出内循环~~")
			flag = true
			break // 跳出内层循环
		}
		if flag {
			fmt.Println("外层循环发现flag==true, 跳出外层循环！！")
			break // 跳出外层循环
		}
	}

	// 骚操作 goto
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				fmt.Println("跳到当前指定标签：breakTag~~~")
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

}
