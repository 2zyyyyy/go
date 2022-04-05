package main

import (
	"fmt"
	"os"
)

// 学生管理系统

// 声明全局的学生管理对象
var smr stuManager

// 菜单函数
func menu() {
	fmt.Println("--------------欢迎使用学生管理系统-------------")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出系统
	`)
}

func main() {
	var smr = stuManager{
		allStu: make(map[int64]student, 100),
	}
	for {
		menu()
		fmt.Println("请输入序号：")
		var chioce int
		fmt.Scan(&chioce)
		fmt.Println("你输入的是：", chioce)
		switch chioce {
		case 1:
			smr.queryStu()
		case 2:
			smr.addStu()
		case 3:
			smr.updateStu()
		case 4:
			smr.delStu()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入异常！！！")
		}
	}
}
