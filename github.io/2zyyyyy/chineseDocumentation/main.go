package main

import (
	"fmt"
	"math"
	"os"
	//_ "go/github.io/2zyyyyy/hello"
)

// 类型定义
type NewInt int

type MyString string

// 类型别名
type MyInt = int

// struct
type Cat struct {
	breed   string
	name    string
	age     int8
	Address Address
}

// 结构体嵌套
type Address struct {
	Province, City, County, Time string
}

// 嵌套匿名结构体
type User struct {
	name    string
	age     int
	Address // 匿名结构体字段 只有类型没有字段名
	Email
}

type Email struct {
	Account string
	Time    string
}

type test struct {
	a, b, c, d int8
}

// 结构体继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会移动！\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d Dog) run() {
	fmt.Printf("%s会跑！\n", d.name)
}

//结构体与JSON序列化
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

type Class struct {
	Title    string
	Students []*Student
}

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

// 构造函数
func newCat(breed, name string, age int8) *Cat {
	return &Cat{
		breed: breed,
		name:  name,
		age:   age,
	}
}

// 方法和接收者
func (c Cat) Eat() {
	fmt.Printf("%s每天就是吃吃吃~\n", c.name)
}

// 指针接收者
func (c *Cat) setAge(newAge int8) {
	fmt.Printf("我是修改年龄的方法，将%d修改为%d\n", c.age, newAge)
	c.age = newAge
}

// 值接收者
func (c Cat) setName(newName string) {
	fmt.Printf("我是修改name的方法，将%s修改为%s\n", c.name, newName)
	c.name = newName
}

// 自定义类型MyString的方法
func (m MyString) OutPut() {
	fmt.Println("Hello, 我是一个string。")
}

func length(s string) int {
	fmt.Println("call length")
	return len(s)
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
	// var num int
	// fmt.Println(&num)
	// ptr := &num
	// fmt.Println(reflect.TypeOf(ptr))
	// *ptr = 20
	// fmt.Println(num)

	// scoreMap := make(map[string]int, 8)
	// scoreMap["张三"] = 90
	// scoreMap["李四"] = 95
	// fmt.Println(scoreMap)
	// fmt.Println(scoreMap["李四"])
	// fmt.Printf("type of:%T\n", scoreMap)

	// 如果key存在ok为true，v为对应的值；不存在OK=false v为值类型的零值
	// v, ok := scoreMap["张三1"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("查无此人~")
	// }

	// scoreMap["王五"] = 100
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }

	// userInfo := map[string]string{
	// 	"userName": "wanli",
	// 	"passWord": "123456",
	// }
	// fmt.Println(userInfo)

	// 按照指定顺序遍历map
	// rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	// scoreMap := make(map[string]int, 200)

	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("test%2d", i) // 生产test开头的字符串
	// 	value := rand.Intn(100)          // 生产0~99随机整数
	// 	scoreMap[key] = value
	// }

	// // 取出map中所有的key存入切片keys
	// keys := make([]string, 0, 200)
	// for key := range scoreMap {
	// 	keys = append(keys, key)
	// }
	// // 对切片排序
	// sort.Strings(keys)
	// // 按照排序后的key遍历map
	// for _, key := range keys {
	// 	fmt.Println(key, scoreMap[key])
	// }

	// 元素为map类型的切片
	// mapSlice := make([]map[string]string, 3)
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }
	// fmt.Println("after init")
	// // 对切片中的map元素进行初始化
	// mapSlice[0] = make(map[string]string, 10)
	// mapSlice[0]["name"] = "张三"
	// mapSlice[0]["passWord"] = "123456"
	// mapSlice[0]["address"] = "未来park"
	// mapSlice[0]["age"] = "10"
	// mapSlice[0]["sex"] = "男"
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d, value:%v\n", index, value)
	// }
	// fmt.Printf("------%T\n", mapSlice)

	// 值为切片类型的map
	// sliceMap := make(map[string][]string, 3)
	// fmt.Println(sliceMap)
	// fmt.Println("after init~~~")
	// key := "杭州"
	// value, ok := sliceMap[key]
	// if !ok {
	// 	value = make([]string, 0, 2)
	// }
	// value = append(value, "北京", "上海")
	// sliceMap[key] = value
	// fmt.Println(sliceMap)
	// fmt.Printf("sliceMap Type is:%T, sliceMap[杭州] Type is:%T\n", sliceMap, sliceMap[key])

	// 类型定义和类型别名区别示例
	// var a NewInt
	// var b MyInt

	// fmt.Printf("type of a:%T\n", a) // type of a:main.NewInt
	// fmt.Printf("type of b:%T\n", b) // type of b:int

	// struct
	// cats := Cat{}
	// cats.name = "二狗子"
	// cats.breed = "加菲"
	// cats.age = 3
	// fmt.Printf("cats=%v\n", cats)  // cats={加菲 二狗子 3}
	// fmt.Printf("cats=%#v\n", cats) // cats=main.Cat{breed:"加菲", name:"二狗子", age:3}

	// cats2 := new(Cat)
	// fmt.Printf("cats2 type=%T\n", cats2) // cats2 type=*main.Cat
	// fmt.Printf("cats2=%#v\n", cats2)     // cats2=&main.Cat{breed:"", name:"", age:0}
	// cats2.age = 100
	// cats2.name = "西西"                // cats2=&main.Cat{breed:"", name:"", age:0}
	// fmt.Printf("cats2:%#v\n", cats2) // cats2:&main.Cat{breed:"", name:"西西", age:100}

	// cats := &Cat{}
	// fmt.Printf("%T\n", cats)
	// fmt.Printf("cats:%v\n", cats)
	// cats.breed = "中华田园猫"
	// cats.age = 10
	// cats.name = "技艺"
	// fmt.Printf("cats:%#v\n", cats)

	// cats := &Cat{
	// 	breed: "美短",
	// }
	// fmt.Printf("cats:%T\n", cats)

	// n := test{
	// 	1, 2, 3, 4,
	// }
	// fmt.Printf("n.a %p\n", &n.a)
	// fmt.Printf("n.b %p\n", &n.b)
	// fmt.Printf("n.c %p\n", &n.c)
	// fmt.Printf("n.d %p\n", &n.d)

	// 调用构造函数newCat()
	// cats := newCat("加菲", "西西", 3)
	// fmt.Printf("%#v\n", cats)

	// // 方法和接收者
	// cats.Eat()
	// cats.setAge(127)
	// fmt.Println(cats.age)
	// cats.setName("咚咚咚")
	// fmt.Println(cats.name)

	// var str MyString
	// str.OutPut()
	// str = "test"
	// fmt.Printf("%#v  %T\n", str, str) // "test"  main.MyString

	// cats := Cat{
	// 	"加菲",
	// 	"西西",
	// 	3,
	// 	Address{
	// 		"浙江省",
	// 		"杭州市",
	// 		"拱墅区",
	// 	},
	// }
	// fmt.Printf("%#v\n", cats)

	// var user User
	// user.age = 10
	// user.name = "嵌套匿名结构体"
	// user.Address.Province = "浙江省" //通过匿名结构体.字段名访问
	// user.City = "杭州市"             // 直接访问匿名结构体的字段名
	// fmt.Printf("%#v\n", user)

	// 嵌套结构体字段冲突
	// var user User
	// user.name = "字段冲突"
	// user.age = 15
	// // 指定结构体中的字段给与赋值
	// user.Address.Time = "address.time"
	// user.Email.Time = "email.time"
	// fmt.Printf("%#v\n", user)

	// dog := &Dog{
	// 	4,
	// 	&Animal{
	// 		"嘻嘻",
	// 	},
	// }
	// dog.move()
	// dog.run()
	// fmt.Printf("%#v\n", dog)

	// 结构体与JSON序列化
	// class := &Class{
	// 	Title:    "中队长",
	// 	Students: make([]*Student, 0, 200),
	// }
	// for i := 0; i < 10; i++ {
	// 	stu := &Student{
	// 		Name:   fmt.Sprintf("stu%02d", i),
	// 		Gender: "男",
	// 		ID:     i,
	// 	}
	// 	class.Students = append(class.Students, stu)
	// }
	// // JSON序列化：结构体——>JSON格式字符串
	// data, err := json.Marshal(class)
	// if err != nil {
	// 	fmt.Printf("json marshal failed!%s\n", err)
	// 	return
	// }
	// fmt.Printf("json:%s\n", data)

	// // JSON反序列化：JSON格式字符串-->结构体
	// str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	// class1 := &Class{}
	// err = json.Unmarshal([]byte(str), &class1)
	// if err != nil {
	// 	fmt.Printf("json unmarshal failed! %s\n", err)
	// 	return
	// }
	// fmt.Printf("%#v\n", class1)

	// 结构体标签
	// student := &Student{
	// 	ID:     1001,
	// 	Gender: "女",
	// 	Name:   "小丑杰克",
	// }
	// data, err := json.Marshal(student)
	// if err != nil {
	// 	fmt.Printf("json marshal failed, err:%s\n", err)
	// }
	// fmt.Printf("json:%s\n", data)

	// 删除map类型结构体
	// animals := make(map[int]Animal)
	// animals[0] = Animal{"花花"}
	// animals[1] = Animal{"西西"}
	// fmt.Println(animals)

	// delete(animals, 0)
	// fmt.Println(animals)

	// 实现map有序输出
	// mapSort := make(map[int]int)
	// mapSort[10] = 128
	// mapSort[8] = 256
	// mapSort[2] = 64
	// mapSort[9] = 100
	// fmt.Println(mapSort)

	// sl := []int{}
	// for k := range mapSort {
	// 	fmt.Println(k)
	// 	sl = append(sl, k)
	// }
	// sort.Ints(sl)
	// fmt.Println(sl)
	// for i := 0; i < len(mapSort); i++ {
	// 	fmt.Printf("key:%d, value:%d\n", sl[i], mapSort[sl[i]])
	// }

	// 定义局部变量
	// a := 100
	// if a < 20 {
	// 	// 如果条件为true
	// 	fmt.Printf("a小于20\n")
	// }
	// fmt.Printf("a的值为：%d\n", a)

	// // switch
	// grade := "B"
	// // marks := 90
	// var marks int

	// switch marks {
	// case 90:
	// 	grade = "A"
	// case 80:
	// 	grade = "B"
	// case 60, 70:
	// 	grade = "C"
	// case 50:
	// 	grade = "E"
	// default:
	// 	grade = "D"
	// }
	// fmt.Println(grade, marks)

	// switch {
	// case grade == "A":
	// 	fmt.Printf("优秀：%s\n", grade)
	// case grade == "B", grade == "C":
	// 	fmt.Printf("良好：%s\n", grade)
	// case grade == "D":
	// 	fmt.Printf("及格：%s\n", grade)
	// case grade == "E":
	// 	fmt.Printf("不及格：%s\n", grade)
	// default:
	// 	fmt.Printf("及格：%s\n", grade)
	// }

	// type switch
	// var x interface{}
	// // 写法1
	// switch i := x.(type) {
	// case nil:
	// 	fmt.Printf("x的类型为:%T\n", i)
	// case int:
	// 	fmt.Println("x是int类型")
	// case float64:
	// 	fmt.Println("x是float64类型")
	// case func(int):
	// 	fmt.Println("x是fun(int)类型")
	// case bool, string:
	// 	fmt.Println("x是bool或string类型")
	// default:
	// 	fmt.Println("??未知类型")
	// }

	// // 写法2
	// j := 0
	// switch j {
	// case 0:
	// case 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("default")
	// }

	// // 写法3
	// k := 0
	// switch k {
	// case 0:
	// 	println("fallthrough")
	// 	fallthrough
	// 	/*
	// 	   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
	// 	   而如果switch没有表达式，它会匹配true。
	// 	   Go里面switch默认相当于每个case最后带有break，
	// 	   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
	// 	   但是可以使用fallthrough强制执行后面的case代码。
	// 	*/
	// case 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("default")
	// }

	// //写法三
	// var m = 0
	// switch m {
	// case 0, 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("default")
	// }

	// //写法四
	// var n = 0
	// switch { //省略条件表达式，可当 if...else if...else
	// case n > 0 && n < 10:
	// 	fmt.Println("i > 0 and i < 10")
	// case n > 10 && n < 20:
	// 	fmt.Println("i > 10 and i < 20")
	// default:
	// 	fmt.Println("default")
	// }

	// select
	// var c1, c2, c3 chan int
	// var i1, i2 int
	// fmt.Printf("c1:%v, c2:%v, c3:%v\n", c1, c2, c3)
	// fmt.Printf("i1:%d, i2:%d\n", i1, i2)
	// select {
	// case i1 = <-c1:
	// 	fmt.Printf("received %d from c1\n", i1)
	// case c2 <- i2:
	// 	fmt.Printf("sent %d to c2\n", i2)
	// case i3, ok := <-c3:
	// 	fmt.Printf("i3:%d\n", i3)
	// 	if ok {
	// 		fmt.Printf("received %v from c3\n", i3)
	// 	} else {
	// 		fmt.Printf("c3 is closed\n")
	// 	}
	// default:
	// 	fmt.Printf("no communivation\n")
	// }

	// s := "abc"
	// println(s[0])
	// for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
	// 	println(s[i])
	// }

	// n := len(s)
	// for n > 0 { // 替代 while (n > 0) {}
	// 	n--
	// 	println(s[n]) // 替代 for (; n > 0;) {}
	// }

	// for { // 替代 while (true) {}
	// 	println(s) // 替代 for (;;) {}
	// }

	// str := "abcd"
	// for i, n := 0, length(str); i < n; i++ { // 避免多次调用length函数
	// 	fmt.Println(i, str[i])
	// }
	// a := 0
	// count := 0
	// b := 15

	// numbers := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(len(numbers))

	// // for 1
	// for a := 0; a < 10; a++ {
	// 	fmt.Printf("a=:%d\n", a)
	// 	count++
	// }
	// fmt.Printf("for循环1执行了:%d次\n", count)

	// count = 0
	// // for 2
	// for a < b {
	// 	a++
	// 	fmt.Printf("a的值为：%d\n", a)
	// 	count++
	// }
	// fmt.Printf("for循环2执行了:%d次\n", count)

	// // for 3
	// for i, x := range numbers {
	// 	i++
	// 	fmt.Printf("第%d位x的值为%d\n", i, x)
	// }

	// 双层循环筛选素数
	// var i, j int
	// for i = 2; i < 10; i++ {
	// 	fmt.Printf("i=%d\n", i)
	// 	for j = 2; j <= (i / j); j++ {
	// 		// fmt.Println(j)
	// 		if i%j == 0 {
	// 			break // 如果发现因子，则不是素数
	// 		}
	// 	}
	// 	if j > (i / j) {
	// 		fmt.Printf("%d  是素数\n", i)
	// 	}
	// }

	// range会复制对象
	// a := [3]int{1, 2, 3}

	// for i, v := range a { // i,v都是从复制品中取出
	// 	if i == 1 {
	// 		// 在修改前我们先修改原数组
	// 		a[1], a[2] = 900, 1000
	// 		fmt.Println(a) // 确认修改是有效的, 输出[1, 900, 1000]
	// 	}
	// 	a[i] = v + 100
	// }
	// fmt.Println(a) // 输出[101, 102, 103]

	// 改用引用类型，其底层数据不会被复制
	// s := []int{4, 5, 6, 7, 8}

	// for i, v := range s { // 复制struct slice（pointer，len，cap）
	// 	if i == 0 {
	// 		fmt.Println(s)
	// 		s = s[:3]    // 对slice的修改 不会影响range
	// 		s[2] = 10086 // 对底层数据的修改
	// 		fmt.Printf("在range中修改s, s:%v\n", s)
	// 	}
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("s:%v\n", s)
}
