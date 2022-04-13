# Go 常用标准库

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

##### 解析字符串格式时间

```GO
func parseStringTime() {
	now := time.Now()
	fmt.Println("当前时间：", now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/04/07 10:46:47", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("指定字符串的时间：", timeObj)
  fmt.Println("指定时间与当前时间差：", timeObj.Sub(now))
}
```

输出：

```GO
master ± go run main.go 
当前时间： 2022-04-07 10:51:54.699272 +0800 CST m=+0.000080455
指定字符串的时间： 2022-04-07 10:46:47 +0800 CST
指定时间与当前时间差： -5m7.699272s
```



### Log

Go语言内置的log包实现了简单的日志服务。本文介绍了标准库log的基本使用。

#### 使用Logger

log包定义了Logger类型，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的“标准”logger，可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用。

例如，我们可以像下面的代码一样直接通过log包来调用上面提到的方法，默认它们会将日志信息打印到终端界面：

```GO
func logDemo() {
	log.Println("普通的 log~~")
	msg := "普通的"
	log.Printf("这是一条%s 的日志", msg)
	log.Fatalln("这是一条会触发 fatal 的日志")
	log.Panicln("这是一条会触发 panic 的日志")
}
```

输出：

```GO
master ±✚ go run .
2022/04/07 11:03:33 普通的 log~~
2022/04/07 11:03:33 这是一条普通的 的日志
2022/04/07 11:03:33 这是一条会触发 fatal 的日志
exit status 1
```

logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。

#### 配置 Logger

默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。log标准库中为我们提供了定制这些设置的方法。

log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。

```go
    func Flags() int
    func SetFlags(flag int) 
```

##### flag选项

log标准库提供了如下的flag选项，它们是一系列定义好的常量。

```go
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
) 
```

下面我们在记录日志之前先设置一下标准logger的输出选项如下：

```GO
func logFlagSet() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条普通的日志")
}
```

输出：

```GO
 master ±✚ go run .
2022/04/07 16:03:33.998119 /Users/gilbert/go/src/go/github.io/2zyyyyy/standardLibrary/log/main.go:16: 这是一条普通的日志
```

#### 配置日志前缀

log标准库中还提供了关于日志信息前缀的两个方法：

```
    func Prefix() string
    func SetPrefix(prefix string) 
```

其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。

```go
// 配置日志前缀
func setPrefixDemo() {
   log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
   log.Println("这是一条普通的日志")
   // 设置输出前缀
   log.SetPrefix("[测试日志前缀]")
   log.Println("这是一条普通的日志")
}
```

输出：

```GO
master ±✚ go run .
2022/04/07 16:09:14.843830 /Users/gilbert/go/src/go/github.io/2zyyyyy/standardLibrary/log/main.go:22: 这是一条普通的日志
[测试日志前缀]2022/04/07 16:09:14.843939 /Users/gilbert/go/src/go/github.io/2zyyyyy/standardLibrary/log/main.go:25: 这是一条普通的日志
```

这样我们就能够在代码中为我们的日志信息添加指定的前缀，方便之后对日志信息进行检索和处理。

#### 配置日志输出位置

```
func SetOutput(w io.Writer)
```

SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。

例如，下面的代码会把日志输出到同目录下的mylog.log文件中:

```GO
func logOutput() {
	logFile, err := os.OpenFile("./mylog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[测试日志前缀]")
	log.Println("这是一条很普通的日志。")
}
```

查看指定的日志文件内容：

```bash
master ±✚ cat ./mylog.log    
2022/04/07 17:13:59.799010 /Users/gilbert/go/src/go/github.io/2zyyyyy/standardLibrary/log/main.go:17: 这是一条很普通的日志。
[测试日志前缀]2022/04/07 17:13:59.799157 /Users/gilbert/go/src/go/github.io/2zyyyyy/standardLibrary/log/main.go:19: 这是一条很普通的日志。
```

如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中。

```go
func init() {
    logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("open log file failed, err:", err)
        return
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
```

#### 创建logger

log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下：

```
    func New(out io.Writer, prefix string, flag int) *Logger
```

New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。

举个例子：

```go
func main() {
    logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
    logger.Println("这是自定义的logger记录的日志。")
}
```

将上面的代码编译执行之后，得到结果如下：

```
    <New>2019/10/11 14:06:51 main.go:34: 这是自定义的logger记录的日志。
```

总结 :
Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等。

### IO 操作

#### 输入输出底层原理

终端其实是一个文件，相关实例如下：

- `os.Stdin`：标准输入的文件实例，类型为`*File`
- `os.Stdout`：标准输出的文件实例，类型为`*File`
- `os.Stderr`：标准错误输出的文件实例，类型为`*File`

以文件的方式操作终端:

```go
package main

import "os"

func main() {
    var buf [16]byte
    os.Stdin.Read(buf[:])
    os.Stdin.WriteString(string(buf[:]))
}
```

#### 文件操作相关 API

- ```
  func Create(name string) (file *File, err Error)
  ```

  - 根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666

- ```
  func NewFile(fd uintptr, name string) *File
  ```

  - 根据文件描述符创建相应的文件，返回一个文件对象

- ```
  func Open(name string) (file *File, err Error)
  ```

  - 只读方式打开一个名称为name的文件

- ```
  func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
  ```

  - 打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限

- ```
  func (file *File) Write(b []byte) (n int, err Error)
  ```

  - 写入byte类型的信息到文件

- ```
  func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
  ```

  - 在指定位置开始写入byte类型的信息

- ```
  func (file *File) WriteString(s string) (ret int, err Error)
  ```

  - 写入string信息到文件

- ```
  func (file *File) Read(b []byte) (n int, err Error)
  ```

  - 读取数据到b中

- ```
  func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
  ```

  - 从off开始读取数据到b中

- ```
  func Remove(name string) Error
  ```

  - 删除文件名为name的文件

  #### 打开和关闭文件

  `os.Open()`函数能够打开一个文件，返回一个`*File`和一个`err`。对得到的文件实例调用close()方法能够关闭文件。

  ```GO
  // 打开和关闭文件
  func openFile() {
  	// 只读方式打开当前目录下的 main.go 文件
  	file, err := os.Open("./main.go")
  	defer func(file *os.File) {
  		err := file.Close()
  		if err != nil {
  			fmt.Printf("close file failed, err%s\n", err)
  			return
  		}
  	}(file)
  	if err != nil {
  		fmt.Println("open file failed! err:", err)
  		return
  	}
  	log.Println("文件打开成功~")
  }
  ```

  #### 写文件

  ```GO
  func writeFile() {
  	// 新建文件
  	file, err := os.Create("./create.txt")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}
  	defer func(file *os.File) {
  		err := file.Close()
  		if err != nil {
  			fmt.Println(err)
  			return
  		}
  	}(file)
  	for i := 0; i < 5; i++ {
  		_, _ = file.WriteString("ab\n")
  		_, _ = file.Write([]byte("cd\n"))
  	}
  }
  ```

  #### 读文件

  文件读取可以用file.Read()和file.ReadAt()，读到文件末尾会返回io.EOF的错误

  ```GO
  func readFile() {
  	// 打开文件
  	file, err := os.Open("./create.txt")
  	if err != nil {
  		fmt.Println("open file failed, err:", err)
  		return
  	}
  	defer func(file *os.File) {
  		err := file.Close()
  		if err != nil {
  			fmt.Println(err)
  		}
  	}(file)
  	// 定义接收文件读取的字节数组
  	var buf [128]byte
  	var content []byte
  	for {
  		n, err := file.Read(buf[:])
  		if err == io.EOF {
  			// 读取结束
  			break
  		}
  		if err != nil {
  			fmt.Println("read file err", err)
  			return
  		}
  		content = append(content, buf[:n]...)
  	}
  	fmt.Print(string(content))
  }
  ```

  #### 拷贝文件

  ```go
  // 拷贝文件
  func copyFile() {
  	// 打开文件
  	srcFile, err := os.Open("./create.txt")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}
  	// 创建新文件
  	newFile, err2 := os.Create("./copy.txt")
  	if err2 != nil {
  		fmt.Println(err2)
  		return
  	}
  	
  	// defer 关闭文件
  	defer func(srcFile *os.File) {
  		err := srcFile.Close()
  		if err != nil {
  			fmt.Println(err)
  		}
  	}(srcFile)
  	defer func(newFile *os.File) {
  		err := newFile.Close()
  		if err != nil {
  			fmt.Println(err)
  		}
  	}(newFile)
  
  	// 缓存读取
  	buf := make([]byte, 1024)
  	for {
  		// 从源文件读数据
  		n, err := srcFile.Read(buf)
  		if err == io.EOF {
  			fmt.Println("读取完毕~")
  			break
  		}
  		// 写进去
  		_, err = newFile.Write(buf[:n])
  		if err != nil {
  			return
  		}
  	}
  }
  ```
  
  ##### bufio
  
  - bufio包实现了带缓冲区的读写，是对文件读写的封装
  - bufio缓冲写数据
  
  | 模式        | 含义     |
  | :---------- | :------- |
  | os.O_WRONLY | 只写     |
  | os.O_CREATE | 创建文件 |
  | os.O_RDONLY | 只读     |
  | os.O_RDWR   | 读写     |
  | os.O_TRUNC  | 清空     |
  | os.O_APPEND | 追加     |

- bufio读数据

  ```GO
  // bufIo
  func writeFile() {
  	// w(写) 2 r(读) 4 x(执行) 1
  	file, err := os.OpenFile("./bufIo.txt", os.O_CREATE|os.O_WRONLY, 0666)
  	if err != nil {
  		fmt.Print("open file failed, err:", err)
  		return
  	}
  	defer func(file *os.File) {
  		_ = file.Close()
  	}(file)
  	// 获取write对象
  	write := bufio.NewWriter(file)
  	for i := 0; i < 10; i++ {
  		_, err = write.WriteString("月满轩尼诗\n")
  		if err != nil {
  			return
  		}
  	}
  	// 刷新缓冲区，强制写出
  	err = write.Flush()
  	if err != nil {
  		return
  	}
  }
  
  func readFile() {
  	file, err := os.Open("./bufIo.txt")
  	if err != nil {
  		fmt.Println("open file failed, err:", err)
  		return
  	}
  	defer func(file *os.File) {
  		_ = file.Close()
  	}(file)
  	reader := bufio.NewReader(file)
  	for {
  		line, _, err := reader.ReadLine()
  		if err == io.EOF {
  			break
  		}
  		if err != nil {
  			return
  		}
  		fmt.Println(string(line))
  	}
  }
  
  func main() {
  	writeFile()
  	readFile()
  }
  ```
  
  #### ioutil工具包
  
  - 工具包写文件
  - 工具包读取文件
  
  ```go
  func writeFile() {
  	err := ioutil.WriteFile("./ioUtil.txt", []byte("月满轩尼诗"), 0666)
	if err != nil {
  		fmt.Println("ioUtil write file failed, err:", err)
		return
  	}
  }
  
  func readFile() {
  	content, err := ioutil.ReadFile("./ioUtil.txt")
  	if err != nil {
  		fmt.Println("ioUtil read file failed, err:", err)
  		return
  	}
  	fmt.Println(string(content))
  }
  
  func main() {
  	writeFile()
  	readFile()
  }
  ```

#### 实例

##### 实现一个cat命令

使用文件操作相关知识，模拟实现Linux的cat命令功能

```GO
// 使用文件操作相关知识，模拟实现linux平台cat命令的功能

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') // 注意是字符
		if err == io.EOF {
			break
		}
		_, err = fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
```

### Strconv

strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。

更多函数请查看[官方文档](https://golang.org/pkg/strconv/)。

#### string与int类型转换

这一组函数是我们平时编程中用的最多的函数。

#### Atoi()

Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下：

```go
func Atoi(s string) (i int, err error)
```

如果传入的字符串参数无法转换为int类型，就会返回错误。

```go
// Atoi()如果传入的字符串参数无法转换为int类型，就会返回错误。
func atoiDemo() {
	str := "100"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%s can't convert to int\n", str)
	} else {
		fmt.Printf("type：%T\nvalue：%#v\n", i, i) // type: int  value: 100
	}
}
```

输出：

```GO
$go run .
type：int value：100
```

如果我们修改str的值，使得无法转换为int值，再次执行就会返回对应的错误：

```GO
str := "100测试"

// 输出
$ go run .
100测试 can't convert to int
```

#### Itoa()

Itoa()函数用于将int类型数据转换为对应的字符串表示，具体的函数签名如下。

```go
func Itoa(i int) string 
```

示例代码：

```GO
// Itoa() 数字转字符串
func itoaDemo() {
	i := 200
	str := strconv.Itoa(i)
	fmt.Printf("type:%T\nvalue:%#v\n", str, str) // type:string  value:"200"
}

// 输出
$ go run .
type:string
value:"200"
```

#### Parse系列函数

Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。

##### ParseBool()

```go
func ParseBool(str string) (value bool, err error)
```

返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。

##### ParseInt()

```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

返回字符串表示的整数值，接受正负号。

base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；

bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；

返回的err是`*NumErr`类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

##### ParseUnit()

```go
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
```

ParseUint类似ParseInt但不接受正负号，用于无符号整型。

##### ParseFloat()

```go
func ParseFloat(s string, bitSize int) (f float64, err error)
```

解析一个表示浮点数的字符串并返回其值。

如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。

bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；

返回值err是`*NumErr`类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。

#### Parse系列函数代码示例

```go
// Parse系列函数
func parseFunc() {
	// ParseBool
	b, _ := strconv.ParseBool("true")
	fmt.Println("parseBool:", b)
	// ParseFloat
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println("parseFloat:", f)
	// ParseInt
	i, _ := strconv.ParseInt("-2", 10, 64)
	fmt.Println("parseInt:", i)
	// ParseUnit
	u, _ := strconv.ParseUint("-2", 10, 64)
	fmt.Println("parseUint:", u)
}

// 输出
$ go run .
parseBool: true   
parseFloat: 3.1415
parseInt: -2      
parseUint: 0    
```

#### Format系列函数

Format系列函数实现了将给定类型数据格式化为string类型数据的功能。

##### FormatBool()

```go
func FormatBool(b bool) string
```

根据b的值返回”true”或”false”。

##### FormatInt()

```go
func FormatInt(i int64, base int) string
```

返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。

##### FormatUint()

```go
func FormatUint(i uint64, base int) string
```

是FormatInt的无符号整数版本。

##### FormatFloat()

```go
func FormatFloat(f float64, fmt byte, prec, bitSize int) string 
```

函数将浮点数表示为字符串并返回。

bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。

fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。

prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。

#### 代码示例：

```GO
// Format系列函数
func formatFunc() {
	// formatBool
	s1 := strconv.FormatBool(true)
	fmt.Println("formatBool:", s1) // formatBool: true
	// formatBool
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println("formatFloat:", s2) // formatFloat: 3.1415E+00
	// formatBool
	s3 := strconv.FormatInt(-2, 16)
	fmt.Println("formatInt:", s3) // formatInt: -2
	// formatBool
	s4 := strconv.FormatUint(2, 16)
	fmt.Println("formatUint:", s4) // formatUint: 2
}

// 输出 
$ go run .
formatBool: true
formatFloat: 3.1415E+00
formatInt: -2          
formatBoos: 2          
```

#### 其他

##### isPrint()

```go
func IsPrint(r rune) bool 
```

返回一个字符是否是可打印的，和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格。

##### CanBackquote()

```go
func CanBackquote(s string) bool 
```

返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。

##### 其他

除上文列出的函数外，strconv包中还有Append系列、Quote系列等函数。具体用法可查看[官方文档](https://golang.org/pkg/strconv/)。

### Template

html/template包实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。它提供了和text/template包相同的接口，Go语言中输出HTML的场景都应使用text/template包。

****

#### 模板

在基于MVC的Web架构中，我们通常需要在后端渲染一些数据到HTML文件中，从而实现动态的网页效果。

#### 模板示例

通过将模板应用于一个数据结构（即该数据结构作为模板的参数）来执行，来获得输出。模板中的注释引用数据接口的元素（一般如结构体的字段或者字典的键）来控制执行过程和获取需要呈现的值。模板执行时会遍历结构并将指针表示为’.‘（称之为”dot”）指向运行过程中数据结构的当前位置的值。

用作模板的输入文本必须是utf-8编码的文本。”Action”—数据运算和控制单位—由”{{“和”}}“界定；在Action之外的所有文本都不做修改的拷贝到输出中。Action内部不能有换行，但注释可以有换行。

HTML文件代码如下：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hello</title>
</head>
<body>
    <p>Hello {{.}}</p>
</body>
</html>
```

我们的HTTP server端代码如下：

```go
// HTTP server端代码
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmp, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入 w
	_ = tmp.Execute(w, "www.google.com")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}
}
```

#### 模板语法

`{{.}}`

模板语法都包含在`{{和}}`中间，其中`{{.}}`中的点表示当前对象。

当我们传入一个结构体对象时，我们可以根据`.`来访问结构体的对应字段。

略过。

### Http

Go语言内置的net/http包十分的优秀，提供了HTTP客户端和服务端的实现。

****

#### net/httpnet/http介绍

Go语言内置的net/http包提供了HTTP客户端和服务端的实现。

#### HTTP 协议

超文本传输协议（HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络传输协议，所有的WWW文件都必须遵守这个标准。设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。

#### HTTP 客户端

基本的HTTP/HTTPS请求
Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。

```go
resp, err := http.Get("http://www.google.com/")
...
resp, err := http.Post("http://www.google.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://www.google.com/form",
    url.Values{"key": {"Value"}, "id": {"123"}}) 
```

程序在使用完response后必须关闭回复的主体。

```go
resp, err := http.Get("http://www.google.com/")
if err != nil {
    // handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

#### Get 请求示例

使用net/http包编写一个简单的发送HTTP请求的Client端，代码如下：

```go
// net/http get demo
func main() {
	res, err := http.Get("https://2zyyyyy.github.io/")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("defer res.Body close failed, err:", err)
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed, err:", err)
		return
	}
	fmt.Println(string(body))
}
```

将上面的代码保存之后编译成可执行文件，执行之后就能在终端打印网站首页的内容了，我们的浏览器其实就是一个发送和接收HTTP协议数据的客户端，我们平时通过浏览器访问网页其实就是从网站的服务器接收HTTP数据，然后浏览器会按照HTML、CSS等规则将网页渲染展示出来。

#### 带参数的 Get 请求

关于GET请求的参数需要使用Go语言内置的net/url这个标准库来处理。

```go
// 关于GET请求的参数需要使用Go语言内置的net/url这个标准库来处理。
func main() {
	apiUrl := "http://127.0.0.1:9090/get"
	// url param
	data := url.Values{}
	data.Set("name", "月满轩尼诗")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // url encode
	fmt.Println(u.String())
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("res body close failed, err:", err)
		}
	}(res.Body)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("readAll failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
```

对应的Server端HandlerFunc如下：

```go
// server 端
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("r.body.close failed, err:", err)
		}
	}(r.Body)
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status" : "ok"}`
	_, _ = w.Write([]byte(answer))
}
```

### Context

在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

#### 基本示例

```go
var wg sync.WaitGroup

// 基本示例
func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束子goroutine
	wg.Wait()
	fmt.Println("over~")
}
```

#### 全局变量方式

```go
// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。
func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exit = true                 // 修改全局变量实现子goroutine的退出
	wg.Wait()
	fmt.Println("over")
}
```

#### 通道方式

```go
// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel
func worker(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan: // 等待接收上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 3)
	// 给予 goroutine 发送推出的信号
	exitChan <- struct{}{}
	close(exitChan)
	wg.Wait()
	fmt.Println("over~")
}
```

##### 官方方案

```GO
// 官方版
func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	// 通知 goroutine 结束
	cancel()
	wg.Wait()
	fmt.Println("over~")
}
```

当子goroutine又开启另外一个goroutine时，只需要将ctx传入即可：

```GO
func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	// 通知 goroutine 结束
	cancel()
	wg.Wait()
	fmt.Println("over~")
}
```

#### Context 初识

Go1.7加入了一个新的标准库context，它定义了Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。

#### Context 接口

context.Context是一个接口，该接口定义了四个需要实现的方法。具体签名如下：

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

其中：

- Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
- Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
- Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
  - 如果当前Context被取消就会返回Canceled错误；
  - 如果当前Context超时就会返回DeadlineExceeded错误；
- Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

#### Background()和 TODO()

Go内置两个函数：Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。

Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。

background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

#### With 系列函数

此外，context包中还定义了四个With系列函数。

##### WithCancel()

WithCancel的函数签名如下：

```go
    func WithCancel(parent Context) (ctx Context, cancel CancelFunc) 
```

WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

```go
func gen(ctx context.Context) <-chan int {
        dst := make(chan int)
        n := 1
        go func() {
            for {
                select {
                case <-ctx.Done():
                    return // return结束该goroutine，防止泄露
                case dst <- n:
                    n++
                }
            }
        }()
        return dst
    }
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // 当我们取完需要的整数后调用cancel

    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
}
```

上面的示例代码中，gen函数在单独的goroutine中生成整数并将它们发送到返回的通道。gen的调用者在使用生成的整数之后需要取消上下文，以免gen启动的内部goroutine发生泄漏。

##### WithDeadline()

WithDeadline的函数签名如下：

```
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```

返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

```go
func main() {
    d := time.Now().Add(50 * time.Millisecond)
    ctx, cancel := context.WithDeadline(context.Background(), d)

    // 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
    // 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
    defer cancel()

    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
} 
```

上面的代码中，定义了一个50毫秒之后过期的deadline，然后我们调用context.WithDeadline(context.Background(), d)得到一个上下文（ctx）和一个取消函数（cancel），然后使用一个select让主程序陷入等待：等待1秒后打印overslept退出或者等待ctx过期后退出。 因为ctx50毫秒后就过期，所以ctx.Done()会先接收到值，上面的代码会打印ctx.Err()取消原因。

##### WithTimeout

WithTimeout的函数签名如下：

```
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) 
```

WithTimeout返回WithDeadline(parent, time.Now().Add(timeout))。

取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。具体示例如下：

```go
var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("do connection...")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): // 50ms 后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over~")
}
```

##### WithValue

WithValue函数能够将请求作用域的数据与 Context 对象建立关系。声明如下：

```
    func WithValue(parent Context, key, val interface{}) Context 
```

WithValue返回父节点的副本，其中与key关联的值为val。

仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。

所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。WithValue的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。或者，导出的上下文关键变量的静态类型应该是指针或接口。

```go
type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在 goroutine 获取 trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): // 50ms 后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("work done~")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
```

#### 注意事项

- 推荐以参数的方式显示传递Context
- 以Context作为参数的函数方法，应该把Context作为第一个参数。
- 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
- Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
- Context是线程安全的，可以放心的在多个goroutine中传递

#### 客户端超时取消示例

调用服务端API时如何在客户端实现超时控制？

##### server 端

```GO
// server 端 随机出现慢响应
func indexHandle(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(2)
	if num == 0 {
		time.Sleep(time.Second * 10) // 耗时 10s 的慢响应
		_, _ = fmt.Fprintf(w, "slow response!")
		return
	}
	_, _ = fmt.Fprintf(w, "quick respinse!")
}

func main() {
	http.HandleFunc("/", indexHandle)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
```

#####client 端

```GO
// client端
type resData struct {
	res *http.Response
	err error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		// 请求频繁可定义全局的 client 对象并启用长链接 不频繁的使用短连接
		DisableKeepAlives: true}
	client := http.Client{
		Transport: &transport,
	}
	resChan := make(chan *resData, 1)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/", nil)
	if err != nil {
		fmt.Printf("new request failed, err:%v\n", err)
		return
	}
	// 使用带超时的 ctx 创建一个新的 client request
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		res, err := client.Do(req)
		fmt.Printf("client.do res:%v, err:%v\n", res, err)
		rd := &resData{
			res: res,
			err: err,
		}
		resChan <- rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-resChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(result.res.Body)
		data, _ := ioutil.ReadAll(result.res.Body)
		fmt.Printf("res:%v\n", string(data))
	}
}

func main() {
	// 定义 100ms 超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel() // 调用 cancel 释放 goroutine 资源
	doCall(ctx)
}
```

### 数据格式

#### 数据格式介绍

- 是系统中数据交互不可缺少的内容
- 这里主要介绍JSON、XML、MSGPack

#### JSON

- json是完全独立于语言的文本格式，是k-v的形式 name:zs

- 应用场景：前后端交互，系统间数据交互

  ```json
  {
      "meta":{
          "filter_count":18
      },
      "data":[
          {
              "created_on":"2022-03-06 06:13:18",
              "version":"16.0.9",
              "content":"Navicat for MySQL 16.0.9\nFixed:\n- The \"Compare\" button did not work in Data Synchronization in some cases\n- \"Couldn't open known_hosts file\" error occurred when the SSH path contained Chinese characters\n- Minor bug fixes and improvements"
          }
      ],
      "public":true
  }
  ```

  - json使用go语言内置的encoding/json 标准库
  - 编码json使用json.Marshal()函数可以对一组数据进行JSON格式的编码

  ```
      func Marshal(v interface{}) ([]byte, error) 
  ```

  示例过结构体生成json：

  ```go
  // 数据格式
  type Equip struct {
  	Name         string
  	Introduction string
  	Occupation   string
  	Estate       int64
  }
  
  func main() {
  	equip := Equip{
  		"破军",
  		"北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋",
  		"东夷战士",
  		120,
  	}
  	// 编码 json
  	b, err := json.Marshal(equip)
  	if err != nil {
  		fmt.Println("json err:", err)
  	}
  	fmt.Println(string(b))
  
  	// 格式化输出
  	b, err = json.MarshalIndent(equip, "", "	")
  	if err != nil {
  		fmt.Println("json err ", err)
  	}
  	fmt.Println(string(b))
  }
  
  // 输出
   $ go run .
  {"Name":"破军","Introduction":"北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋","Occupation":"东夷战士","Estate":120}
  {
          "Name": "破军",
          "Introduction": "北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋",
          "Occupation": "东夷战士",
          "Estate": 120
  }
  ```

  ##### struct tag

  ```go
  type Equip struct {
      Name         string `json:"name"`
      Introduction string `json:"introduction"`
      Occupation   string `json:"occupation"`
      Estate       int64 `json:"estate"`
  }
  ```

  示例通过map生成json：

  ```GO
  // 示例通过map生成json
  func mapJson() {
  	student := make(map[string]interface{})
  	student["name"] = "星河万里"
  	student["age"] = 18
  	student["sex"] = "man"
  	b, err := json.Marshal(student)
  	if err != nil {
  		fmt.Println(err)
  	}
  	// 格式化输出
  	b, err = json.MarshalIndent(student, "", "	")
  	if err != nil {
  		fmt.Println("json err ", err)
  	}
  	fmt.Println(string(b))
  }
  
  // 输出
  $ go run .
  {                 
          "age": 18,
          "name": "星河万里",
          "sex": "man"
  }
  ```

  - 解码json使用json.Unmarshal()函数可以对一组数据进行JSON格式的解码

  ```
      func Unmarshal(data []byte, v interface{}) error
  ```

  示例解析到结构体:

  ```go
  // 示例解析到结构体
  func jsonStruct() {
  	b := []byte(`{"name":"破军","introduction":"北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋","occupation":"东夷战士","estate":120}`)
  	var e Equip
  	err := json.Unmarshal(b, &e)
  	if err != nil {
  		fmt.Println(err)
  	}
  	// 格式化输出
  	b, err = json.MarshalIndent(e, "", "	")
  	if err != nil {
  		fmt.Println("json err ", err)
  	}
  	fmt.Println(string(b))
  }
  
  // 输出
  $ go run main.go
  {
          "name": "破军",
          "introduction": "北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋",
          "occupation": "东夷战士",
          "estate": 120
  }
  ```

  示例解析到接口：

  ```GO
  // 示例解析到接口
  func jsonInterface() {
  	// 声明接口
  	var i interface{}
  	err := json.Unmarshal(b, &i)
  	if err != nil {
  		fmt.Println(err)
  	}
  	// 自动转到map
  	fmt.Println(i)
  	// 可以判断类型
  	m := i.(map[string]interface{})
  	for k, v := range m {
  		switch vv := v.(type) {
  		case float64:
  			fmt.Println(k, "是float64类型", vv)
  		case string:
  			fmt.Println(k, "是string类型", vv)
  		default:
  			fmt.Println("other type", vv)
  		}
  	}
  }
  
  // 输出
  $ go run main.go
  map[estate:120 introduction:北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋 name:破军 occupation:东夷战士]
  name 是string类型 破军
  introduction 是string类型 北斗第七星，有名破军。破军主破，其利天下无敢拂其锋芒，刀魂暴戾，易走偏锋
  occupation 是string类型 东夷战士
  estate 是float64类型 120
  ```

#### XML

- 是可扩展标记语言，包含声明、根标签、子元素和属性
- 应用场景：配置文件以及webService

示例：

```xml
    <?xml version="1.0" encoding="UTF-8" ?>
    <servers version="1">
        <server>
            <serverName>Shanghai_VPN</serverName>
            <serverIP>127.0.0.1</serverIP>
        </server>
        <server>
            <serverName>Beijing_VPN</serverName>
            <serverIP>127.0.0.2</serverIP>
        </server>
    </servers>
```

```go
// Server 抽取单个server对象
type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version int      `xml:"version"`
	Servers []Server `xml:"server"`
}

func main() {
	data, err := ioutil.ReadFile("./xml.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 格式化输出
	b, err := json.MarshalIndent(servers, "", "	")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}

// 输出
$ go run .
{
        "Name": {
                "Space": "",
                "Local": ""
        },
        "Version": 0,
        "Servers": [
                {
                        "ServerName": "Shanghai_VPN",
                        "ServerIP": "127.0.0.1"
                },
                {
                        "ServerName": "Beijing_VPN",
                        "ServerIP": "127.0.0.2"
                }
        ]
}
```

#### MSGPack

- MSGPack是二进制的json，性能更快，更省空间
- 需要安装第三方包：go get -u github.com/vmihailenco/msgpack

```go
package MSGPack

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"math/rand"
)

type Person struct {
	Name, Sex string
	Age       int
}

// 二进制写出
func writeJson(filename string) (err error) {
	var persons []*Person
	// 假数据
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Sex:  "male",
			Age:  rand.Intn(100),
		}
		persons = append(persons, p)
	}
	// 二进制json序列化
	data, err := msgpack.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// 二进制读取
func readJson(filename string) (err error) {
	var persons []*Person
	// 读文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 反序列化
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range persons {
		fmt.Printf("%#v\n", v)
	}
	return
}

func main() {
	err := readJson("filepath")
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

























