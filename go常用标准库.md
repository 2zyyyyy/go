# go 常用标准库

### fmt

fmt包实现了雷类似语言print和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。

#### 向外输出

标准库fmt提供了以下几种输出相关函数。

****

#### Print

Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Print函数支持格式化输出字符串，Print函数会在输出内容的结尾添加一个换行符。

```go
func Print(a ...interface{}) (n int, err error)
func Print(format stringm a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

举个简单的例子：

```go
package main

import "fmt"

// Print
func printDemo() {
    fmt.Print("终端输出该文本信息:")
	name := "月满轩尼诗"
    fmt.Printf("我是:%s\n", name)
	fmt.Println("终端单独一行输出内容。")
}

func main() {
	printDemo()
}
```

输出：

```GO
终端输出该文本信息:我是月满轩尼诗
终端单独一行输出内容。       
```

#### Fprint

Fprint系列函数会将内容输出到一个io.Write接口类型的变量w中，我们通常用这个函数往文件中写入内容。

```GO
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) 
```

举个简单的例子：

```GO
// Fprint
func fprintDemo() {
	_, _ = fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./Fprint.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%s\n", err)
		return
	}
	name := "月满轩尼诗"
	// 向打开的文件句柄中写入内容
	_, _ = fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)
}
```

*注意：只要瞒住io.Write接口的类型都支持写入*

#### Sprint

Sprint系列函数会把传入的数据生成并返回一个字符串。

```GO
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```

举个简单的例子：

```go
// Sprint
func sprintDemo() {
	str1 := fmt.Sprint("月满轩尼诗1")
	name := "月满轩尼诗2"
	age := 18
	str2 := fmt.Sprintf("name:%s, age:%d", name, age)
	str3 := fmt.Sprint("月满轩尼诗3")
	fmt.Println(str1, str2, str3)
}
```

#### Errorf

Errorf函数根据format参数生成格式化字符串并返回包含该字符串的错误。

```GO
func Errorf(format string, a...interface{}) error
```

通常使用这种方式来自定义错误类型，例如：

```go
err := fmt.Errorf("这是一个错误")
```

#### 格式化占位符

`*printf`系列函数都支持format格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

#### 通用占位符

| 占位符 | 说明                                 |
| ------ | ------------------------------------ |
| %v     | 值的默认格式表示                     |
| %+v    | 类似%v，但是输出结构体时会添加字段名 |
| %#v    | 值的Go语法表示                       |
| %T     | 打印值的类型                         |
| %%     | 百分号                               |

示例代码如下：

```GO
// 格式化占位符
func formatDemo() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", true)
	person := struct {
		name string
		age  int32
	}{"月满轩尼诗", 18}
	fmt.Printf("%v\n", person)
	fmt.Printf("%#v\n", person)
	fmt.Printf("%T\n", person)
	fmt.Printf("100%%\n")
}
```

输出结果：

```go
100                                                         
true                                                        
{月满轩尼诗 18}                                             
struct { name string; age int32 }{name:"月满轩尼诗", age:18}
struct { name string; age int32 }                           
100%    
```

#### 布尔型

| 占位符 | 说明        |
| ------ | ----------- |
| %t     | true或false |

#### 整型

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %b     | 表示二进制                                                   |
| %c     | 该值对应的unicode码值                                        |
| %d     | 表示为十进制                                                 |
| %o     | 表示为八进制                                                 |
| %x     | 表示为十六进制，使用a-f                                      |
| %X     | 表示为十六进制，使用A-F                                      |
| %U     | 表示为Unicode格式：U+1234，等价于“U+%04X”                    |
| %q     | 改值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |

实例代码：

```GO
n := 65
fmt.Printf("%b\n", n)
fmt.Printf("%c\n", n)
fmt.Printf("%d\n", n)
fmt.Printf("%o\n", n)
fmt.Printf("%x\n", n)
fmt.Printf("%X\n", n) 
```

输出结果：

```go
    1000001
    A
    65
    101
    41
    41 
```

#### 浮点数与复数

| 占位符 | 说明                                                   |
| ------ | ------------------------------------------------------ |
| %b     | 无小数部分、二进制指数的科学计数法，如-123456p-78      |
| %e     | 科学计数法，如-1234.456e+78                            |
| %E     | 科学计数法，如-1234.456E+78                            |
| %f     | 有小数部分但无指数部分，如123.456                      |
| %F     | 等价于%f                                               |
| %g     | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出） |
| %G     | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出） |

示例代码如下：

```go
f := 12.34
fmt.Printf("%b\n", f)
fmt.Printf("%e\n", f)
fmt.Printf("%E\n", f)
fmt.Printf("%f\n", f)
fmt.Printf("%g\n", f)
fmt.Printf("%G\n", f) 
```

输出结果如下：

```go
    6946802425218990p-49
    1.234000e+01
    1.234000E+01
    12.340000
    12.34
    12.34 
```

#### 字符串和[]byte

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %s     | 直接输出字符串或者[]byte                                     |
| %q     | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %x     | 每个字节用两字符十六进制数表示（使用a-f                      |
| %X     | 每个字节用两字符十六进制数表示（使用A-F）                    |

示例代码：

```GO
// 字符串和[]byte
func byteDemo() {
	s := "月满轩尼诗"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
}
```

输出：

```GO
月满轩尼诗
"月满轩尼诗"
e69c88e6bba1e8bda9e5b0bce8af97
E69C88E6BBA1E8BDA9E5B0BCE8AF97
```

#### 指针

| 占位符 | 说明                         |
| ------ | ---------------------------- |
| %p     | 表示十六进制，并加上前导的0x |

示例代码：

```go
a := 18
fmt.Printf("%p\n", &a)
fmt.Printf("%#p\n", &a) 
```

输出：

```GO
    0xc000054058
    c000054058 
```

#### 宽度标识符

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：

| 占位符 | 说明               |
| :----- | :----------------- |
| %f     | 默认宽度，默认精度 |
| %9f    | 宽度9，默认精度    |
| %.2f   | 默认宽度，精度2    |
| %9.2f  | 宽度9，精度2       |
| %9.f   | 宽度9，精度0       |

示例代码如下：

```go
n := 88.88
fmt.Printf("%f\n", n)
fmt.Printf("%9f\n", n)
fmt.Printf("%.2f\n", n)
fmt.Printf("%9.2f\n", n)
fmt.Printf("%9.f\n", n)
```

输出结果如下：

```go
    88.880000
    88.880000
    88.88
        88.88
           89
```

#### 其他 flag

| 占位符 | 说明                                                         |
| :----- | :----------------------------------------------------------- |
| ’+’    | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）； |
| ’ ‘    | 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格 |
| ’-’    | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）； |
| ’#’    | 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值； |
| ‘0’    | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面； |

简单例子：

```GO
// 其他 flag
func otherFlagDemo() {
	s := "月满轩尼诗"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}
```

输出结果：

```go
测试
   测试
测试   
   测试
测试   
   测试
000测试
```

#### 获取输入

Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。

##### fmt.Scan

函数定签名如下：

```go
func Scan(a ...interface{}) (n int, err error)
```

- Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
- 本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因

简单例子：

```GO
// scan
func scanDemo() {
	var (
		name    string
		age     int
		married bool
	)
	_, _ = fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name：%s age：%d married：%t \n", name, age, married)
}
```

运行并输入：

```GO
$ go run main.go
万里 18 false
扫描结果 name：万里 age：18 married：false 
```

fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。

##### fmt.Scanf

函数签名如下：

```go
func Scanf(format string, a ...interface{}) (n int, err error)
```

- Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

代码示例如下：

```go
// scanf
func scanfdEmo() {
	var (
		name    string
		age     int
		married bool
	)
	_, _ = fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name：%s age：%d married：%t \n", name, age, married)
}
```

输入和输出：

```go
master ± go run main.go
1:wanli 2:18 3:true
扫描结果 name：wanli age：18 married：true 
```

fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。

例如，我们还是按照上个示例中以空格分隔的方式输入，fmt.Scanf就不能正确扫描到输入的数据。

##### fmt.Scanfln

函数签名如下：

```go
func Scanln(a ...interface{}) (n int, err error) 
```

- Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

具体代码示例如下：

```go
func scanlnDemo() {
	var (
		name    string
		age     int
		married bool
	)
	_, _ = fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name：%s age：%d married：%t \n", name, age, married)
}
```

将上面的代码编译后在终端执行，在终端依次输入1、2和true使用空格分隔。

```go
master ± go run main.go  
1  2  true
扫描结果 name：1 age：2 married：true 
```

fmt.Scanln遇到回车就结束扫描了，这个比较常用。

##### bufio.NewReader

有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。示例代码如下：

```go
func removeSpaces() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Printf("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
```

输入及输出结果：

```go
go run main.go
请输入内容：      1   2   3        
"1   2   3"
```

##### Sscan系列

这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
```



### Time

时间和日期是我们编程中经常会用到的，本文主要介绍了Go语言内置的time包的基本用法。

#### time 包

time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。

#### 时间类型

time.Time类型表示时间。我们可以通过time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。示例代码如下：

```GO
func timeNow() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```

输出结果：

```go
master ±✚ go run main.go
current time:2022-04-06 20:05:33.088168 +0800 CST m=+0.000073523
2022-4-6 20:05:33
```

#### 时间戳

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。

基于时间对象获取时间戳的示例代码如下：

```GO
// 时间戳
func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}
```

输出：

```GO
$ go run main.go
current timestamp1:1649247038
current timestamp2:1649247038178534000
```

使用time.Unix()函数可以将时间戳转为时间格式。

```go
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```

输出结果：

```go
master ±✚ go run main.go
2022-04-06 20:10:38 +0800 CST
2022-04-06 20:10:38
```

#### 时间间隔

time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。time.Duration表示一段时间间隔，可表示的最长时间段大约290年。

time包中定义的时间间隔类型的常量如下：

```go
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

例如：time.Duration表示1纳秒，time.Second表示1秒。

#### 时间操作

****

##### Add

我们在日常的编码过程中可能会遇到要求时间+时间间隔的需求，Go语言的时间对象有提供Add方法如下：

```go
    func (t Time) Add(d Duration) Time
```

举个例子，求2个小时之后的时间：

```go
// 求2个小时之后的时间：
func timeDifference() {
	now := time.Now()
	later := now.Add(time.Hour * 2) // 当前时间加 2 个小时后的时间
	fmt.Println(later)
}
```

##### Sub

求两个时间之间的差值：

```
    func (t Time) Sub(u Time) Duration 
```

返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。

##### Equal

```
func (t Time) Equal(u Time) bool
```

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

##### Before

```
func (t Time) Before(u Time) bool
```

如果t代表的时间点在u之前，返回真；否则返回假。

##### After

```
func (t Time) After(u Time) bool 
```

如果t代表的时间点在u之后，返回真；否则返回假。

#### 定时器

使用 time.Tick（时间间隔）来设置定时器，定时器的本质是一个通道（channel）。

```GO
// 定时器
func timer() {
	ticker := time.Tick(time.Second) // 定义一个间隔 1s 的定时器
	for j := range ticker {
		fmt.Println(j) // 每秒都会执行的任务
	}
}
```

输出结果：

```GO
2022-04-06 20:49:18.570535 +0800 CST m=+2.000806572
2022-04-06 20:49:19.570315 +0800 CST m=+3.000556264
2022-04-06 20:49:20.57078 +0800 CST m=+4.000991453
2022-04-06 20:49:21.570969 +0800 CST m=+5.001149924
2022-04-06 20:49:22.570875 +0800 CST m=+6.001025610
2022-04-06 20:49:23.570413 +0800 CST m=+7.000534147
2022-04-06 20:49:24.570922 +0800 CST m=+8.001012810
……
```

#### 时间格式化

时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧。

补充：如果想格式化为12小时方式，需指定PM。

```GO
func formatTime() {
	now := time.Now()
	// 格式化的模板为 Go 的出生时间（2006 1/2 15:04）
	// 24 小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12 小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
```

输出结果：

```GO
 master ±✚ go run main.go 
2022-04-06 21:00:56.594 Wed Apr
2022-04-06 09:00:56.594 Wed Apr
2022/04/06 21:00
21:00 2022/04/06
2022/04/06
```































