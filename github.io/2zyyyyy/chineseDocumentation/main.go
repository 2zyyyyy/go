package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	//_ "go/github.io/2zyyyyy/hello"
)

// 下划线在代码中
func underline() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/gilbert/go/src/go/golang基础.md")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		_, _ = os.Stdout.Write(buf[:n])
	}
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 修改字符串
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	var c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func sumArray(a [10]int) int {
	var sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

// 找出数组中和为给定值的两个元素的下标
func testsSum(n []int, sum int) {
	for i := 0; i < len(n); i++ {
		for j := i; j < len(n); j++ {
			if n[i]+n[j] == sum {
				fmt.Println(i, j)
			}
		}
	}
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	// import使用了'_',则会编译报错：./main.go:9:2: undefined: hello
	// hello.Hello()
	// underline()

	// traversalString()
	// 若想做一个真正的随机数，要种子 seed()种子默认是1 rand.Seed(1)
	// rand.Seed(time.Now().Unix())

	// var b [10]int
	// for i := 0; i < len(b); i++ {
	// 	// 产生一个1~1000随机数
	// 	b[i] = rand.Intn(1000)
	// }
	// fmt.Println(b)
	// sum := sumArray(b)
	// fmt.Printf("sum:%d\n", sum)

	// n := []int{0, 1, 3, 5, 7, 8}
	// fmt.Println(n)
	// testsSum(n, 8)

	// data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s := data[:2:3]
	// fmt.Println(s, len(s), cap(s))

	// s = append(s, 10, 20)        // 添加2个值，超出容量3限制
	// fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	// fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。

	// s11 := make([]int, 0, 1)
	// c := cap(s11)
	// fmt.Printf("初始容量：%d\n", c)

	// for i := 0; i < 50; i++ {
	// 	// 追加
	// 	s11 = append(s11, i)
	// 	if n := cap(s11); n > c {
	// 		fmt.Printf("追加前cap为: %d -> 追加后cap为:%d\n", c, n)
	// 		c = n
	// 	}
	// }

	// s1 := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	// fmt.Printf("slice s1:%v\n", s1)

	// s2 := s1[len(s1)-1:]
	// fmt.Printf("slice s2:%v\n", s2)

	// copy(s2, s1) // 只会复制s1的第一个元素到s2中
	// fmt.Printf("copied slice s1 : %v\n", s1)
	// fmt.Printf("copied slice s2 : %v\n", s2)

	// s3 := []int{1, 2, 3}
	// fmt.Printf("slice s3:%v\n", s3)
	// s3 = append(s3, s2...)
	// fmt.Printf("append s3:%v\n", s3)
	// s3 = append(s3, 4, 5, 6)
	// fmt.Printf("last s3:%v\n", s3)

	// data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println("array data:", data)

	// s1 := data[8:]
	// s2 := data[:5]
	// fmt.Printf("slice s1:%v\n", s1)
	// fmt.Printf("slice s2:%v\n", s2)

	// copy(s2, s1)
	// fmt.Printf("copied slice s1 : %v\n", s1)
	// fmt.Printf("copied slice s2 : %v\n", s2)
	// fmt.Println("last array data : ", data)

	// data12 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s12 := data12[:]
	// for i, v := range s12 {
	// 	fmt.Printf("i:%v, v:%v\n", i, v)
	// }

	// a12 := []int{1, 2, 3, 4}
	// fmt.Printf("slice a12:%v, len(a12):%v\n", a12, len(a12))

	// b12 := a12[1:2]
	// fmt.Printf("slice b12:%v, len(b12):%v\n", b12, len(b12))

	// c12 := b12[0:3]
	// fmt.Printf("slice c12:%v, len(c12):%v\n", c12, len(c12))

	// str := "software tester"
	// s1 := str[0:8]
	// fmt.Println(s1)

	// s2 := str[9:]
	// fmt.Println(s2)

	// str := "software tester"
	// s13 := []byte(str) // 中文字符需要用rune
	// s13[0] = 'S'
	// s13 = append(s13, '!')
	// str = string(s13)
	// fmt.Println(str)

	// str := "欲穷千里日，更上一层娄！"
	// s := []rune(str)
	// s[4] = '目'
	// s[10] = '楼'
	// str = string(s)
	// fmt.Println(str)

	// slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// data1 := slice[6:8]
	// fmt.Printf("data1：%v, len(data1):%d, cap(data1):%d\n", data1, len(data1), cap(data1))

	// data2 := slice[:6:8]
	// fmt.Printf("data2：%v, len(data2):%d, cap(data2):%d\n", data2, len(data2), cap(data2))

	// a20 := 10
	// b20 := &a20
	// fmt.Printf("a20:%d, ptr:%p\n", a20, &a20)
	// fmt.Printf("b20:%p, type:%T\n", b20, b20)
	// fmt.Printf("b20:%p, *b20:%d\n", &b20, *b20)

	// a20 := 10
	// modify1(a20)
	// fmt.Println(a20)

	// modify2(&a20)
	// fmt.Println(a20)
	// var a *int = new(int)
	// *a = 100
	// fmt.Println(*a)

	// var b map[string]int
	// b["测试"] = 100
	// fmt.Println(b)

	// 指针练习
	// 程序定义一个int变量num的地址并打印
	// 将num的地址赋给指针ptr，并通过ptr去修改num的值
	var num int
	fmt.Println(&num)
	ptr := &num
	fmt.Println(reflect.TypeOf(ptr))
	*ptr = 20
	fmt.Println(num)
}
