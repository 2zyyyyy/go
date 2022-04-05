package main

import "fmt"

// 复合数据类型

/*
数组
定义：var 数组变量名 [元素数量]T (类型)
存放元素的容器
必须制定存放元素的类型和长度
golang：数组的长度是数组类型的一部分
*/

func main() {
	var list1 [5]int
	var list2 [10]int

	fmt.Printf("list1:%T list2:%T\n", list1, list2)

	// 数组的初始化 不初始化默认是0   int: 0  string:“”  bool:false
	fmt.Println(list1, list2)
	// 1、初始化方式01  创建变量的时候赋值
	list1 = [5]int{0, 1, 2, 3, 4}
	fmt.Println(list1)
	// 2、初始化方式02 根据初始值自动腿短数组的长度是多少
	list100 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\n", list100)

	// 3、初始化方式03  根据索引初始化
	list3 := [5]int{0: 1, 4: 2}
	fmt.Println(list3)

	// 数组的遍历
	city := [...]string{"杭州", "宁波", "金华"}
	// 1、根据索引
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	// 2、for range 遍历
	for i, v := range city {
		fmt.Println(i, v)
	}

	// 多维数组
	// [[1 2] [3 4] [5 6]]
	var all [3][2]int
	all = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(all)
	// 多维数组的遍历
	for _, v1 := range all {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

	// 练习题1 求数组[1,3,5,7,8]元素的和
	c1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range c1 {
		sum = sum + v
	}
	fmt.Println(sum)

	/* 找出数组中和位置性质的两个元素的下标，比如[1, 3, 5, 7, 8]中找出和为8两个元素的下标分别为(0,3)(1,2)
	定义2个for循环  外层从第n个开始遍历
	内层for循环从n+1开始找
	1和2的和为8 */
	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			if c1[i]+c1[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}

}
