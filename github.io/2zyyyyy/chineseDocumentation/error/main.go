package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// 自定义异常
type CustomError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("path=%s \n op=%s \n createTime=%s message=%s",
		c.path, c.op, c.createTime, c.message)
}

func Open(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &CustomError{
			path:       fileName,
			op:         "read",
			createTime: fmt.Sprintf("%v\n", time.Now()),
			message:    err.Error(),
		}
	}
	defer file.Close()
	return nil
}

// system error
func systemError() {
	a := [5]int{1, 2, 3, 4, 5}
	a[1] = 123
	fmt.Println(a)
	index := 10
	a[index] = 10
	fmt.Println(a)
}

// 自己抛
func getCircleArea01(radius float32) (area float32) {
	if radius <= 0 {
		panic("半径必须大于0")
	}
	return 3.14 * radius * radius
}

// 返回异常
func getCircleArea02(radius float32) (area float32, err error) {
	if radius <= 0 {
		// 构建一个异常对象
		err = errors.New("半径必须大于0")
		return
	}
	area = 3.14 * radius * radius
	return
}

func test01() {
	// 延时执行匿名函数
	// 延时到何时 1.程序正常结束 2.发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Printf("defer fun() err=%s\n", err)
		}
	}()
	getCircleArea01(-1)
	fmt.Println("getCircleArea函数报错，此处不执行。")
}

func test02() {
	test01()
	fmt.Println("test02()")
}

func main() {
	// test02()
	area, err := getCircleArea02(1.2)
	if err != nil {
		fmt.Printf("err=%s\n", err)
	} else {
		fmt.Printf("area=%f\n", area)
	}

	//自定义error
	err = Open("/Users/gilbert/go/src/go/README1.md")
	switch v := err.(type) {
	case *CustomError:
		fmt.Println("get path error,", v)
	default:
	}
}
