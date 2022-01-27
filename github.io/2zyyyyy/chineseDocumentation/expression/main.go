package main

import (
	"fmt"
)

// 表达式

type User struct {
	id   int
	name string
}

func (u *User) Test() {
	fmt.Printf("u.p=%p, u.v=%v\n", u, u)
}

func main() {
	user := User{
		100,
		"user.name",
	}
	user.Test()

	methodValue := user.Test
	methodValue() // 隐式传递 receiver

	methodExpression := (*User).Test
	methodExpression(&user) // 显式传递 receiver
}
