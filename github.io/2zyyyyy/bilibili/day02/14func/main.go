package main

import "fmt"

/* golang 函数
函数是一段代码的封装
把一段逻辑抽象出来封装到一个函数，定义函数名称，每次使用调用该函数即可
使用函数可以使代码更加清晰、简洁 */

// 函数的定义
func sum(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值可以直接return
}

// 无返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 无参无返回
func f2() {
	fmt.Println("f2")
}

// 无参有返回值
func f3() int {
	return 3
}

func f4() (int, string) {
	return 1, "golang~~~"
}

// 参数的类型简写
func f5(x, y int) int {
	return x + y
}

// 可变长参数
func f6(x string, y ...int) { // 可变长参数必须放在参数最后
	fmt.Println(x)
	fmt.Println(y) // y的类型是slice 切片[]int
}

func main() {
	var s = sum(10, 20)
	fmt.Println(s)

	_, n := f4()
	fmt.Println(n)

	f6("python")
	f6("python", 1, 2, 3, 4, 5)

	/* 回文判断
	山西运煤车煤运西山
	s[0] = s[(len(s) - 1)] */
	ss := "山西运煤车煤运西山1"
	/* 山 ss[0] ss[len(ss) - 1]
	山 ss[1] ss[len(ss) - 1 - 1]
	山 ss[2] ss[len(ss) - 1 - 2]
	山 ss[30] ss[len(ss) - 1 - 3] */

	r := make([]rune, 0, len(ss))
	for _, v := range ss {
		r = append(r, v)
	}
	fmt.Println(r)

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文！！！")
			return
		}
	}
	fmt.Println("是回文~~~")

}
