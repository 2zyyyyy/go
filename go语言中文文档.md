## Go 语言中文文档记录

### Go基础

#### 1、Go语言的主要特征

- 自动立即回收
- 更丰富内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

#### Go文件名

​	所有的go源码都是以`*.go`结尾

#### Go语言命名

##### Go的函数、变量、常量、自定义类型、包（package）的命名方式遵循以下规则：

- 首字符可以使任意的Unicode字符或者下户线
- 剩余字符可以使Unicode字符、下划线、数字
- 字符长度不限

##### Go只有25个关键字

```go
	 break        default      func         interface    select
 case         defer        go           map          struct
 chan         else         goto         package      switch
 const        fallthrough  if           range        type
 continue     for          import       return       var
```

##### Go还有37个保留字

```go
Constants: true  false  iota  nil

Types: int  int8  int16  int32  int64  
       uint  uint8  uint16  uint32  uint64  uintptr
       float32  float64  complex128  complex64
       bool  byte  rune  string  error

Functions: make  len  cap  new  append  copy  close  delete
           complex  real  imag
           panic  recover
```

#### 可见性

- 声明在函数内部，是函数的本地值，类似**private**
- 声明在函数外部，是对当前包可见（包内所有.go文件都可见）的全局值，类似**protect**
- 声明在函数外部且首字母大写是对所有包可见的全局值，类似**public**

#### Go语言声明

​	有4中主要的声明方式：

```go
var (声明变量)
const (声明常量)
type (声明类型)
func (声明函数)
```

​	Go的程序是保存在多个.go文件中，文件的第一行就是package XXX声明，用来说明该文件属于哪个包(package)，package声明下来就是import声明，再下来是类型，变量，常量，函数的声明。

#### Go项目构建及编译

​	一个Go工程中主要包含以下三个目录：

```go
src:源代码文件
pkg:包文件
bin:相关bin文件
```

#### Go编译问题

​	golang的编译使用命令`go build`，`go install`；除非仅写一个main函数，否则还是准备目录结构；

​	gopath=工程目录；其下应创建src、pkg、bin目录，bin目录中用语生产可执行文件，pkg目录中用于生产.a文件。golang中的import name，实际是到GOPATH中去寻找name.a, 使用时是该name.a的源码中生命的package 名字；



#### 2、Go内置类型及函数

##### 内置类型

- 值类型

  ```go
      bool
      int(32 or 64), int8, int16, int32, int64
      uint(32 or 64), uint8(byte), uint16, uint32, uint64
      float32, float64
      string
      complex64, complex128
      array    -- 固定长度的数组
  ```

- 引用类型（指针类型）：

  ```go
      slice   -- 序列数组(最常用)
      map     -- 映射
      channel    -- 管道
  ```

##### 内置函数

​	Go语言拥有一些不需要进行导入操作就可以使用的内置函数。他们又是可以针对不同的类型进行操作，如：`len()` `cap()` 和`append()`，或必须用于系统级的操作，如:`panic`，因此，他们需要直接获得编译器的支持。

```go
		append          -- 用来追加元素到数组、slice中,返回修改后的数组、slice
    close           -- 主要用来关闭channel
    delete            -- 从map中删除key对应的value
    panic            -- 停止常规的goroutine  （panic和recover：用来做错误处理）
    recover         -- 允许程序定义goroutine的panic动作
    real            -- 返回complex的实部   （complex、real imag：用于创建和操作复数）
    imag            -- 返回complex的虚部
    make            -- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
    new                -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
    cap                -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
    copy            -- 用于复制和连接slice，返回复制的数目
    len                -- 来求长度，比如string、array、slice、map、channel ，返回长度
    print、println     -- 底层打印函数，在部署环境中建议使用 fmt 包
```

##### 内置接口error

```go
type error interface { // 只要实现了Error()函数，返回值为String的都实现了err接口
  Error() String
}
```



#### 3、Init函数和main函数

##### init函数

​	go语言中`init`函数用语包`package`的初始化，该函数是go语言的一个重要特性。其有以下特征：

```go
		1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
    2 每个包可以拥有多个init函数
    3 包的每个源文件也可以拥有多个init函数
    4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
    5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
    6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
```

##### main函数

​	Go语言程序的默认入口函数（主函数）:func main()函数体用{}包裹。

```go
func main() {
  // 函数体
}
```

##### init函数和main函数的异同

- 共同点
  - 两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
- 不同点
  - init可以应用于任何包中，且可以重复定义多个
  - main函数只能用于main包中，且只能定义一个

两个函数的执行顺序：

​	对同一个go文件的`init()`调用顺序是从上到下的。

​	对同一个package中不同文件是按文件名字符串比较从小到大的顺序调用个文件中的`init()`函数。

​	对于不同的package，如果不相互依赖的话，按照main包中”先import的后调用“的顺序调用其包中的`init()`，如果package存在依赖，则先调用最早被依赖的package中的`init()`，最后调用`main()`函数。

​	如果`init()`函数中使用了`PrintLn()`或者`Print()`你会发现在执行过程中这两个不会按照你想象中的顺序执行。这两个函数官方只推荐在测试环境中使用，对于正式环境不要使用。

#### 4、命令

​	假设你已安装了golang环境，我们可以在终端通过go命令查看相关的go语言命令：

```go
 ~/go/src/go   master ±  go            
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        module-private  module configuration for non-public modules
        packages        package lists and patterns
        testflag        testing flags
        testfunc        testing functions

Use "go help <topic>" for more information about that topic.
```

|      命令       | 说明                                                         |
| :-------------: | ------------------------------------------------------------ |
|    `go env`     | 打印Go语言的环境信息                                         |
|    `go run`     | 编译并运行命令源码文件                                       |
|    `go get`     | 根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装 |
|   `go build`    | 编译我们指定的源码文件或代码包以及它们的依赖包               |
|  `go install`   | 编译并安装指定的代码包及它们的依赖包                         |
|   `go clean`    | 删除掉执行其它命令时产生的一些文件和目录                     |
|    `go doc`     | 打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的 |
|    `go test`    | 用于对Go语言编写的程序进行测试                               |
|    `go list`    | 列出指定的代码包的信息                                       |
|    `go fix`     | 把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码 |
|    `go vet`     | 检查Go语言源码中静态错误的简单工具                           |
| `go tool pprof` | 交互式的访问概要文件的内容                                   |

#### 5、运算符

​	Go内置的运算符有：

```go
算数运算符
关系运算符
逻辑运算符
位运算符
赋值运算符
```

##### 算术运算符

| 运算符 | 描述 |
| ------ | ---- |
| +      | 相加 |
| -      | 相减 |
| *      | 相乘 |
| /      | 相除 |
| %      | 求余 |

**注意：**++（自增）和—（自减）在Go语言中是单独的语句，并不是运算符。

##### 关系运算符

| 运算符 | 描述                                                         |
| ------ | ------------------------------------------------------------ |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 |

##### 逻辑运算符

| 运算符 | 描述                                                         |
| :----- | :----------------------------------------------------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。 |
| ll     | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。 |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。 |

##### 位运算符

​	位运算符对整数在内存中的二进制位进行操作。

| 运算符 | 描述                                                         |
| :----- | :----------------------------------------------------------- |
| &      | 参与运算的两数各对应的二进位相与。（两位均为1才为1）         |
| l      | 参与运算的两数各对应的二进位相或。（两位有一个为1就为1）     |
| ^      | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1） |
| <<     | 左移n位就是乘以2的n次方。“a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。 |
| >>     | 右移n位就是除以2的n次方。“a>>b”是把a的各二进位全部右移b位。  |

##### 赋值运算符

| 运算符 | 描述                                           |
| :----- | :--------------------------------------------- |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
| +=     | 相加后再赋值                                   |
| -=     | 相减后再赋值                                   |
| *=     | 相乘后再赋值                                   |
| /=     | 相除后再赋值                                   |
| %=     | 求余后再赋值                                   |
| <<=    | 左移后赋值                                     |
| >>=    | 右移后赋值                                     |
| &=     | 按位与后赋值                                   |
| l=     | 按位或后赋值                                   |
| ^=     | 按位异或后赋值                                 |

#### 6、下划线

​	`_`是特殊标识符，用来忽略结果。

##### 下划线在import中

​	在Go语言中，`import`的作用是导入其他`package`.

​	import下划线（如：import  _hello/imp）的作用：当导入一个包时，该包下的文件里所有的`init()`函数都会被执行。然而，有些时候我们并不需要把整个包都导入进来，仅仅是希望他执行`init()`函数而已。这个时候就可以使用`import _`引用该包。即使用：import _包路径，只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。

​	代码结构：

```go
src
 |
 +--- main.go
 |
 +--- hello
				|
				+--- hello.go
```

```go
package main

import (
	"go/github.io/2zyyyyy/hello"
	_ "go/github.io/2zyyyyy/hello" // import使用了'_',则会编译报错：./main.go:9:2: undefined: hello
)

func main() {
	hello.Hello()
}

```

hello.go

```go
package hello

import "fmt"

// _ 下划线在import中的应用
func init()  {
	fmt.Println("import--init() comme here!")
}

func Hello()  {
	fmt.Println("hello~")
}
```

使用`import _ "./hello"`,编译程序：

```go
./main.go:9:2: undefined: hello
```

正常`import "./hello"`,编译程序：

```go
import--init() comme here!
hello~

Process finished with exit code 0
```

##### 下划线在代码中

```go
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

func main() {
  underline()
}
```

解释1：

```go
		下划线意思是忽略这个变量.

    比如os.Open，返回值为*os.File，error

    普通写法是f, err := os.Open("xxxxxxx")

    如果此时不需要知道返回的错误值

    就可以用f, _ := os.Open("xxxxxx")

    如此则忽略了error变量
```

解释2：

```go
		占位符，意思是那个位置本应赋给某个值，但是咱们不需要这个值。
    所以就把该值赋给下划线，意思是丢掉不要。
    这样编译器可以更好的优化，任何类型的单个值都可以丢给下划线。
    这种情况是占位用的，方法返回两个结果，而你只想要一个结果。
    那另一个就用 "_" 占位，而如果用变量的话，不使用，编译器是会报错的。
```

补充：

```go
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
```

第二个import就是不直接使用MySQL包，只是执行这个包的init函数，把MySQL的驱动注册到sql包里，然后程序里就可以直接使用sql包来访问MySQL数据库了。



#### 7、变量和常量

##### 变量

###### 变量的来历

​	程序运行过程中的数据都是保存在内存中，我们想要在代码中操作某个数据时就需要去内存上找到这个变量，但是如果我们直接在代码中通过内存地址去操作变量的话，代码的可读性会非常差而且还容易出错，所以我们就利用变量将这个数据的内存地址保存起来，以后直接通过这个变量就能找到内存上对应的数据了。

##### 变量类型

​	变量（Variable）的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：整型、浮点型、布尔型等。

​	Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

##### 变量声明

​	Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用。

##### 标准声明

​	Go语言的变量声明格式为：

```go
		var 变量名 变量类型
```

​	变量声明以关键字`var`开头，变量类型放在变量的后面，行尾无需分号。 举个例子：

```go
		var name string
		var age int 
		var sex bool
```

##### 批量声明

​	每声明一个变量就需要`var`关键字会比较繁琐，Go语言中还支持批量声明变量：

```go
var (
		a string
  	b int
  	c bool
  	d float64
)
```

##### 变量的初始化

​	Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为`false`。 切片、函数、指针变量的默认为`nil`。

​	当然我们也可在声明变量的时候为其指定初始值。变量初始化的标准格式如下：

```go
		var 变量名 类型 = 表达式
```

​	举个例子：

```go
		var name string = "用户名"
		var age int = 18
```

​	或者一次性初始化多个变量：

```go
		var name, age = "用户名", 18
```

##### 类型推导

​	有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
		var name = "用户名"
		var age = 18
```

##### 短变量声明

​	在函数内部，可以使用更简略的 := 方式声明并初始化变量。

```go
package main
import "fmt"

// 全局变量m
var m = 100
func main() {
  n := 10
  m := 11
  fmt.Println(m, n)
}
```

##### 匿名变量

​	在使用多重赋值时，如果想要忽略某个值，可以使用`匿名变量（anonymous variable）`。 匿名变量用一个下划线_表示，例如：

``````go
func foo()(int, string) {
  return 10, "test"
}

func main() {
  x, _ := foo()
  _, y := foo()
  fmt.Println("x=", x)
  fmt.Ptintln("y=", y)
}
``````

​	匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。

**注意事项：**

- 函数外的每个语句都必须以关键字开始（var、const、func等）
- :=不能使用在函数外
- _多用于占位符，表示忽略值

##### 常量

​	相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。敞亮的声明和变量声明非常累是，只是把`var` 变成了`const`，常量在定义的时候必须赋值。

```go
const pi = 3.1415
const e = 2.7182
```

​	声明了`pi` 和 `e`这两个常量后，在整个程序运行期间它们的值都不能再发生变化了。

多个常量也可以一起声明：

```go
const (
		pi = 3.1415
  	e = 2.7182
)
```

`const`同时声明多个常量时，如果忽略了值则表示和上面一行的值相同，例如：

```go
const (
		n1 = 100
  	n2
  	n3
)
```

上面的示例中，常量`n1 = n2 = n3 = 100`

##### iota

​	iota是go语言的常量计数器，只能在常量的表达式中使用。

iota在const关键字出现时将被重置为0.const中每增加一行常量声明将使iota计数一次(iota可理解为const与剧中的行索引)。使用iota能简化定义，在定义枚举时很有用。

举个例子：

```go
const (
		n1 = iota // 0
  	n2				// 1
  	n3				// 2
  	n4				// 3
)

// 使用_跳过某些值
const (
		n1 = iota // 0
  	n2				// 1
  	_				
  	n4				// 3
)

// 多个iota定义在一行
const (
		a, b = iota + 1, iota + 2 // 1, 2
  	c, d											// 2, 3
  	e, f											// 3, 4
)
```

#### 8、基本类型

##### 基本类型介绍

​	golang更明确的数字类型命名，支持Unicode，支持常用数据结构

| 类型          | 长度（byte） | 默认值 | 说明                                      |
| ------------- | ------------ | ------ | ----------------------------------------- |
| bool          | 1            | false  |                                           |
| byte          | 1            | 0      | unit8                                     |
| rune          | 4            | 0      | Unicode Code Point, int32                 |
| int, uint     | 4或8         | 0      | 32位或64位                                |
| int8， uint8  | 1            | 0      | -128 ~ 127, 0 ~ 255，byte是uint8 的别名   |
| int16, uint16 | 2            | 0      | -32768 ~ 32767, 0 ~ 65535                 |
| int32, uint32 | 4            | 0      | -21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名 |
| int64, uint64 | 8            | 0      |                                           |
| float32       | 4            | 0.0    |                                           |
| float64       | 8            | 0.0    |                                           |
| complex64     | 8            |        |                                           |
| complex128    | 16           |        |                                           |
| uintptr       | 4或8         |        | 以存储指针的 uint32 或 uint64 整数        |
| array         |              |        | 值类型                                    |
| struct        |              |        | 值类型                                    |
| string        |              | “”     | UTF-8 字符串                              |
| slice         |              | nil    | 引用类型                                  |
| map           |              | nil    | 引用类型                                  |
| channel       |              | nil    | 引用类型                                  |
| interface     |              | nil    | 接口                                      |
| function      |              | nil    | 函数                                      |

支持八进制、 六进制，以及科学记数法。标准库 math 定义了各数字类型取值范围。

```go
     a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
```

空指针值 nil，而非C/C++ NULL。

##### 整型

​	整型分为以下两个大类： 按长度分为：`int8`、`int16`、`int32`、`int64`对应的无符号整型：`uint8`、`uint16`、`uint32`、`uint64`

​	其中，`uint8`就是我们熟知的`byte`型，`int16`对应C语言中的`short`型，`int64`对应C语言中的`long`型。

##### 浮点型

​	Go语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循`IEEE 754`标准： `float32` 的浮点数的最大范围约为`3.4e38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`。

##### 复数

​	`complex64`和`complex128`

复数有实部和虚部，`complex64`的实部和虚部为32位，`complex128`的实部和虚部为64位。

##### 布尔值

​	Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true（真）`和`false（假）`两个值。

**注意：**

- 布尔类型默认为false
- go语言中不允许将整型强制转换为布尔型
- 布尔值无法参与数值运算，也无法与其它类型进行转换

##### 字符串

​	Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型`（int、bool、float32、float64 等）`一样。 Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(“)中的内容，可以在Go语言的源码中直接添加非`ASCII`码字符，例如：

```go
str1 := "test"
str2 := "测试"
```

##### 字符串转义符

​	Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义 | 含义                               |
| :--- | :--------------------------------- |
| \r   | 回车符（返回行首）                 |
| \n   | 换行符（直接跳到下一行的同列位置） |
| \t   | 制表符                             |
| '    | 单引号                             |
| "    | 双引号                             |
| \    | 反斜杠                             |

举个例子，我们要打印Windows平台下的一个文件路径：

```go
package main
imoirt (
		"fmt"
)

func main() {
  fmt.Println("str := \"C:\\pprof\\main.exe\"")
}
```

##### 多行字符串

​	Go语言中要定义一个多行字符串时，就必须使用`反引号`字符：

```go
 s1 := `第一行
    第二行
    第三行
    `
 fmt.Println(s1)
```

##### 字符串的常用操作

| 方法                                | 介绍           |
| :---------------------------------- | :------------- |
| len(str)                            | 求长度         |
| +或fmt.Sprintf                      | 拼接字符串     |
| strings.Split                       | 分割           |
| strings.Contains                    | 判断是否包含   |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string) | join操作       |

##### byte和rune类型

​	组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：

```go
    var a := '中'
    var b := 'x'
```

​	Go 语言的字符有以下两种：

```go
    uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
    rune类型，代表一个 UTF-8字符。
```

	当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。

​	Go 使用了特殊的 `rune` 类型来处理 `Unicode`，让基于 `Unicode`的文本处理更为方便，也可以使用 `byte` 型进行默认字符串处理，性能和扩展性都有照顾

##### 修改字符串

​	要修改字符串，需要先将其转换成`[]rune或[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```go
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
```

##### 类型转换

​	Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

​	强制类型转换的基本语法如下：

```go
    T(表达式)
```

​	其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等。

​	比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。

#### 9、数据Array

​	Golang Array和以往认知的数组有很大不同。

1. 数组：是同一种数据类型的固定长度的序列

2. 数组定义：var a [len]int，var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。

3. 长度是数组类型的一部分，因此，var a[5]int 和var a[10]int是不同的类型。

4. 数组可以通过下标进行访问，下表是从0开始，最后一个元素的下标是：len(array) - 1

   ```go
   for i :=0; i < len(a); i++ {
     ……
   }
   for index, v := range a {
     ……
   }
   ```

5. 访问越界，如果下标在数组的合法范围之外，则触发访问越界，会panic

6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值

7. 支持`++`、`!=`操作符，因为内存总是被初始化过得

8. 指针数组`[n]*T`，数组指针` *[n]T`。

##### 数组初始化

**一维数组**：

```go
// 全局
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "test", 4: "demo"}

// 局部
a := [3]int{1, 2} // 未初始化元素值为0
b := [...]int{1, 2, 3, 4}  // 通过初始化值来确定数组长度
c := [5]int{2: 100, 3: 400} // 使用索引号初始化元素
d := [...]struct {
  name string
  age uint8
}{
  {"user1", 10} // 可省略元素类型
  {"user2", 20}, // 最后一行逗号
}
```

**多维数组：**

```go
// 全局
var arr0[3][5]int
var arr1[2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

// 局部
a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
b := [...][2]int{{1,1}, {2, 2}, {3, 3}}  // 第二维度不能用...
```

​	值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。

```go
package main

import (
    "fmt"
)

func test(arr [2]int) {
  fmt.Printf("x:%p\n", &x)
  x[1] = 1000
}

func main() {
  a := [2]int{}
  fmt.Printf("a:%p\n", &a)
  test(a)
  fmt.Println(a)
}
```

​	内置函数 len 和 cap 都返回数组长度 (元素数量)。

```go
package main

func main() {
  a := [2]int{}
  println(len(a), cap(a))
}
```

**多维数组遍历：**

```go
package main
import (
    "fmt"
)

func main() {
 var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(key=%d, value=%d) = %d\n", k1, k2, v2)
		}
	}
}
```

![image-20211029151822612](https://tva1.sinaimg.cn/large/008i3skNgy1gvw7fuuu78j309i04swem.jpg)

**数据拷贝和传参：**

```go
package main

import "fmt"

func printArray(arr *[5]int) {
  arr[0] = 10
  for i, v := range arr {
    fmt.Println(i, v)
  }
}

func main() {
  var arr1 [5]int
  printArray(&arr1)
  fmt.Println(arr1)
  
  arr2 := [...]int{2, 4, 6, 8, 10}
  printArray(&arr2)
  fmt.Println(arr2)
}
```

**数组练习**

- **求所有元素之和**

  ```go
  package main
  
  import (
  		"fmt"
    	"math/rand"
    	"time"
  )
  
  // 求元素和
  func sumArray(a [10]int) int {
    var sum = 0
    for i :=0; i < len(a); i++ {
      sum += a[i]
    }
    return sum
  }
  
  func main() {
    // 若想做一个真正的随机数，要种子 seed()种子默认是1 rand.Seed(1)
    rand.Seed(time.Now().Unix())
    
    var b [10]int
    for i :=0; i < len(b); i++{
      // 产生一个1~1000随机数
      b[i] = rand.Intn(1000)
    }
    sum := sumArray(b)
    fmt.Printf("sum:%d", sum)
  }
  ```

  ![image-20211029154950940](https://tva1.sinaimg.cn/large/008i3skNgy1gvw8clnyraj30ik02aaa1.jpg)

- **找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）**

  ```go
  package main
  
  import "fmt"
  // 找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
  // 找出两个元素之和等于8的下标分别是（0，4）和（1，2）
  // 找出数组中和为给定值的两个元素的下标
  func testsSum(n [6]int, sum int) {
  	for i := 0; i < len(n); i++ {
  		for j := i; j < len(n); j++ {
  			if n[i]+n[j] == sum {
  				fmt.Println(i, j)
  			}
  		}
  	}
  }
  
  func main() {
    n := [6]int{0, 1, 3, 5, 7, 8}
  	fmt.Println(n)
  	testsSum(n, 8)
  }
  ```

  ![image-20211029171229088](https://tva1.sinaimg.cn/large/008i3skNly1gvwaqkw1wwj307o032a9v.jpg)

#### 10、切片Slice

​	需要说明，slice并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
2. 切片的长度可以改变，因此，切片是一个可变数组。
3. 切片遍历方式和数组一样可以用`len()`求长度。表示可用元素数量，读写操作不能超过该限制。
4. `cap()`可以求出slice最大扩张容量，不能超过数组限制。`0 <= len(slice) <=len(array)`，其中array是slice引用的数组。
5. 切片的定义。var 变量名 []类型，比如：`var s []string    var a []int`。
6. 如果slice == nil，那么len、cap结果都等于0。

##### 创建切片的各种方式

```go
package main 

import "fmt"

func main() {
  // 1.声明切片
  var s1 []int
  if s == nil {
    fmt.Println("s1为空~")
  } else {
    fmt.Println("s1不为空~")
  }
  
  // 2.:=
  s2 := []int{}
  
  // 3.make
  var s3 []int = make([]int, 0)
  
  // 4.初始化赋值
  var s4 []int = make([]int, 0, 0)
  fmt.Println(s4)
  s5 := []int{1, 2, 3}
  fmt.Println(s5)
  
  // 5.从数组切片
  arr := [5]int{1, 2, 3, 4, 5}
  var s6 []int
  // 前包后不包
  s6 = arr[:4]
  fmt.Println(s6)
}
```

**切片初始化**

```go
// 全局
var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, ,9, 10}

var slice0 []int = arr[start: end]
var slice1 []int = arr [:end]
car slice2 []int = arr[start:]
var slice3 []int = arr[:]
var slice4 []int = arr[:len(arr) -1] // 去尾
var slice5 []int = arr.slice(1:len(arr)) // 去头
```

| 操作              | 含义                                                         |
| ----------------- | ------------------------------------------------------------ |
| s[n]              | 切片s中索引位置为n的项                                       |
| s[:]              | 从切片s的索引位置0到len(s)-1处所获得的切片                   |
| s[low:]           | 从切片s的索引位置low到len(s)-1处所获得的切片                 |
| s[:high]          | 从切片s的索引位置0到high处所获得的切片（len=high）           |
| s[low: high]      | 从切片s的索引位置low到high处所获得的切片（len=high-low）     |
| s[low: high: max] | 从切片s的索引位置low到high处所获得的切片(len=high-low,cap=max-low) |
| len(s)            | 切片s的长度，总是<=cap(s)                                    |
| cap(s)            | 切片s的容量，总是>=len(s)                                    |

**通过make来创建切片**

```go
var slice []type = make([]type, len)
slice := make([]type, len)
slice := make([]type, len, cap)
```

![image-20211101145952773](https://tva1.sinaimg.cn/large/008i3skNly1gvznrhttedj30uu09yt98.jpg)

切片内存布局：![image-20211101150130456](https://tva1.sinaimg.cn/large/008i3skNly1gvznt6wt0sj30y60c8mxx.jpg)

读写操作实际目标是底层数组，只需注意索引号的差别。

```go
package main

import "fmt"

func main() {
  data := [...]int{0, 1, 2, 3, 4, 5}
  s := data[2:4]
  s[0] += 200
  s[1] += 100
  
  fmt.Println(s)
  fmt.Println(data)
}
```

![image-20211101150954561](https://tva1.sinaimg.cn/large/008i3skNly1gvzo1x3ggcj309601qa9w.jpg)

​								 修改的是s，但实际是对底层数组（data）的操作

直接创建slice对象，自动分配底层数组。

```go
package main 

import "fmt"

func main() {
  s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号
  fmt.Println(s1, len(s1), cap(s1))
  
  s2 := make([]int, 6, 8) // 使用make创建，指定长度和容量
  fmt.Println(s2, len(s2), cap(s2))
  
  s3 := make([]int, 6) // 初始化未指定cap， 相当于cap = len
  fmt.Println(s3, len(s3), cap(s3))
}
```

使用make动态创建slice，避免了数组必须用常量做长度的麻烦。还可以用指针直接访问底层数组，退化成普通数组操作。

```go
package main

import "fmt"

func main() {
  s := []int{0, 1, 2, 3}
  p := &s[0]  // *int,获取底层数组元素指针
  *p += 1
  
  fmt.Println(s)
  
  // [][]T,是指元素类型为[]T
  data := [][]int{
    {1, 2, 3}
    {10, 20 ,30}
    {100, 200, 300}
  }
  fmt.Println(data)
}
```

![image-20211101154342237](https://tva1.sinaimg.cn/large/008i3skNly1gvzp13sos9j30fi01qgli.jpg)

可直接修改struct、array、slice成员。

```go
package main

import (
    "fmt"
)

func main() {
  d := [5]struct{
    x int
  }{}
  
  s := d[:]
  d[1].x = 10
  d[2].x = 20
  
  fmt.Println(d)
  fmt.Printf("%p, %p\n", &d, &d[0])
}
```

**append操作切片（追加）**

```go
package main

import "fmt"

func main() {
  a := []int{1, 2, 3}
  fmt.Printf("slice a:%v\n", a)
  
  b := []int{4, 5, 6}
  fmt.Printf("slice b:%v\n", b)
  
  c := append(a, b)
  fmt.Printf("slice c:%v\n", c)
}
```

![image-20211101165313156](https://tva1.sinaimg.cn/large/008i3skNly1gvzr1fkb41j308y028mx2.jpg)

append ：向 slice 尾部添加数据，返回新的 slice 对象。

```go
package main

import (
    "fmt"
)

func main() {
  	s1 := make([]int, 0, 5)
    fmt.Printf("%p\n", &s1)
  
    s2 := append(s1, 1)
    fmt.Printf("%p\n", &s2)

    fmt.Println(s1, s2)
}
```

![image-20211101165721937](https://tva1.sinaimg.cn/large/008i3skNly1gvzr5q3p0qj307w02e3yd.jpg)

**超过slice.cap限制，就会重新分配底层数组，即便原有数组未填满**

```go
package main

import (
    "fmt"
)

func main() {
  data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  s := data[:2:3] // 取索引0~1，容量为3
  
  s = append(s, 10, 20) // 添加2个值，超出容量3限制
  fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
  fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}
```

![image-20211101201407347](https://tva1.sinaimg.cn/large/008i3skNgy1gvzwuk0ds5j30dw032aa3.jpg)

​	从输出结果可以看出，append后s重新分配了底层数组，并复制数据。如果只追加一个值，则不会超过s.cap限制，也就不会分配。

​	通常以2倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。或初始化组构成的len属性，改用索引号进行操作。及时释放不再使用的slice对象，避免持有过期数组，造成GC无法回收。

**slice中cap重新分配规律：**

```go
package main

import "fmt"

func main() {
  s11 := make([]int, 0, 1)
	c := cap(s11)
	fmt.Printf("初始容量：%d\n", c)

	for i := 0; i < 50; i++ {
		// 追加
		s11 = append(s11, i)
		if n := cap(s11); n > c {
			fmt.Printf("追加前cap为: %d -> 追加后cap为:%d\n", c, n)
			c = n
		}
	}
}
```

![image-20211102100728896](https://tva1.sinaimg.cn/large/008i3skNgy1gw0kxku3a3j30dk05mq3k.jpg)

**切片拷贝**

```go
package main

import "fmt"

func main() {
  s1 :=[]int{1, 2, 3, 4, 5}
  fmt.Printf("slice s1:%v\n", s1)
  
  s2 := make([]int, 10)
  fmt.Printf("slice s2:%v\n", s2)
  
  copy(s2, s1)
  fmt.Printf("copied slice s1 : %v\n", s1)
  fmt.Printf("copied slice s2 : %v\n", s2)
  
  s3 := []int{1, 2, 3}
  fmt.Printf("slice s3:%v\n", s3)
  s3 = append(s3, s2...)
  fmt.Printf("append s3:%v\n", s3)
  s3 = append(s3, 4, 5, 6)
  fmt.Printf("last s3:%v\n", s3)
}
```

![image-20211102104358412](https://tva1.sinaimg.cn/large/008i3skNgy1gw0lzjm0oij30gw05c0t5.jpg)

**copy()：函数copy在两个slice间复制数据，复制长度以len最小的为准。两个slice可指向同一底层数组，允许元素区间重叠。**

```go
package main

import "fmt"

func main() {
  data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Println("array data:", data)
  
  s1 := data[8:]
  s2 := data[:5]
  fmt.Printf("slice s1:%v\n", s1)
  fmt.Printf("slice s2:%v\n", s2)

  copy(s2, s1)
  fmt.Printf("copied slice s1 : %v\n", s1)
  fmt.Printf("copied slice s2 : %v\n", s2)
  fmt.Println("last array data : ", data)
}
```

![image-20211102112138666](https://tva1.sinaimg.cn/large/008i3skNgy1gw0n2qgkimj30gg04oq38.jpg)

**slice遍历**

```go
package main

import "fmt"

func main() {
  data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  sl := data[:]
  for i, v := range s1 {
    fmt.Printf("i:%v, v:%v\n", i ,v)
  }
}
```

![image-20211102112603408](https://tva1.sinaimg.cn/large/008i3skNly1gw0n7bqh1zj305207uaa2.jpg)

**切片resize（调整大小）**

```go
package main

import "fmt"

func main() {
  a12 := []int{1, 2, 3, 4}
	fmt.Printf("slice a12:%v, len(a12):%v\n", a12, len(a12))

	b12 := a12[1:2]
	fmt.Printf("slice b12:%v, len(b12):%v\n", b12, len(b12))

	c12 := b12[0:3]
	fmt.Printf("slice c12:%v, len(c12):%v\n", c12, len(c12))
}
```

![image-20211102140943922](https://tva1.sinaimg.cn/large/008i3skNgy1gw0rxmpqn7j30cs028jrf.jpg)

**数组和切片的内存布局**

![image-20211102141750577](https://tva1.sinaimg.cn/large/008i3skNgy1gw0s62qjfbj30yk0buwf5.jpg)

**字符串和切片（string and slice）**

​	string底层就是一个byte的数组，因此，也可以进行切片操作。

```go
package main 

import "fmt"

func main() {
 	str := "software tester"
	s1 := str[0:8]
	fmt.Println(s1)

	s2 := str[9:]
	fmt.Println(s2)
}
```

![image-20211102150256204](https://tva1.sinaimg.cn/large/008i3skNly1gw0tgzg9jjj305001mjr6.jpg)

​	string本身是不可变的，因此要改变string中字符。需要如下操作：

```go
package main

import "fmt"

func main() {
  str := "software tester"
  s := []byte(str)  // 中文字符需要用rune
  s[0] := 'S'
  s = append(s, '!')
  str = string(s)
  fmt.Println(str)
}
```

`Software tester!`

**含有中文字符串**

```go
package main

import (
    "fmt"
)

func main() {
  str := "欲穷千里日，更上一层娄！"
	s := []rune(str)
	s[4] = '目'
	s[10] = '楼'
	str = string(s)
	fmt.Println(str)
}
```

![image-20211102151503033](https://tva1.sinaimg.cn/large/008i3skNly1gw0ttl6mncj309c010dfo.jpg)

​	golang slice data[:6:8]两个冒号理解：

​	常规slice，data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为4（6-9）。

​	另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8。

​	a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x。

```go
package main

import (
    "fmt"
)

func main() {
  slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	data1 := slice[6:8]
	fmt.Printf("data1：%v, len(data1):%d, cap(data1):%d\n", data1, len(data1), cap(data1))

	data2 := slice[:6:8]
	fmt.Printf("data2：%v, len(data2):%d, cap(data2):%d\n", data2, len(data2), cap(data2))
}
```

![image-20211102152833650](https://tva1.sinaimg.cn/large/008i3skNgy1gw0u7nwuyxj30je01k3yl.jpg)

**数组或切片转字符串：**

```go
strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
```

#### 11、Slice底层实现

##### 本章不属于基础部分但是面试会经常问建议学学

暂时跳过~

​	切片是 Go 中的一种基本的数据结构，使用这种结构可以用来管理数据集合。切片的设计想法是由动态数组概念而来，为了开发者可以更加方便的使一个数据结构可以自动增加和减少。但是切片本身并不是动态数据或者数组指针。切片常见的操作有 reslice、append、copy。与此同时，切片还具有可索引，可迭代的优秀特性。

![img](https://tva1.sinaimg.cn/large/008i3skNly1gw0uc4yq45j31ai0grgmj.jpg)

……

#### 12、指针

​	区别于C/C++中的指针，go语言中的指针不能进行便宜和运算，是安全指针。

要搞明白go语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值。

**Go语言的指针**

​	Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。Go语言中的指针操作非常简单，只需要记住两个符号：`&`（取地址）和`*`（根据地址取值）。

**指针地址和指针类型**

​	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用&字符放在变量前面对变量进行“取地址”操作。 Go语言中的值类型`（int、float、bool、string、array、struct）`都有对应的指针类型，如：`*int、*int64、*string`等。

```go
// 取变量指针的语法如下：
ptr := &v  // v的类型为T

// 其中：
// 1.v代表被取地址的变量，类型为T
// 2.ptr用于接收地址的变量，ptr的类型就是*T，乘坐T的指针类型。*代表指针

// 举例说明
func main() {
  a20 := 10
	b20 := &a20
	fmt.Printf("a20:%d, ptr:%p\n", a20, &a20)
	fmt.Printf("b20:%p, type:%T\n", b20, b20)
	fmt.Printf("b20:%p, *b20:%d\n", &b20, *b20)
}
```

![image-20211102165007256](https://tva1.sinaimg.cn/large/008i3skNgy1gw0wkiq4g6j30bm02ewek.jpg)

我们来看一下`b := &a`的图示：

![image-20211102165058766](https://tva1.sinaimg.cn/large/008i3skNgy1gw0wlf1k4uj310q0eawf5.jpg)

**指针取值**

​	在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值，代码如下：

```go
func main() {
  	//指针取值
    a := 10
    b := &a // 取变量a的地址，将指针保存到b中
    fmt.Printf("type of b:%T\n", b)
    c := *b // 指针取值（根据指针去内存取值）
    fmt.Printf("type of c:%T\n", c)
    fmt.Printf("value of c:%v\n", c)
}
```

​	总结： 取地址操作符&和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址取出地址指向的值。

​	变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

1. 对变量进行取地址（&）操作，可以获得这个变量的指针地址。
2. 指针变量的值是指针地址。
3. 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

指针传值实例：

```go
func modify1() {
  x = 100
}

func modify2(x *int) {
  *x = 100
}

func main() {
  a := 10
  modify1(a) // 10
  fmt.Println(a)
  
  modify2(&a)
  fmt.Println(a) // 100
}
```

**空指针**

- 当一个指针被定义后没有分配任何变量时，它的值为nil

- 空指针判断

  ```go
  package main
  
  import "fmt"
  
  func main() {
    var p *string
    fmt.Println(p)
    fmt.Printf("p的值是%v\n", p)
     if p != nil {
        fmt.Println("非空")
     } else {
        fmt.Println("空值")
     }
  }
  ```

**new和make**

先看实例：

```go
func main() {
  var a *int
  *a = 100
  fmt.Println(*a)
  
  var b map[string]int
  b["测试"] =  100
  fmt.Println(b)
}
```

执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存

**new**

​	new是一个内置的函数，它的函数签名如下：

```go
func new(Type) *Type
```

```go
其中：
1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
```

​	new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：

```go
func main() {
    a := new(int)
    b := new(bool)
    fmt.Printf("%T\n", a) // *int
    fmt.Printf("%T\n", b) // *bool
    fmt.Println(*a)       // 0
    fmt.Println(*b)       // false
} 
```

​	本节开始的示例代码中`var a *int`只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：

```go
func main() {
  var a *int
  a = new(int)
  *a = 10
  fmt.Println(*a)
}
```

**make**

​	mak也是用语内存分配的，区别于new，他只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：

```go
func make(t Type, size ...InterType) Type
```

​    make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对他们进行操作。这个我们上一章中都有说明，关于channel我们会在后续的章节详细说明。

​	本节开始的示例中`var b map[string]int `只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其赋值：

```go
func main() {
  var b map[string]int
  b = make(map[string]int, 10)
  b["测试"] = 100
  fmt.Println(b)
}
```

**new和make的区别**

```go
1.两者都是用来做内存分配的
2.make只用于slice、map和channel的初始化，返回的值还是这三个引用类型本身
3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
```

**指针练习**

- 程序定义一个int变量num的地址并打印
- 将num的地址赋给指针ptr，并通过ptr去修改num的值

```go
func main() {
  var num int
  fmt.Println(num)
  
  ptr := &num
  *ptr = 20
  fmt.Println(num)
}
```

#### 13、Map

​	map是一种无序的基于k-v的数据结构，go语言中的map是引用类型，必须初始化才能使用。

**map定义**

​	go语言中map定义语法如下：

```go
map[keyType] valueType
```

其中：

```go
keyType:表示键的类型
valueType:表示键对应值的类型
```

map类型的变量默认初始值为nil，需要使用make函数来分配内存。语法为：

```go
make(map[KeyType]ValueType, [cap])
```

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

**map的基本使用**

​	map中的数据都是成对出现的，map的基本使用实例代码如下：

```go
func main() {
  scoreMap := make(map[string]int, 8)
  scoreMap["张三"] = 90
  scoreMap["李四"] = 95
  fmt.Println(scoreMap)
  fmt.Println(scoreMap["李四"])
  fmt.Printf("type of:%T\n", scoreMap)
}
```

![image-20211124102412157](https://tva1.sinaimg.cn/large/008i3skNly1gwq11qomdhj3096034q31.jpg)

map也支持在声明的时候填充元素：

```go
func main() {
  userInfo := map[string]string {
    "userName" : "wanli",
    "passWord" : "123456",
  }
  fmt.Println(userInfo)
}
```

**判断某个键是否存在**

go语言中有个判断map中是否存在的特殊写法，格式如下：

```go
value, ok := map[key]

// 举例说明
func main() {
  scoreMap := make(map[string]int, 8)
  scoreMap["张三"] = 90
  scoreMap["李四"] = 95
  
  // 如果key存在ok为true，v为对应的值；不存在OK=false v为值类型的零值
  v, ok := scoreMap["张三"]
  if ok {
    fmt.Println(v)
  } else {
    fmt.Println("查无此人~")
  }
}
```

**map的遍历**

​	go语言中使用for range遍历map。

```go
func main() {
  scoreMap := make(map[string]int, 8)
  scoreMap["张三"] = 90
  scoreMap["李四"] = 95
  scoreMap["王五"] = 100
  for k, v := range scoreMap {
    fmt.Println(k, v)
  }
}
```

*注意：遍历map时元素的顺序与添加键值对的顺序无关*

**使用delete()函数删除键值对**

​	使用delete()内建函数从map中该删除一组键值对，delete()函数格式如下：

```go
delete(map, key)

// 其中：
// map：表示要删除键值对的map
// key：表示要删除键值对的键

// 示例：
func main() {
  scoreMap := make(map[string]int, 8)
  scoreMap["张三"] = 90
  scoreMap["李四"] = 95
  scoreMap["王五"] = 100
  
  // 删除 李四:95
  delete(scoreMap, "李四")
  for k, v := range scoreMap{
    fmt.Println(k, v)
  }
}
```

**按照指定顺序遍历map**

```go
func main() {
  rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	scoreMap := make(map[string]int, 200)
	
  // 循环写入数据至map
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("test%2d", i) // 生产test开头的字符串
		value := rand.Intn(100)          // 生产0~99随机整数
		scoreMap[key] = value
	}

	// 取出map中所有的key存入切片keys
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```

**元素为map类型的切片**

​	下面演示了切片中的元素为map类型时的操作：

```go
func main() {
  mapSlice := make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "张三"
	mapSlice[0]["passWord"] = "123456"
	mapSlice[0]["address"] = "未来park"
	mapSlice[0]["age"] = "10"
	mapSlice[0]["sex"] = "男"
	for index, value := range mapSlice {
		fmt.Printf("index:%d, value:%v\n", index, value)
	}
}
```

![image-20211124143712413](https://tva1.sinaimg.cn/large/008i3skNly1gwq8d044kdj30u005ogmb.jpg)

**值为切片类型的map**

​	下面的代码演示了map中值为切片类型的操作：

```go
sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init~~~")
	key := "杭州"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
	fmt.Printf("sliceMap Type is:%T, sliceMap[杭州] Type is:%T\n", sliceMap, sliceMap[key])
```

![image-20211124145422228](https://tva1.sinaimg.cn/large/008i3skNly1gwq8uv20pkj30s003adg3.jpg)

#### 14、Map实现原理

​	**什么是Map**

**key，value存储**

非基础内容后续学习……

#### 15、结构体

​	go语言中没有”类（class）“的概念，也不支持类的继承等面向对象的概念。go语言通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

##### 类型别名和自定义类型

​	**自定义类型**

​	在go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型，go语言中可以使用type关键字来定义自定义类型。

​	自定义类型是定义了一个全新的类型，我们可以基于内置的基本类型定义，也可以通过struct定义。例如：

```go
// 将MyInt定义为int类型
type MyInt int
```

通过Type关键字的定义，MyInt就是一种新的类型，它具有int的特性。

**类型别名**

类型别名是Go1.9版本添加的新功能。

类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。类似张三有小名、乳名、英文名，这些名字都是指向张三这个人。

```go
type TypeAlias = Type
```

我们之前见过的rune和byte就是类型的别名，他们的定义如下：

```go
type byte = uint8
type rune = int32
```

**类型定义和类型别名的区别**

类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解他们之间的区别。

```go
//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
    var a NewInt
    var b MyInt

    fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
    fmt.Printf("type of b:%T\n", b) //type of b:int
} 
```

结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。

##### 结构体

​	go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足要求的，go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型就叫做结构体（struct）。也就是我们可以通过struct来定义自己的类型了。

​	go语言中通过struct来实现面向对象。

**结构体的定义**

使用type和struct关键字来定义结构体，具体代码格式如下：

```go
type 类型名 struct {
  字段名1 字段类型
  字段名2 字段类型
  ……,
}
```

其中：

```go
1.类型名：标识自定义结构体的名称，在同一个包中不能重复
2.字段名：表示结构体字段名。结构体中的字段名必须唯一
3.字段类型：表示结构体字段的具体类型
```

举例说明，我们定义一个Cat（猫）结构体，代码如下：

```go
type Cat struct{
  name string
  city string
  age int8
  sex string
}
// 相同类型的字段也可以写在一行
type Cat struct{
  name, city, sex string
  age int8
}
```

这样我们就拥有一个Cat的自定义类型，它有name、city、age、sex三个字段。分别表示姓名、城市、年龄、性别。这样我们使用这个Cat结构体就能够很方便的在程序中表示和存储猫咪的信息了。

语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型。

**结构体实例化**

只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

```go
var 结构体实例 结构体类型
```

**基本实例化**

```go
func main() {
  // struct
	cats := Cat{}
	cats.name = "二狗子"
	cats.breed = "加菲"
	cats.age = 3
	fmt.Printf("cats=%v\n", cats)  // cats={加菲 二狗子 3}
	fmt.Printf("cats=%#v\n", cats) // cats=main.Cat{breed:"加菲", name:"二狗	子", age:3}
}
```

![image-20211124193326123](https://tva1.sinaimg.cn/large/008i3skNly1gwqgx9wxelj30ji01sjrj.jpg)

我们通过.来访问结构体的字段（成员变量），例如cats.breed和cats.name等。

**匿名结构体**

在定义一些临时数据结构等场景下还可以使用匿名结构体。

```go
func main() {
  var user struct{Name string; Age int}
  user.Name = "测试匿名结构体"
  user.Age = 3
  fmt.Printf("%#v\n", user)
}
```

**创建指针类型结构体**

我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。格式如下：

```go
cats := new(Cat)
fmt.Printf("%T\n", cats)
fmt.PrintF("cats=%#v\n", cats)
```

从打印的结果中我们可以看出cats是一个结构体指针。

需要注意的是在go语言中支持对结构体指针直接使用.开访问结构体的成员。

```go
func main() {
  cats2 := new(Cat)
	fmt.Printf("cats2 type=%T\n", cats2) // cats2 type=*main.Cat
	fmt.Printf("cats2=%#v\n", cats2)     // cats2=&main.Cat{breed:"", 		name:"", age:0}
	cats2.age = 100
	cats2.name = "西西"                // cats2=&main.Cat{breed:"", 	name:"", age:0}
	fmt.Printf("cats2:%#v\n", cats2) // cats2:&main.Cat{breed:"", name:"西西", age:100}
}
```

![image-20211124194917068](https://tva1.sinaimg.cn/large/008i3skNly1gwqhdpqxeuj30jm02a74g.jpg)

**取结构体的地址实例化**

使用&对结构体取地址操作相当于对该结构体类型进行了一次new实例化操作。

```go
func main() {
  cats := &Cat{}
	fmt.Printf("%T\n", cats)
	fmt.Printf("cats:%v\n", cats)
	cats.breed = "中华田园猫"
	cats.age = 10
	cats.name = "技艺"
	fmt.Printf("cats:%#v\n", cats)
}
```

![image-20211126103143678](https://tva1.sinaimg.cn/large/008i3skNly1gwsci89lofj30ne02sjrq.jpg)

`cats.breed = "中华田园猫"`其实在底层是`(*cats).breed = "中华田园猫"`,这是go语言帮我们实现的语法糖。

**结构体初始化**

```go
type Cat struct {
	breed string
	name  string
	age   int8
}

func main() {
  cats := Cat{}
  fmt.Printf("cats:%#v\n", cats)
}
```

**使用键值对初始化**

使用键值对对结构体初始化时，键对应结构体的字段，值对应该字段的初始值。

```go
cats := Cat{
  breed: "布偶",
  age: 10,
  name: "旺财",
}
fmt.Printf("cats:%#v\n", cats)
```

也可以对结构体指针进行键值对初始化，例如：

```go
cats := &Cat{
  breed: "布偶",
  age: 10,
  name: "旺财",
}
fmt.Printf("cats:%#v\n", cats)
```

当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

```go
cats := &Cat{
  breed: "美短",
}
fmt.Printf("cats:%#v\n", cats) // cats:&main.Cat{breed:"美短", name:"", age:0}
```

**使用值的列表初始化**

初始化结构体的时候可以简写，也就是初始化的时候不用写键，直接写值：

```go
cats := &Cat{
  "加菲",
  5,
  "哈哈",
}
fmt.Printf("cats:%#v\n", cats)
```

使用这种格式初始化时，需要注意：

1. 必须初始化结构体的所有字段
2. 初始化的填充顺序必须与字段在结构体中声明的顺序一直
3. 该方式不能和键值初始化方式混用

**结构体内存布局**

```go
type test struct{
  a, b, c, d int8
}
n := test{
  1, 2, 3, 4,
}
fmt.Printf("n.a %p\n", &n.a)
fmt.Printf("n.b %p\n", &n.b)
fmt.Printf("n.c %p\n", &n.c)
fmt.Printf("n.d %p\n", &n.d)
```

![image-20211129150732338](https://tva1.sinaimg.cn/large/008i3skNgy1gww1c5iownj307y03ewej.jpg)

**构造函数**

go语言的结构体没有构造函数，我们可以自己实现。例如，下方的代码就实现了一个cat的构造函数。因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。

```go
func newCat(breed, name string, age int 8) *Cat{
  return &Cat{
    breed: breed,
    name, name,
    age, age,
  }
}
```

调用构造函数

```go
cats := newCat("加菲", "西西" , 3)
fmt.Printf("%#v\n", cats) // &main.Cat{breed:"加菲", name:"西西", age:3}
```

**方法和接收者**

go语言中的方法（method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（receiver）。接收者的概念就类似于其他语言中的this或者self。

方法的定义格式如下：

```go
func(接收者变量 接收者类型) 方法名(参数列表)(返回参数){
  函数体
}
```

其中：

1. 接收者变量：接收者中的参数变量名在命名是，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如Cat类型的接收者变量应该命名为c，Person类型的接收者变量应该命名为p等。
2. 接收者类型：接收者类型和参数类似，可以使指针类型和非指针类型。
3. 方法名、参数列表、返回参数：具体格式与函数定义相同。

举例说明：

```go
// cat结构体
type Cat struct{
  breed, name string,
	age int8
}

// newCat构造函数
func newCat(breed, name string, age int8)*Cat {
  return &Cat{
    breed: breed,
    name: name,
    age: age,
  }
  
// Cat eat的方法
func (c Cat) Eat(){
  fmt.Printf("%s每天就是吃吃吃~", c.name)
}  

  func main() {
    cats := newCat("加菲", "西西", 3)
    cats.Eat()
  }
```

*方法和函数的区别是：函数不属于任何类型，方法属于特定的类型。*

**指针类型的接收者**

指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法的时候修改就守着指针的任意成员变量，在方法结束后，修改都是有效的。这种方式九十分接近于其他语言中面向对象中的this或者self。例如我们为Cat添加一个SetAge方法，来修改实例变量的年龄。

```go
// setAge设置cat的年龄
// 使用指针接收者
func (c *Cat) setAge(newAge int8) {
	fmt.Printf("我是修改年龄的方法，将%d修改为%d\n", c.age, newAge)
	c.age = newAge
}
```

**值类型接收者**

当方法作用域值类型接受者时，go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

```go
// 值接收者
func (c Cat) setName(newName string) {
	fmt.Printf("我是修改name的方法，将%s修改为%s\n", c.name, newName)
	c.name = newName
}

func main() {
  // 调用构造函数newCat()
	cats := newCat("加菲", "西西", 3)
	fmt.Printf("%#v\n", cats)
	
  // 使用值类型接收者修改成员变量
	cats.setName("咚咚咚")
	fmt.Println(cats.name)
}
```

**什么时候应该使用指针类型接收者**

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者

**任意类型添加方法**

在go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

```go
// MyString 将string定义为自定义MyString类型
type MyString string

// 自定义类型MyString的方法
func (m MyString) OutPut() {
	fmt.Println("Hello, 我是一个string。")
}

func main() {
  var str MyString
	str.OutPut()
	str = "test"
	fmt.Printf("%#v  %T\n", str, str) // "test"  main.MyString
}
```

*注意事项：非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。*

**结构体匿名字段**

结构体允许其成员字段再声明时没有字段名而只有类型，这种没有名字的字段就称为你名字短。

```go
// Cat 结构体Cat类型
type Cat struct {
	string
	int // 匿名字段
}
func main() {
  cats := Cat{
    "加菲,"
    11,
  }
  fmt.Printf("%#v\n", cats)
  fmt.Println(cats.string, cats.int)
}
```

匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

**嵌套结构体**

一个结构体中可以嵌套包含另一个结构体或结构体指针。

```go
// struct
type Cat struct {
	breed   string
	name    string
	age     int8
	Address Address
}

// 结构体嵌套
type Address struct {
	Province, City, County string
}

func main() {
  cats := Cat{
		"加菲",
		"西西",
		3,
		Address{
			"浙江省",
			"杭州市",
			"拱墅区",
		},
	}
	fmt.Printf("%#v\n", cats) // main.Cat{breed:"加菲", name:"西西", age:3, Address:main.Address{Province:"浙江省", City:"杭州市", County:"拱墅区"}}
}
```

**嵌套匿名结构体**

```go
type Address struct {
	Province, City, County string
}

// 嵌套匿名结构体
type User struct {
	name    string
	age     int
	Address // 匿名结构体字段 只有类型没有字段名
}
func main() {
  var user User
	user.age = 10
	user.name = "嵌套匿名结构体"
	user.Address.Province = "浙江省" //通过匿名结构体.字段名访问
	user.City = "杭州市"             // 直接访问匿名结构体的字段名
	fmt.Printf("%#v\n", user) // main.User{name:"嵌套匿名结构体", age:10, Address:main.Address{Province:"浙江省", City:"杭州市", County:""}}
}
```

当访问结构体成员是先会在结构体中查找该字段，找不到再去匿名结构体中去查找

**嵌套结构体的字段名冲突**

嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。

```go
type Address struct {
	Province, City, County, Time string
}

type Email struct {
	Account string
	Time    string
}

// 嵌套匿名结构体
type User struct {
	name    string
	age     int
	Address // 匿名结构体字段 只有类型没有字段名
	Email
}

func main() {
  // 嵌套结构体字段冲突
	var user User
	user.name = "字段冲突"
	user.age = 15
  // 指定结构体中的字段给与赋值
	user.Address.Time = "address.time"
	user.Email.Time = "email.time"
	fmt.Printf("%#v\n", user)
}
```

**结构体中的”继承“**

Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

```go
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

func main() {
  dog := &Dog{
		4,
		&Animal{
			"嘻嘻",
		},
	}
	dog.move()
	dog.run()
	fmt.Printf("%#v\n", dog)
}
```

**结构体字段的可见性**

结构体中字段大写开头表示公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

**结构体与JSON序列化**

JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成。JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号””包裹，使用冒号:分隔，然后紧接着值；多个键值之间使用英文,分隔。

```go
//结构体与JSON序列化
type Student struct {
	ID     int
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []*Student
}

func main() {
  // 结构体与JSON序列化
	class := &Class{
		Title:    "中队长",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		class.Students = append(class.Students, stu)
	}
	// JSON序列化：结构体——>JSON格式字符串
	data, err := json.Marshal(class)
	if err != nil {
		fmt.Printf("json marshal failed!%s\n", err)
		return
	}
	fmt.Printf("json:%s\n", data)

	// JSON反序列化：JSON格式字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	class1 := &Class{}
	err = json.Unmarshal([]byte(str), &class1)
	if err != nil {
		fmt.Printf("json unmarshal failed! %s\n", err)
		return
	}
	fmt.Printf("%#v\n", class1)
}
```

**结构体标签（TAG）**

Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。

Tag在结构体字段的后方定义，由一对反引号包裹起来，具体格式如下：

```go
`key1: "value" key2:"value"`
```

结构体标签由一个或者多个键值组成。键和值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

例如我们为Student结构体的每个字段定义json序列化时使用的Tag：

```go
// 结构体标签 tag
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
  // 结构体标签
	student := &Student{
		ID:     1001,
		Gender: "女",
		Name:   "小丑杰克",
	}
	data, err := json.Marshal(student)
	if err != nil {
		fmt.Printf("json marshal failed, err:%s\n", err)
	}
	fmt.Printf("json:%s\n", data) // json:{"id":1001,"gender":"女","name":"小丑杰克"}
}
```

**删除map类型的结构体**

```go
package main

import "fmt"

type Animal struct {
	name string
}

func main(){
  // 删除map类型结构体
	animals := make(map[int]Animal)
	animals[0] = Animal{"花花"}
	animals[1] = Animal{"西西"}
	fmt.Println(animals)

	delete(animals, 0)
	fmt.Println(animals)
}
```

![image-20211206102927003](https://tva1.sinaimg.cn/large/008i3skNly1gx3wn8wjuoj309g01uaa0.jpg)

**实现map有序输出（面试经常问到）**

```go
package main

import(
		"fmt"
  	"sort"
)

func main() {
  // 实现map有序输出
	mapSort := make(map[int]int)
	mapSort[10] = 128
	mapSort[8] = 256
	mapSort[2] = 64
	mapSort[9] = 100
	fmt.Println(mapSort)

	sl := []int{}
	for k := range mapSort {
		fmt.Println(k)
		sl = append(sl, k)
	}
	sort.Ints(sl)
	for i := 0; i < len(mapSort); i++ {
		fmt.Printf("key:%d, value:%d\n", sl[i], mapSort[sl[i]])
	}
}
```

![image-20211206110854993](https://tva1.sinaimg.cn/large/008i3skNly1gx3xryxuicj30ei07g0t2.jpg)



### 流程控制

#### 1、条件语句if

**Go语言条件语句：**

​	条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为true来决定是否执行指定语句，并在条件为false的情况在执行另外的语句。

​	Go语言提供了以下几种条件判断语句：

**if 语句 if 语句 由一个布尔表达式后紧跟一个或多个语句组成。**

Go 编程语言中 if 语句的语法如下：

```go
1.可省略条件表达式括号。
2.持初始化语句，可定义代码块局部变量。
3.代码块左括号必须在条件表达式尾部。

if 布尔表达式 {
  /* 在布尔表达式为true时执行 */
}
```

if 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则不执行。

```go
 x := 0

// if x > 10        // Error: missing condition in if statement
// {
// }

if n := "abc"; x > 0 {     // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
    println(n[2])
} else if x < 0 {    // 注意 else if 和 else 左大括号位置。
    println(n[1])
} else {
    println(n[0])
}     
```

**实例**

```go
package main

import "fmt"

func main(){
  // 定义局部变量
  a := 10
  if a < 20 {
    // 如果条件为true
    fmt.Printf("a小于20\n")
  }
  fmt.Printf("a的值为：%d\n", a)
}
```

**if … else语句if语句后可以使用可选的else语句，else语句中的表达式在布尔表达式为false时执行**

Go 编程语言中 if…else 语句的语法如下：

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}
```

if 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则执行 else 语句块。

**实例**

```go
package main

import "fmt"

func main() {
   /* 局部变量定义 */
   var a int = 100
   /* 判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   } else {
       /* 如果条件为 false 则执行以下语句 */
       fmt.Printf("a 不小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)

}
```

**if嵌套语句 你可以在if或者else if语句中嵌入一个或多个if或else if语句**

Go 编程语言中 if…else 语句的语法如下：

```go
if 布尔表达式 1 {
   /* 在布尔表达式 1 为 true 时执行 */
   if 布尔表达式 2 {
      /* 在布尔表达式 2 为 true 时执行 */
   }
}
```

你可以以同样的方式在 if 语句中嵌套 else if…else 语句

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
   /* 判断条件 */
   if a == 100 {
       /* if 条件语句为 true 执行 */
       if b == 200 {
          /* if 条件语句为 true 执行 */
          fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
       }
   }
   fmt.Printf("a 值为 : %d\n", a )
   fmt.Printf("b 值为 : %d\n", b )
}     
```

#### 2、条件语句switch

​	switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
​	Golang switch 分支表达式可以是任意类型，不限于常量。可省略 break，默认自动终止。

**语法**

Go 编程语言中 switch 语句的语法如下：

```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。

**实例**

```go
// switch
	grade := "B"
	// marks := 90
	var marks int

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	case 50:
		grade = "E"
	default:
		grade = "D"
	}
	fmt.Println(grade, marks)

	switch {
	case grade == "A":
		fmt.Printf("优秀：%s\n", grade)
	case grade == "B", grade == "C":
		fmt.Printf("良好：%s\n", grade)
	case grade == "D":
		fmt.Printf("及格：%s\n", grade)
	case grade == "E":
		fmt.Printf("不及格：%s\n", grade)
	default:
		fmt.Printf("及格：%s\n", grade)
	}
```

**Type Switch**

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

Type Switch语法：

```go
switch x.(Type){
  case type:
  	statement(s)
  case type:
  	statement(s)
  // 你可以定义任意个数的case
  default: // 可选
  	statement(s)
}
```

**实例**

```go
// type switch
	var x interface{}
	// 写法1
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型为:%T\n", i)
	case int:
		fmt.Println("x是int类型")
	case float64:
		fmt.Println("x是float64类型")
	case func(int):
		fmt.Println("x是fun(int)类型")
	case bool, string:
		fmt.Println("x是bool或string类型")
	default:
		fmt.Println("??未知类型")
	}

	// 写法2
	j := 0
	switch j {
	case 0:
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	// 写法3
	k := 0
	switch k {
	case 0:
		println("fallthrough")
		fallthrough
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	//写法三
	var m = 0
	switch m {
	case 0, 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	//写法四
	var n = 0
	switch { //省略条件表达式，可当 if...else if...else
	case n > 0 && n < 10:
		fmt.Println("i > 0 and i < 10")
	case n > 10 && n < 20:
		fmt.Println("i > 10 and i < 20")
	default:
		fmt.Println("default")
	}
```

![image-20211206201310815](https://tva1.sinaimg.cn/large/008i3skNly1gx4dib55idj30ck04aaa2.jpg)















### 3、条件语句select



### 4、循环语句for



#### 5、循环语句range



#### 6、循环控制Goto、Break、Continue





### 函数





### 方法



### 面向对象





### 网络编程



### 并发编程



### 数据操作



### 常用标准库















