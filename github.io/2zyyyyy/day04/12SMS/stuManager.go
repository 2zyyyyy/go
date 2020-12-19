package main

import "fmt"

/* 学生管理系统
1、他保存了一些数据（结构体字段）
2、他有四个功能（结构体方法） */

type student struct {
	id   int64
	name string
}

// 造一个学生管理者
type stuManager struct {
	allStu map[int64]student
}

// 四个方法，学生的增删改查

// 增加学生
func (s stuManager) addStu() {
	// 根据用户输入的内容创建一个新的学生 将该学生放入map中
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)

	// 根据用户输入创造结构体对象
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	// 把新的学生放到s.allStu这个map中
	s.allStu[newStu.id] = newStu
}

// 删除学生
func (s stuManager) delStu() {
	/* 1.获取用户输入的学号
	2.如果查询不到输出查无此人
	3.如果匹配到用户数据执行删除操作 */
	var stuID int64
	fmt.Print("请输入要删除学生的学号：")
	fmt.Scanln(&stuID)
	_, ok := s.allStu[stuID]
	if !ok {
		fmt.Println("查无此人！！！")
		return
	}
	delete(s.allStu, stuID)
	fmt.Println("删除成功！")
}

// 修改学生
func (s stuManager) updateStu() {
	/* 	1.获取用户输入的学号
	   	2.展示该学号对应的学生，如没有输出查无此人
	   	3.让用户输入修改后的学生名字
	   	4.更新学生的姓名 */
	var stuID int64
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&stuID)
	stuObj, ok := s.allStu[stuID]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	fmt.Print("请输入修改后的学生姓名：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStu[stuID] = stuObj
}

// 查看学生
func (s stuManager) queryStu() {
	// 从s.allStu中把所有的学生遍历出来
	for _, stu := range s.allStu { // stu=每个 学生
		fmt.Printf("学号%d 姓名%s\n", stu.id, stu.name)
	}
}
