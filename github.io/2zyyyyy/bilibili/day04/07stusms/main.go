package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	实现系统：允许 查看/新增/删除 学生
*/
var (
	allStudent map[int64]*student // 变量声明
)

// 定义student类型
type student struct {
	id   int64
	name string
}

// 定义student构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

// 查看所有学生数据
func showAllStudent() {
	// 打印所有学生
	fmt.Println(allStudent)
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

// 创建学生
func addStudent() {
	// 向allStudent中添加一个新的学生
	// 1、创建一个学生 2、获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生的姓名：")
	fmt.Scanln(&name)
	// 3.创造学生（）调用newStudent构造函数
	newStu := newStudent(id, name)
	// 4.追加学生数据到allStudent这个map中
	allStudent[id] = newStu
}

// 删除学生
func delStudent() {
	// 1、请输入要删除的学生的学号
	var (
		delID int64
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&delID)
	// 2、去allStuden这个map中根据id删除对应的键值对
	delete(allStudent, delID)
}

func main() {
	allStudent = make(map[int64]*student, 48) // 初始化，申请内存空间
	for {
		/*
			1.输出功能菜单
			2.等待用户选择菜单
			3.执行对应函数
		*/
		// 1.输出功能菜单
		fmt.Println("欢迎使用学生管理系统V1.0.0~~")
		fmt.Println(`
			1、查看学生数据
			2、创建学生
			3、删除学生
			4、退出系统
		`)
		// 2.等待用户选择菜单
		fmt.Print("输入菜单编号选择对应用功能：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项！\n", choice)
		// 3.执行对应函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入不合法！")
		}
	}
}
