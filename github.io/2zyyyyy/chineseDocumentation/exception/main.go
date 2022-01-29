package main

import (
	"errors"
	"fmt"
)

// go语言之异常

func main() {
	// test()
	// panicChannel(1)
	// deferPanic()
	// deferRecover()
	// testExcept()
	// protectFunc(2,1)
	// defer func() {
	// 	fmt.Println(recover())
	// }()
	// switch z, err := division(10, 0); err{
	// case nil:
	// 	fmt.Println(z)
	// case ErrorDieByZero:
	// 	panic(err)
	// }

	Try(func() {
		panic("test panic!")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

// GO try catch
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

// 标准库errors.New和fmt.Errorf函数用于创建实现error接口的错误对象。通过判断错误对象实例来确定具体错误类型
var ErrorDieByZero = errors.New("division by zero!")

func division(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrorDieByZero
	}
	return x / y, nil
}

// 如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执行
func protectFunc(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic!")
		z = x / y
		return
	}()
	fmt.Printf("x / y = %d\n", z)
}

// 使用延迟匿名函数或下面这样都是有效的。
func except() {
	fmt.Println(recover())
}

func testExcept() {
	defer except()
	panic("test panic!")
}

// 捕获函数recover只有在延迟调用内直接调用才会终止错误，否则总是返回nil。任何未捕获的错误都会沿调用堆栈向外传递
func deferRecover() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()

	defer recover() // 无效

	defer fmt.Println(recover()) // 无效

	defer func() {
		func() {
			fmt.Println("defer inner!")
			recover() // 无效
		}()
	}()
	panic("test panic!!")
}

// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func deferPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic error!")
	}()

	panic("deferPanic test!")
}

// 向已关闭的通道发送数据引发panic
func panicChannel(n int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch := make(chan int, 10)
	close(ch)
	ch <- n
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			// fmt.Printf("%T\n", err)
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型
		}
	}()
	panic("panic error!")
}
