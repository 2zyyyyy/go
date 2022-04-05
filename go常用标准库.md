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



































