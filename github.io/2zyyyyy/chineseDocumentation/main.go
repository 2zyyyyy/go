package main

import (
	"fmt"
	"go/github.io/2zyyyyy/hello"
	"math"
	"math/rand"
	"os"
	"time"
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

func main() {
	// import使用了'_',则会编译报错：./main.go:9:2: undefined: hello
	hello.Hello()
	// underline()

	// traversalString()
	// 若想做一个真正的随机数，要种子 seed()种子默认是1 rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个1~1000随机数
		b[i] = rand.Intn(1000)
	}
	fmt.Println(b)
	sum := sumArray(b)
	fmt.Printf("sum:%d\n", sum)

	n := []int{0, 1, 3, 5, 7, 8}
	fmt.Println(n)
	testsSum(n, 8)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[:2:3]
	fmt.Println(s, len(s), cap(s))

	s = append(s, 10, 20)        // 添加2个值，超出容量3限制
	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}
