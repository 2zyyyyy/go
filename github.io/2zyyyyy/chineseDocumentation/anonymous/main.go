package main

import (
	"fmt"
)

// 匿名字段

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (u *User) toString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

// 通过匿名字段，可获得和继承类似的复用能力。依据编译器查找顺序，只需在外层定义同名方法，就可以实现”override“。
func (m *Manager) toString() string {
	return fmt.Sprintf("Manager: %p, %v", m, m)
}

func main() {
	manager := Manager{User{
		100,
		"法外狂徒张三",
	}, "manager title"}

	fmt.Printf("manager type=%T, manager=%p\n", manager, &manager)
	fmt.Println(manager.toString())
	fmt.Println(manager.User.toString())
}
