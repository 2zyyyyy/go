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
  package main
  
  import (
  	"bufio"
  	"fmt"
  	"io"
  	"os"
  )
  
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
  package main
  
  import (
  	"fmt"
  	"io/ioutil"
  )
  
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
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

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





















