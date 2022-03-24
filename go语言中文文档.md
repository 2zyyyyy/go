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

**select语句**

​	select语句类似于switch语句，但是select会随机执行一个可运行的case。如果没有case可运行，他将阻塞，知道有case可运行。

​	select是go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。

​	select随机执行一个可运行的case。如果没有case可运行，他将阻塞，知道有case可运行。一个默认的子句应该总是可运行的。

**语法**

Go 编程语言中 select 语句的语法如下：

```go
select{
  case communication clause:
  	statement(s)
  case communication clause:
  	statement(s)
  // 你可以定义任意数量的case
  default: // 可选
  	statement(s)
}
```

以下描述了select语句的语法:

```go
1.每个case都必须是一个通信。
2.所有channel表达式都会被求值。
3.所有被发送的表达式都会被求值。
4.如果任意某个通信可以进行，它就执行；其他被忽略。
5.如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
否组：
1.如果有default语句，则执行该语句。
2.如果没有default语句，select将阻塞，知道某个通信可以运行；go不会重新对channel或值进行求值。
```

**实例**

```go
package main

import "fmt"

func main() {
  // select
	var c1, c2, c3 chan int
	var i1, i2 int
	fmt.Printf("c1:%v, c2:%v, c3:%v\n", c1, c2, c3)
	fmt.Printf("i1:%d, i2:%d\n", i1, i2)
	select {
	case i1 = <-c1:
		fmt.Printf("received %d from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %d to c2\n", i2)
	case i3, ok := <-c3:
		fmt.Printf("i3:%d\n", i3)
		if ok {
			fmt.Printf("received %v from c3\n", i3)
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communivation\n")
	}
}
```

![image-20211207165158021](https://tva1.sinaimg.cn/large/008i3skNly1gx5db7gxqqj30dq02oq2z.jpg)

select可以监听channel的数据流动

select的用法与switch语法非常类似，由select开始的一个新的选择块，每个选择条件由case语句来描述。与switch语句可以选择任何使用相等比较的条件相比，select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作。

```go
select{ // 不停的在这里检测
  case <-chan1: // 检测有没有数据可以读
  // 如果chan成功读取到数据，则进行该case处理语句
  case chan2 <- 1: // 检测有没有数据可以写
  // 如果成功向chan2写入数据，则进行该case处理语句
  
  // 假设没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞(一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源)
  default:
  // 如果以上都没有符合条件，则会进行default处理流程
}
```

在一个select语句中，go会按照从头到尾评估每一个发送和接收的语句。

如果其中的任意一个语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来执行。

如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况：

1. 如果给出了default语句，那么就会执行default的流程，同时程序的执行会从select语句后的语句中恢复。
2. 如果没有default语句，那么select语句将被阻塞，直到至少有一个case可以进行下去。

**Go select的使用及典型用法**

**基本使用**

​	select是go的一个控制结构，类似于switch语句，用户处理异步io操作。select会监听case语句中channel的读写操作，当case中的channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

​	select中的case语句必须是一个channel操作。

​	select中的default子句总是可运行的。

​	如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。

​	如果没有可运行的case语句，且有default语句，那么就会执行default的动作。

​	如果没有可运行的case语句，且没有default语句，select将阻塞，知道某个case通信可以运行。

例如：

```go
func main() {
  var c1, c2, c3 chan int
  var i1, i2 int
  select{
    case i1 = <- c1:
    	fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}

//输出：no communication 
```

**典型用法**

1. 超时判断

```go
//比如在下面的场景中，使用全局resChan来接收response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行
var resChan = make(chan int)
// do request
func test() {
  select{
    case data := <-resChan:
    	doData(data)
    case <- time.After(time.Second * 3)
    	fmt.Println("request time out")
  }
  
  func doData(data int){
    ……
  }
}
```

2. 退出

   ```go
   // 主线程(协程)中如下：
   var shouldQuit = make(chan struct{})
   func main() {
     {
       // loop
     }
     // ...out of the loop
     select{
       case <-c.shouldQuit:
               cleanUp()
               return
           default:
           }
       //...
   }
   
   //再另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
   close(shouldQuit)
   ```

3. 判断channel是否阻塞

   ```go
   // 在某些情况下是存在不希望channel缓存满了的需求的，可以用如下方法判断
   ch := make(chan int, 5)
   // ...
   data := 0
   select{
     case ch <- data:
     default:
     	// 做相应操作，比如丢弃data。
   }
   ```

### 4、循环语句for

**Golang for 支持三中循环方式，包括类似while的语法**

for循环是一个循环控制结构，可以执行指定次数的循环。

**语法**

Go语言的for有3种形式，只有其中的一种使用分号。

```go
1.for init; condition; post{ }
2.for codition{ }
3.for { }

init:一般为赋值表达式，给控制变量赋初值
condition：关系表达式或逻辑表达式，循环控制条件；
post：一般为赋值表达式，给控制变量增量或减量。
for语句执行过程如下：
1.先对表达式init赋初值；
2.判别赋值表达式init是否满足给定condition条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行post，进入第二次循环，在判别condition；否则判断condition的值为假，不满足条件，就终止for循环，执行循环体外语句。
```

```go
str := "abc"
for i, n := 0, len(str); i < n; i++ { // 常见的for循环，支持初始化语句
  println(s[i])
}

n := len(s)
for n >0 { // 替代while(n>0) {}
  n--
  println(str[n])
}

for { // 替代 while (true) {}
		println(s) // 替代 for (;;) {}
}
```

不要期望编译器能理解你的想法，在初始化语句中计算出全部结果是个好主意。

```go
func length(s string) int {
  println("call length.")
  return len(s)
}

func main() {
  s := "abcd"
  for i, n := 0, lenth(s); i < n; i++{ // 避免多次调用length函数
    println(i, s[i])
  }
}
```

**实例**

```go
func main() {
  a := 0
	count := 0
	b := 15

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(numbers))

	// for 1
	for a := 0; a < 10; a++ {
		fmt.Printf("a=:%d\n", a)
		count++
	}
	fmt.Printf("for循环1执行了:%d次\n", count)

	count = 0
	// for 2
	for a < b {
		a++
		fmt.Printf("a的值为：%d\n", a)
		count++
	}
	fmt.Printf("for循环2执行了:%d次\n", count)

	// for 3
	for i, x := range numbers {
		i++
		fmt.Printf("第%d位x的值为%d\n", i, x)
	}
}

// 输出
5
a=:0
a=:1
a=:2
a=:3
a=:4
a=:5
a=:6
a=:7
a=:8
a=:9
for循环1执行了:10次
a的值为：1
a的值为：2
a的值为：3
a的值为：4
a的值为：5
a的值为：6
a的值为：7
a的值为：8
a的值为：9
a的值为：10
a的值为：11
a的值为：12
a的值为：13
a的值为：14
a的值为：15
for循环2执行了:15次
第1位x的值为1
第2位x的值为2
第3位x的值为3
第4位x的值为4
第5位x的值为5
```

**嵌套循环**

在for循环中嵌套一个或多个for循环。

**语法**

```go
for [condition | (init; condition;increment)| Range]
{
  for [condition |(init; condition; increment) | Range]
  {
    statement(s)
  }
  statement(s)
}
```

**实例**

以下实例使用循环嵌套来输出2-100间的素数：

```go
func main() {
  var i, j int
	for i = 2; i < 10; i++ {
		fmt.Printf("i=%d\n", i)
		for j = 2; j <= (i / j); j++ {
			// fmt.Println(j)
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}
```

**无限循环**

如过循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：

```go
func main() {
    for true  {
        fmt.Printf("这是无限循环。\n");
    }
}  
```

#### 5、循环语句range

​	Golang range类似迭代器操作，返回（索引，值）或（键，值）。

for循环的range格式可以对slice、map、数组、字符串等进行迭代循环。格式如下：

```go
for k, v := range oldMap{
  newMap[k] = v
}
```

可忽略不想要的返回值，或“_”这个特殊变量。

```go
func main() {
  s := "abc"
  // 忽略2nd value，支持string/array/slice/map
  for i := range s {
    println(s[i])
  }
  // 忽略index
  for _, a := range s {
    println(a)
  }
  // 全部忽略 仅迭代
  for range s{
    
  }
  
  m := map[string]int{"a": 1, "b":2}
  // 返回k v
  for k, v := range m{
    println(k, v)
  }
}
```

*注意：range会复制对象*

```go
func main() {
  a :=[3]int{1, 2, 3}
  
  for i, v := range a { // i,v都是从复制品中取出
    if i == 1 {
      // 在修改前我们先修改原数组
      a[1], a[2] = 900, 1000
      fmt.Println(a) // 确认修改是有效的, 输出[1, 900, 1000]
    }
    a[i] = v + 100
  }
  fmt.Println(a) // 输出[101, 102, 103]
}
```

建议改用引用类型，其底层数据不会被复制。

```go
func main() {
  // 改用引用类型，其底层数据不会被复制
	s := []int{4, 5, 6, 7, 8}

	for i, v := range s { // 复制struct slice（pointer，len，cap）
		if i == 0 {
			s = s[:3]    // 对slice的修改 不会影响range
			s[2] = 10086 // 对底层数据的修改
		}
		fmt.Println(i, v)
	}
}

// 输出
0 4
1 5
2 10086
3 7
4 8
```

另外两种引用类型map、channel是指针包装，而不像slice是struct。

**for 和 for range有什么区别?**

- 主要是使用场景不同，for可以
  - 遍历array和slice
  - 遍历key为整型递增的map
  - 遍历string

- for range可以完成所有for可以做的事情，却能做到for不能做的，包括
  - 遍历key为string类型的map并同时获取key和value
  - 遍历channel

#### 6、循环控制Goto、Break、Continue

循环控制语句

循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：

**Goto、Break、Continue**

```go
1.三个语句都可以配合标签(label)使用
2.标签名区分大小写，定以后若不使用会造成编译错误
3.continue、break配合标签(label)可用于多层循环跳出
4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同 
```



### 函数

#### 1、函数定义

Golang函数特点

- 无需声明原型
- 支持不定变参
- 支持多返回值
- 支持命名返回参数
- 支持匿名函数和闭包
- 函数也是一种类型，一个函数可以赋值给变量

****

- 不支持嵌套（nested），一个包不能有两个名字一样的函数
- 不支持重载（overload）
- 不支持默认参数（default parameter）

**函数声明**

函数声明包含一个函数名，参数列表，返回值列表和函数体。如果函数没有返回值，则返回列表可以省略。函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。

函数可以没有参数或者接收多个参数。

注意类型在变量名之后。

当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

函数可以返回任意数量的返回值。

使用关键字func定义函数，左大括号依旧不能另起一行。

```go
func test(x, y int, s string)(int, string) {
  // 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号
  n := x + y
  return n, fmt.Sprintf(s, n)
}

func main() {
	n, str := functionDemo(1, 2, "测试functionDemo函数~")
	fmt.Printf("n:%d\nstr:%s\n", n, str)
}
```

函数是第一类对象，可作为参数传递。建议将复杂签名定义为函数类型，以便于阅读。

```go
func test(fn func() int) int {
	return fn()
}

// 定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main() {
	functionDemo(1, 2, "测试functionDemo函数~")
	// fmt.Printf("n:%d\nstr:%s\n", n, str)

	s1 := test(func() int { // 直接将匿名函数当参数
		return 100
	})
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s1, s2)
}

// 输出
100 10, 20
```

有返回值的函数，必须有明确的终止语句，否则会引发编译错误。

你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数标识符。

```go
func Sin(x float64) float //implemented in assembly language
```

#### 2、参数

函数定义时指出，函数定义时有参数，该变量可称为函数的形参。形参就像定义在函数体内部的变量。

但当调用函数，传递过来的变量就是函数的实参，函数可以通过两种方式来传递参数：

1. 值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响实际参数。

   ```go
   func swap(x, y int) int {
     ……
   }
   ```

2. 引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数进行的修改，将会影响到实际参数。

   ```go
   // 定义相互交换值的函数
   func swap(x, y *int) {
     var temp int
     
     temp = *x
     *x = *y
     *y = temp
   }
   
   func main() {
   	a, b := 1, 50
   	swap(&a, &b)
   	fmt.Printf("a=%v, b=%v\n", a, b)
   }
   
   // 输出
   a=50, b=1
   ```

在默认情况下，Go语言使用的是值传递，即在调用过程中不会影响到实际参数。

*注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝。一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。*

*注意2：map、slice、chan、指针、interface默认以引用的方式传递。*

不定参传值

就是函数的参数不是固定的，后面的类型是固定的。（可变参数）

Golang可变参数本质上就是slice。只能有一个，且必须是最后一个。

在参数赋值时可以不用一个一个地赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上…即可。

```go
func myfun(args ...int) { // 0或多个参数
  ……
}
func add(a int, args ...int) int { // 1个或多个参数
  ……
}
func add(a int, b int, args ...int) { // 2个或多个参数
  ……
}
```

*注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数，通过len(arg)来判断传递参数的个数。*

任意类型的不定参数：

就是函数的参数和每个参数的类型都不是固定的。

用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的。

```go
func myfunc(args ...interface{}) {
  ……
}
```

代码：

```go
func myfunc(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprint(s, x)
}

func main() {
  fmt.Println(myfunc("sum:", 1, 100, 1000, 10000))
}

// 输出
sum:11101
```

使用slice对象做变参时，必须展开。（slice...）

```go
func test(s string, n ...int) {
  var x int
  for _, i := range n {
    x += i
  }
  return fmt.Sprintf(s, x)
}

func main() {
  s := []int{1, 2, 3}
  res := test("sum:", s...) // slice... 展开slice
  fmt.Printlf(res)
}

// 输出
sum:15
```

#### 3、返回值

**函数返回值**

`"_"`标识符，用来忽略函数的某个返回值。

Go的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。

返回值的名称应当具有一定意义，可以作为文档使用。

没有参数的return语句将返回各个返回变量的当前值。这种用法被称作“裸”返回。

直接返回语句仅应用当在下面这样的短函数中。在常函数中他们会影响代码的可读性。

```go
func add(a, b int) (c int) {
  c = a + b
  return
}

func calc(a, b int) (sum int, avg int) {
  sum = a + b 
  avg = (a + b) / 2
  return
}

func main() {
  a, b := 1, 2
  c := add(a, b)
  sum, avg := calc(a, b)
  fmt.Println(a, b ,c , sum , avg)
}
```

Golang返回值不能用容器对象接收多返回值。只能用多个变量，或`“_”`忽略。

```go
func test() (int, int) {
  return 1, 2
}

func main() {
  // 错误写法
  s := make([]int, 2)
  s = test() // Error: multiple-value test() in single-value context
	 
  // 正确写法
  x, _ := test()
  fmt.Println(x)
}

// 输出
1
```

多返回值可直接作为其它函数调用实参。

```go
// 多返回值可直接作为其它函数调用实参。
func add(a, b int) (sum int) {
	sum = a + b
	return
}

func test2() (int, int) {
	return 1, 2
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}

	return x
}

func main() {
  // 多返回值可直接作为其它函数调用实参。
	fmt.Println(add(test2()))
	fmt.Println(sum(test2()))
}

// 输出
3
3
```

命名返回参数可以看做与形参类似的局部变量，最后由return隐式返回。

```go
func add(x, y int) (z int) {
  z = x +y
  return
}

func main() {
  fmt.Println(add(1, 2))
}

// 输出
3
```

命名返回参数可被同名局部变量遮蔽，此时需要显示返回。

```go
func add(x, y int) (z int) {
  { // 不能在一个级别，引发 "z redeclared in this block" 错误。
        var z = x + y
        // return   // Error: z is shadowed during return
        return z // 必须显式返回。
    }
}
```

命名返回参数允许defer延迟调用通过闭包读取和修改。

```go
// 命名返回参数允许 defer 延迟调用通过闭包读取和修改
func deferReturn(x, y int) (z int) {
	defer func() {
		fmt.Printf("defer z=%d\n", z)
		z += 100
	}()
	z = x + y
	return
}

func main() {
  	fmt.Println(deferReturn(1, 2))
}

// 输出
defer z=3
103
```

显式return返回前，会先修改命名返回参数。

```go
func add(x, y int) (z int) {
    defer func() {
        println(z) // 输出: 203
    }()

    z = x + y
    return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func main() {
    println(add(1, 2)) // 输出: 203
}

// 输出
203
203
```

#### 4、匿名函数

​	匿名函数是指不需要定义函数名的一种函数实现方式。1958年LISP首先采用匿名函数。

在Go里面，函数可以想普通变量一样被传递或使用，Go语言支持随时在代码里面定义匿名函数。

匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，无需声明。

```go
func main() {
  getSqrt := func(a float64) float64{
    return math.Sqrt(a)
  }
  fmt.Println(getSqrt(4))
}

// 输出
2
```

上面先定义了一个名为getSqrt 的变量，初始化该变量时和之前的变量初始化有些不同，使用了func，func是定义函数的，可是这个函数和上面说的函数最大不同就是没有函数名，也就是匿名函数。这里将一个函数当做一个变量一样的操作。

Golang匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。

```go
func main() {
  // function variable
  fn := func() {
    fmt.Println("hello world.")
  }
  fn()
  
  // function cllection
  fns := [](func(x int) int) {
    func(x int) int {return x + 1},
    func(x int) int {return x + 2},
  }
  fmt.Println(fns[0](100))
  
  // function as field
  d := struct{
    fn func() string
  }{
    fn : func() string {return "hello world!"},
  }
  fmt.Println(d.fn())
  
  // channel of function
  fc := make(chan func() string, 2)
  fc <- func() string {return "hello world."}
  fmt.Println((<-fc)())
}

// 输出
hello world.
101
hello world!
hello world.
```

#### 5、闭包、递归

**闭包详解**

闭包是由函数及其相关引用环境组合而成的实体（即：闭包 = 函数 + 引用环境）。

官方解释：所谓闭包，指的是一个拥有许多变量和绑定了这些变量的环境的表达式（通常是一个函数），因而这些变量也应该是该表达式的一部分。

**Go的闭包**

```GO
package main

import "fmt"

// 闭包（Closure）

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {
	c := a()
	c()
	c()
	c()

	a() // 不会输出i
}

// 输出
1
2
3
```

闭包复制的是原对象指针，这就很容易解释延迟引用现象。

```go
func test() func() {
	x := 100
	fmt.Printf("x(%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x(%p) = %d\n", &x, x)
	}
}

func main() {
	f := test()
	f()
}

// 输出
x(0xc0000140b8) = 100
x(0xc0000140b8) = 100
```

在汇编层 ，test 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、闭包对象指针。当调 匿名函数时，只需以某个寄存器传递该对象即可。

```go
FuncVal { func_address, closure_var_pointer ... }
```

外部引用函数参数局部变量

```go
func add(base int) func(int) int {
  return func(i int) int {
    base += 1
    return base
  }
}

func main() {
  tmp1 := add(10)
  fmt.Println(tmp1(1), tmp(2))
  
  // 此时tmp1和tmp2不是一个实体了
  tmp2 := add(100)
  fmt.Println(tmp2(1), tmp2(2))
}

// 输出
11 12
101 102
```

返回两个闭包

```go
// 返回两个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
  // 定义两个函数并返回
  add := func(i int) int {
    base += i
    return base
  }
  
  sub := func(i int) int {
    base -= i
    return base
  }
  
  return add, sub
}

func main() {
  f1, f2 := test01(10)
  fmt.Println(f1(1), f2(2))
  // 此时base是9
  fmt.Println(f1(3), f2(4))
}

// 输出
11 9
12 8
```

**Go递归函数**

递归，就是在运行的过程中调用自己。一个函数调用自己，就叫做递归函数。

构成递归需要具备的条件：

```go
1.子问题须与原始问题为同样的事，且更为简单
2.不能无限地调用本身，须有个出口，化简为非递归状况处理
```

**数字阶乘**

一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，并且0的阶乘为1。自然数n的阶乘写作n!。1808年，基斯顿·卡曼引进这个表示法。

```go
func factoril(i int) int {
  if i <= 1 {
    return 1
  }
  return i * factoril(i-1)
}

func main() {
  fmt.Println(factoril(3))
}

// 输出
6
```

**斐波那契数列（Fibonacci）**

这个数列从第三项开始，每一项都等于前两项的和。

```go
// 斐波那契数列(Fibonacci)
func fionacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fionacci(i-1) + fionacci(i-2)
}

func main() {
  for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", fionacci(i))
	}
}

// 输出
0
1
1
2
3
5
8
13
21
34
```

#### 6、延迟调用（defer）

**golang延迟调用**

**defer特性**

1. 关键字defer用户注册延迟调用
2. 这些调用直到return前才被执行。因此，可以用来做资源清理
3. 多个的反而语句，按先进后出的方式执行
4. defer语句中的变量，在defer声明时就决定了

**defer用途**

1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

**go语言defer**

go语言的defer功能强大，对于资源管理非常方便，但是如果没用好，也会有陷阱。

defer是先进后出。

这个很自然，后面的语句会依赖前面的资源，因此如果前面的资源先释放了，后面的语句就没法执行了。

```go
func main() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println(i)
	}
}

// 输出
4
3
2
1
0
```

**defer遇上闭包**

```go
func deferClosure() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

func main() {
	deferClosure()
}
// 输出
4
4
4
4
4
```

其实go说的很清楚,我们一起来看看go spec如何说的

Each time a “defer” statement executes, the function value and parameters to the call are evaluated as usualand saved anew but the actual function is not invoked.

也就是说函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.

**defer f.Close**

这个大家用的都很频繁,但是go语言编程举了一个可能一不小心会犯错的例子.

```go
type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func main() {
	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}
	for _, v := range ts {
		defer v.Close()
	}
}
// 输出
c  closed
c  closed
c  closed
```

这个输出并不会像我们预计的输出c b a,而是输出c c c

可是按照前面的go spec中的说明,应该输出c b a才对啊.

那我们换一种方式来调用一下.

```go
type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t Test) {
	t.Close()
}

func main() {
	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}
	for _, v := range ts {
		defer Close(v)
	}
}
// 输出
c  closed
b  closed
a  closed
```

​	defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。但是并没有说struct这里的this指针如何处理，通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待。

​	多个defer注册，按FILO次序执行（先进后出）。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。

```go
func main() {
	test(0)
}

func test(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer func() {
		fmt.Println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()
	defer fmt.Println("c")
}

// 输出
c
b
a
panic: runtime error: integer divide by zero
```

`*`延迟调用参数在注册时求值或复制，可用指针或闭包 “延迟” 读取。

后面的太复杂暂时跳过。

#### 7、异常处理

Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。

异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

**panic：**

```go
1.内置函数
2.假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
3.返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
4.直到goroutine整个退出，并报告错误
```

**recover：**

```go
1.内置函数
2.用来控制一个goroutine的panicKing行为，捕获panic，从而影响应用的行为
3.一般的调用建议
	a).在defer函数中，通过recover来终止一个goroutine的panicking的过程，从而恢复正常代码的执行
	b).可以获取通过panic传递的error
```

**注意：**

```go
1.利用recover处理panic指令，defer必须放在panic之前定义，另外recover只有在defer调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散
2.recover处理异常后，逻辑并不会恢复到panic那个点去，函数跑到defer之后的那个点
3.多个defer会形成defer栈，后定义defer的defer语句会先被调用
```

```go
func main() {
	test()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%T\n", err)
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型
		}
	}()
	panic("panic error!")
}
// 输出
string
panic error!
```

由于panic、recover参数类型为interface{}，因此可抛出任何类型对象。

```go
func panic(v interface{})

func recover() interface{}
```

向已关闭的通道发送数据会引发panic

```go
func main() {
	panicChannel(1)
}

// 向已关闭的通道发送数据引发panic
func panicChannel(n int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch := make(chan int, 10)
	close(ch)
	ch <- n
}
// 输出
send on closed channel
```

延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个可被捕获。

```go
func main() {
	deferPanic()
}

// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func deferPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic error!")
	}()

	panic("deferPanic test!")
}
// 输出
defer panic test！
```

捕获函数recover只有在延迟调用内直接调用才会终止错误，否则总是返回nil。任何未捕获的错误都会沿调用堆栈向外传递。

```go
func main() {
	deferRecover()
}

func deferRecover() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()
	
	defer recover() // 无效

	defer fmt.Println(recover()) // 无效

	defer func() {
		func() {
			fmt.Println("defer inner!")
			recover() // 无效
		}()
	}()
	panic("test panic!!")
}
// 输出
defer inner!
<nil>
test panic!!
```

使用延迟匿名函数或下面这样都是有效的。

```go
func main() {
	testExcept()
}

// 使用延迟匿名函数或下面这样都是有效的。
func except() {
	fmt.Println(recover())
}

func testExcept() {
	defer except()
	panic("test panic!")
}
// 输出
test panic!
```

如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执行。

```go
func main() {
	protectFunc(2,1)
}

// 如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执行
func protectFunc(x,y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic!")
		z = x / y
		return
	}()
	fmt.Printf("x / y = %d\n", z)
}
// 输出
x / y = 0
```

除用panic引发中断性错误外，还可返回error类型错误对象来表示函数调用状态。

```go
type error interface{
  Error() string
}
```

 标准库errors.New和fmt.Errorf函数用于创建实现error接口的错误对象。通过判断错误对象实例来确定具体错误类型。

```go
var ErrorDieByZero = errors.New("division by zero!")

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := division(10, 0); err{
	case nil:
		fmt.Println(z)
	case ErrorDieByZero:
		panic(err)
	}
}

func division(x, y int) (int, error) {
	if y == 0{
		return 0, ErrorDieByZero
	}
	return x/y, nil
}
// 输出
division by zero!
```

Go实现类似try catch的异常处理

```go
func main() {
	Try(func(){
		panic("test panic!")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

// GO tyr catch
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
// 输出
test panic!
```

如何区别使用panic和error两种方式？

惯例是：导致关键流程出现不可修复性错误的使用panic，其他使用error。

#### 8、单元测试

**go test工具**

Go语言中的测试依赖go test命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。

go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。

在`*_test.go`文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。

| 类型     | 格式                  | 作用                           |
| -------- | --------------------- | ------------------------------ |
| 测试函数 | 函数名前缀为Test      | 测试程序的一些逻辑行为是否正确 |
| 基准函数 | 函数名前缀为Benchmark | 测试函数的性能                 |
| 实例函数 | 函数名前缀为Example   | 为文档提供实例文档             |

go test命令会遍历所有的`*_test.go`文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。


golang单元测试对文件名和方法名、参数都有很严格的要求。

```go
1.文件名必须以xx_test.go命名
2.方法必须是Test[^a-z]开头
3.方法参数必须t *testing.T
4.使用go test执行单元测试
```

go test的参数解读：

go test是go语言自带的测试工具，其中包含的是两类，单元测试与性能测试

通过go help test可以看到go test的使用说明：

go test [-c] [-i] [build flags] [packages] [flags for test binary]

参数解读：

-c : 编译go test成为可执行的二进制文件，但是不运行测试。

-i : 安装测试包依赖的package，但是不运行测试。

关于build flags，调用go help build，这些是编译运行过程中需要使用到的参数，一般设置为空

关于packages，调用go help packages，这些是关于包的管理，一般设置为空

关于flags for test binary，调用go help testflag，这些是go test过程中经常使用到的参数

-test.v : 是否输出全部的单元测试用例（不管成功或者失败），默认没有加上，所以只输出失败的单元测试用例。

-test.run pattern: 只跑哪些单元测试用例

-test.bench patten: 只跑那些性能测试用例

-test.benchmem : 是否在性能测试的时候输出内存情况

-test.benchtime t : 性能测试运行的时间，默认是1s

-test.cpuprofile cpu.out : 是否输出cpu性能分析文件

-test.memprofile mem.out : 是否输出内存性能分析文件

-test.blockprofile block.out : 是否输出内部goroutine阻塞的性能分析文件

-test.memprofilerate n : 内存性能分析的时候有一个分配了多少的时候才打点记录的问题。这个参数就是设置打点的内存分配间隔，也就是profile中一个sample代表的内存大小。默认是设置为512 * 1024的。如果你将它设置为1，则每分配一个内存块就会在profile中有个打点，那么生成的profile的sample就会非常多。如果你设置为0，那就是不做打点了。

你可以通过设置memprofilerate=1和GOGC=off来关闭内存回收，并且对每个内存块的分配进行观察。

-test.blockprofilerate n: 基本同上，控制的是goroutine阻塞时候打点的纳秒数。默认不设置就相当于-test.blockprofilerate=1，每一纳秒都打点记录一下

-test.parallel n : 性能测试的程序并行cpu数，默认等于GOMAXPROCS。

-test.timeout t : 如果测试用例运行时间超过t，则抛出panic

-test.cpu 1,2,4 : 程序运行在哪些CPU上面，使用二进制的1所在位代表，和nginx的nginx_worker_cpu_affinity是一个道理

-test.short : 将那些运行时间较长的测试用例运行时间缩短

目录结构：

```go
test
	|
	 ——— calc.go
	|
	 ——— calc_test.go
```

**测试函数**

****

**测试函数的格式**

每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：

```go
func TestName(t *testing.T) {
  ……
}
```

测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举个例子：

```go
func TestAdd(t *testing.T) { …… }
func TestSum(t *testing.T) { …… }
func TestLog(t *testing.T) { …… }
```

其中参数t用语报告测试失败和附加的日志信息。testing.T拥有的方法如下：

```go
func (c *T) Error(args ...interface{})
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...interface{})
func (c *T) Fatalf(format string, args ...interface{})
func (c *T) Log(args ...interface{})
func (c *T) Logf(format string, args ...interface{})
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (c *T) Skip(args ...interface{})
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...interface{})
func (c *T) Skipped() bool
```

**测试函数示例**

就像细胞是构成我们身体的基本单位，一个软件程序也是有很多单元组件构成的。单元组件可以是函数、结构体、方法和最终用户可能依赖的任何东西。总之我们需要确保这些组件是能够正常运行的。单元测试是一些利用各种测试单元组件的程序，它会将结果和预期输出进行比较。

接下来，我们定义一个split的包，包中定义了一个Split函数，具体实现如下：

```go
package split

import (
	"fmt"
	"strings"
)

// golang unit testing demo

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		fmt.Printf("s=%s, i=%d, result=%s\n", s, i, result)
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
```

在当前目录下，我们创建一个split_test.go的测试文件，并定义一个测试函数如下：

```go
package split

import (
	"testing"
	"reflect"
)

func TestSplit(t *testing.T) {
	// 程序输出的结果
	got := Split("a:b:c", ":")
	// 期望结果
	want := []string{"a", "b", "c"}

	// 对比期望及实际结果（slice无法直接比较，借助反射包中的方法对比）
	if !reflect.DeepEqual(want, got) {
		// 测试失败输出错误提示
		t.Errorf("excepted:%v, got:%v\n", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
    got := Split("abcd", "bc")
    want := []string{"a", "d"}
    if !reflect.DeepEqual(want, got) {
        t.Errorf("excepted:%v, got:%v", want, got)
    }
}
```

在split包路径下，执行go test命令：

```go
go test
s=a:b:c, i=1, result=[a]
s=b:c, i=1, result=[a b]
s=abcd, i=1, result=[a]
--- FAIL: TestMoreSplit (0.00s)
    split_test.go:25: excepted:[a d], got:[a cd]
FAIL
exit status 1
FAIL
```

可以看到，两条测试用例中有一条测试没有通过，这种情况我们可以通过在go test添加-v参数，查看测试函数名称和运行时间：

```go
go test -v
=== RUN   TestSplit
s=a:b:c, i=1, result=[a]
s=b:c, i=1, result=[a b]
--- PASS: TestSplit (0.00s)
=== RUN   TestMoreSplit
s=abcd, i=1, result=[a]
    TestMoreSplit: split_test.go:25: excepted:[a d], got:[a cd]
--- FAIL: TestMoreSplit (0.00s)
FAIL
exit status 1
FAIL
```

这次我们可以清楚的看到是TestMoreSplit这个测试没有成功。还可以在go test后添加-run参数，它对应一个正则表达式，只有函数名称匹配上的测试函数才会被执行。

```go
go test -run="More"
s=abcd, i=1, result=[a]
--- FAIL: TestMoreSplit (0.00s)
    split_test.go:25: excepted:[a d], got:[a cd]
FAIL
exit status 1
FAIL   
```

现在我们来解决程序中的问题。很显然我们最初的split函数没有考虑到sep为多个字符的情况，我们来修复一下这个bug：

```go
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		fmt.Printf("s=%s, i=%d, result=%s\n", s, i, result)
    // s = s[i+1:] i+1会导致sep长度大于1 产生bug 需要将1改成sep的长度
		s = s[i+(len(sep)):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
```

这一次我们再来测试一下，我们的程序。注意，当我们修改了我们的代码之后不要仅仅执行那些失败的测试函数，我们应该完整的运行所有的测试，保证不会因为修改代码而引入了新的问题。

```go
go test -v
=== RUN   TestSplit
s=a:b:c, i=1, result=[a]
s=b:c, i=1, result=[a b]
--- PASS: TestSplit (0.00s)
=== RUN   TestMoreSplit
s=abcd, i=1, result=[a]
--- PASS: TestMoreSplit (0.00s)
PASS
ok     
```

**测试组**

****

我们现在还想要测试一下split函数对中文字符串的支持。这个时候我们可以编写一个TestChineseSplit测试函数，但是我们也可以使用如下更友好的一种方式来添加更多的测试用例。

```go
package split

import (
	"fmt"
	"testing"
	"reflect"
)

func TestSplit(t *testing.T) {
	// 定义一个测试类型
	type Test struct {
		input string
		sep string
		want []string
	}

	// 定义一个存储测试用例的切片
	tests := []Test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
        {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
        {input: "abcd", sep: "bc", want: []string{"a", "d"}},
        {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}

	// 遍历切片 逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		// 对比期望及实际结果（slice无法直接比较，借助反射包中的方法对比）
		fmt.Printf("input=%s, sep=%s, tc.want=%#v, got=%#v\n", tc.input, tc.sep, tc.want, got)
		if !reflect.DeepEqual(tc.want, got) {
			// 测试失败输出错误提示
			t.Errorf("excepted:%v, got:%v\n", tc.want, got)
		}
	}

	
}

// 输出
=== RUN   TestSplit
input=a:b:c, sep=:, tc.want=[]string{"a", "b", "c"}, got=[]string{"a", "b", "c"}
input=a:b:c, sep=,, tc.want=[]string{"a:b:c"}, got=[]string{"a:b:c"}
input=abcd, sep=bc, tc.want=[]string{"a", "d"}, got=[]string{"a", "d"}
input=枯藤老树昏鸦, sep=老, tc.want=[]string{"枯藤", "树昏鸦"}, got=[]string{"枯藤", "树昏鸦"}
--- PASS: TestSplit (0.00s)
PASS
ok      
```

**子测试**

****

如果测试用例比较多的时候，我们是没办法一眼看出来具体是那个测试用例失败了。我们可能会想到下面的解决办法,Go1.7+中新增了子测试，我们可以按照如下方式使用t.Run执行子测试：

```go
package split

import (
	"fmt"
	"testing"
	"reflect"
)

func TestSplit(t *testing.T) {
	// 定义一个测试类型
	type Test struct {
		input string
		sep string
		want []string
	}

	// 定义一个存储测试用例的切片
	tests := map[string]Test{
		"testcase P0": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
        "testcase P2": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
        "testcase P3": {input: "abcd", sep: "bc", want: []string{"a", "d"}},
        "testcase P4": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}

	// 遍历切片 逐一执行测试用例
	for name, tc := range tests {
		// 使用t.Run()执行子测试
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			// 对比期望及实际结果（slice无法直接比较，借助反射包中的方法对比）
			fmt.Printf("input=%s, sep=%s, tc.want=%#v, got=%#v\n", tc.input, tc.sep, tc.want, got)
			if !reflect.DeepEqual(tc.want, got) {
			// 测试失败输出错误提示
			t.Errorf("excepted:%v, got:%v\n", tc.want, got)
		}
		})
	}
}

// 输出
=== RUN   TestSplit
=== RUN   TestSplit/testcase_P0
input=a:b:c, sep=:, tc.want=[]string{"a", "b", "c"}, got=[]string{"a", "b", "c"}
=== RUN   TestSplit/testcase_P2
input=a:b:c, sep=,, tc.want=[]string{"a:b:c"}, got=[]string{"a:b:c"}
=== RUN   TestSplit/testcase_P3
input=abcd, sep=bc, tc.want=[]string{"a", "d"}, got=[]string{"a", "d"}
=== RUN   TestSplit/testcase_P4
input=枯藤老树昏鸦, sep=老, tc.want=[]string{"枯藤", "树昏鸦"}, got=[]string{"枯藤", "树昏鸦"}
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/testcase_P0 (0.00s)
    --- PASS: TestSplit/testcase_P2 (0.00s)
    --- PASS: TestSplit/testcase_P3 (0.00s)
    --- PASS: TestSplit/testcase_P4 (0.00s)
PASS
ok     
```

我们还可以通过go test -v -run=Split/testcase P0来执行指定的测试用例：

```go
go test -v -run="Split/testcase P0"
=== RUN   TestSplit
=== RUN   TestSplit/testcase_P0
input=a:b:c, sep=:, tc.want=[]string{"a", "b", "c"}, got=[]string{"a", "b", "c"}
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/testcase_P0 (0.00s)
PASS
ok     
```

**测试覆盖率**

****

测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占中代码的比例。

Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。例如：

```go
go test -cover
input=a:b:c, sep=:, tc.want=[]string{"a", "b", "c"}, got=[]string{"a", "b", "c"}
input=a:b:c, sep=,, tc.want=[]string{"a:b:c"}, got=[]string{"a:b:c"}
input=abcd, sep=bc, tc.want=[]string{"a", "d"}, got=[]string{"a", "d"}
input=枯藤老树昏鸦, sep=老, tc.want=[]string{"枯藤", "树昏鸦"}, got=[]string{"枯藤", "树昏鸦"}
PASS
coverage: 100.0% of statements
ok    
```

从上面的结果可以看到我们的测试覆盖率测试用例覆盖了100%的代码。

Go还提供了一个额外的-coverProfile参数，用来将覆盖率相关的记录信息输出到一个文件。例如：

```go
go test -cover -coverprofile=c.out
input=a:b:c, sep=:, tc.want=[]string{"a", "b", "c"}, got=[]string{"a", "b", "c"}
input=a:b:c, sep=,, tc.want=[]string{"a:b:c"}, got=[]string{"a:b:c"}
input=abcd, sep=bc, tc.want=[]string{"a", "d"}, got=[]string{"a", "d"}
input=枯藤老树昏鸦, sep=老, tc.want=[]string{"枯藤", "树昏鸦"}, got=[]string{"枯藤", "树昏鸦"}
PASS
coverage: 100.0% of statements
ok    
```

上面的命令会将覆盖率相关的信息输出到当前文件夹的c.out文件中，然后我们执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。

![image-20220124191520451](https://tva1.sinaimg.cn/large/008i3skNly1gyoz78y98nj31c00s4aey.jpg)

**基准测试**

****

基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：

```go
func BenchmarkName(b *testing.B) {
  // ...
}
```

基准测试以Benchmark为前缀，需要一个`*testing.B`类型的参数b，基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。testing.B拥有的方法如下:

```go
func (c *B) Error(args ...interface{})
func (c *B) Errorf(format string, args ...interface{})
func (c *B) Fail()
func (c *B) FailNow()
func (c *B) Failed() bool
func (c *B) Fatal(args ...interface{})
func (c *B) Fatalf(format string, args ...interface{})
func (c *B) Log(args ...interface{})
func (c *B) Logf(format string, args ...interface{})
func (c *B) Name() string
func (b *B) ReportAllocs()
func (b *B) ResetTimer()
func (b *B) Run(name string, f func(b *B)) bool
func (b *B) RunParallel(body func(*PB))
func (b *B) SetBytes(n int64)
func (b *B) SetParallelism(p int)
func (c *B) Skip(args ...interface{})
func (c *B) SkipNow()
func (c *B) Skipf(format string, args ...interface{})
func (c *B) Skipped() bool
func (b *B) StartTimer()
func (b *B) StopTimer()  
```

**基准测试示例**

我们为split包中的Split函数编写基准测试如下：

```go
// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}
```

基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试，输出结果如下：

```go
go test -bench=Split
input=a:b:c, sep=:, tc.want=[]string{"a", "b", "c"}, got=[]string{"a", "b", "c"}
input=a:b:c, sep=,, tc.want=[]string{"a:b:c"}, got=[]string{"a:b:c"}
input=abcd, sep=bc, tc.want=[]string{"a", "d"}, got=[]string{"a", "d"}
input=枯藤老树昏鸦, sep=老, tc.want=[]string{"枯藤", "树昏鸦"}, got=[]string{"枯藤", "树昏鸦"}
goos: darwin
goarch: amd64
pkg: go/github.io/2zyyyyy/chineseDocumentation/unitTesting/split
BenchmarkSplit-8        10355961               117 ns/op
PASS
```

其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。

我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。

```go
go test -bench=Split -benchmem
input=a:b:c, sep=:, tc.want=[]string{"a", "b", "c"}, got=[]string{"a", "b", "c"}
input=a:b:c, sep=,, tc.want=[]string{"a:b:c"}, got=[]string{"a:b:c"}
input=abcd, sep=bc, tc.want=[]string{"a", "d"}, got=[]string{"a", "d"}
input=枯藤老树昏鸦, sep=老, tc.want=[]string{"枯藤", "树昏鸦"}, got=[]string{"枯藤", "树昏鸦"}
goos: darwin
goarch: amd64
pkg: go/github.io/2zyyyyy/chineseDocumentation/unitTesting/split
BenchmarkSplit-8        10241151               115 ns/op              48 B/op          2 allocs/op
PASS
```

其中，48 B/op表示每次操作内存分配了48字节，2 allocs/op则表示每次操作进行了2次内存分配。 我们将我们的Split函数优化如下：

```go
goos: darwin
goarch: amd64
pkg: go/github.io/2zyyyyy/chineseDocumentation/unitTesting/split
BenchmarkSplit
BenchmarkSplit-8        11468383                99.6 ns/op            32 B/op          1 allocs/op
PASS
ok
```

这个使用make函数提前分配内存的改动，减少了一般的内存分配次数，并且减少了1/3的内存分配。

**性能比较函数**

上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。

性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。

**重置时间**

b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。例如：

```go
func BenchmarkSplit(b *testing.B) {
    time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
    b.ResetTimer()              // 重置计时器
    for i := 0; i < b.N; i++ {
        Split("枯藤老树昏鸦", "老")
    }
} 
```

**并行测试**

func (b *B) RunParallel(body func(*PB))会以并行的方式执行给定的基准测试。

RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。RunParallel通常会与-cpu标志一同使用。

```go
func BenchmarkSplitParallel(b *testing.B) {
    // b.SetParallelism(1) // 设置使用的CPU数
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            Split("枯藤老树昏鸦", "老")
        }
    })
}  
```

执行一下基准测试：

```
split $ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/pprof/studygo/code_demo/test_demo/split
BenchmarkSplit-8                10000000               131 ns/op
BenchmarkSplitParallel-8        50000000                36.1 ns/op
PASS
ok      github.com/pprof/studygo/code_demo/test_demo/split       3.308s
```

还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。

**Setup与TearDown**

测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。

**TestMain**

通过在`*_test.go`文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。

如果测试文件包含函数:`func TestMain(m *testing.M)`那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和拆卸（teardown）。退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit。

一个使用TestMain来设置Setup和TearDown的示例如下：

```go
func TestMain(m *testing.M) {
    fmt.Println("write setup code here...") // 测试之前的做一些设置
    // 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
    retCode := m.Run()                         // 执行测试
    fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
    os.Exit(retCode)                           // 退出测试
}  
```

需要注意的是：在调用TestMain时, flag.Parse并没有被调用。所以如果TestMain 依赖于command-line标志 (包括 testing 包的标记), 则应该显示的调用flag.Parse。

**子测试的Setup与Teardown**

有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。下面我们定义两个函数工具函数如下：

```go
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
    t.Log("如有需要在此执行:测试之前的setup")
    return func(t *testing.T) {
        t.Log("如有需要在此执行:测试之后的teardown")
    }
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
    t.Log("如有需要在此执行:子测试之前的setup")
    return func(t *testing.T) {
        t.Log("如有需要在此执行:子测试之后的teardown")
    }
}  
```

使用方式如下：

```go
func TestSplit(t *testing.T) {
    type test struct { // 定义test结构体
        input string
        sep   string
        want  []string
    }
    tests := map[string]test{ // 测试用例使用map存储
        "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
        "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
        "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
        "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"", "枯藤", "树昏鸦"}},
    }
    teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
    defer teardownTestCase(t)            // 测试之后执行testdoen操作

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
            teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
            defer teardownSubTest(t)           // 测试之后执行testdoen操作
            got := Split(tc.input, tc.sep)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("excepted:%#v, got:%#v", tc.want, got)
            }
        })
    }
} 
```

测试结果如下：

```
    split $ go test -v
    === RUN   TestSplit
    === RUN   TestSplit/simple
    === RUN   TestSplit/wrong_sep
    === RUN   TestSplit/more_sep
    === RUN   TestSplit/leading_sep
    --- PASS: TestSplit (0.00s)
        split_test.go:71: 如有需要在此执行:测试之前的setup
        --- PASS: TestSplit/simple (0.00s)
            split_test.go:79: 如有需要在此执行:子测试之前的setup
            split_test.go:81: 如有需要在此执行:子测试之后的teardown
        --- PASS: TestSplit/wrong_sep (0.00s)
            split_test.go:79: 如有需要在此执行:子测试之前的setup
            split_test.go:81: 如有需要在此执行:子测试之后的teardown
        --- PASS: TestSplit/more_sep (0.00s)
            split_test.go:79: 如有需要在此执行:子测试之前的setup
            split_test.go:81: 如有需要在此执行:子测试之后的teardown
        --- PASS: TestSplit/leading_sep (0.00s)
            split_test.go:79: 如有需要在此执行:子测试之前的setup
            split_test.go:81: 如有需要在此执行:子测试之后的teardown
        split_test.go:73: 如有需要在此执行:测试之后的teardown
    === RUN   ExampleSplit
    --- PASS: ExampleSplit (0.00s)
    PASS
    ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s 
```

**示例函数**

\### 示例函数的格式

被go test特殊对待的第三种函数就是示例函数，它们的函数名以Example为前缀。它们既没有参数也没有返回值。标准格式如下：

```go
func ExampleName() {
    // ...
} 
```

### 示例函数示例

下面的代码是我们为Split函数编写的一个示例函数：

```go
func ExampleSplit() {
    fmt.Println(split.Split("a:b:c", ":"))
    fmt.Println(split.Split("枯藤老树昏鸦", "老"))
    // Output:
    // [a b c]
    // [ 枯藤 树昏鸦]
} 
```

为你的代码编写示例代码有如下三个用处：

```
    示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。

    示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。

        split $ go test -run Example
        PASS
        ok      github.com/pprof/studygo/code_demo/test_demo/split       0.006s
    示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。下图为strings.ToUpper函数在Playground的示例函数效果。
```

![null](https://tva1.sinaimg.cn/large/008i3skNly1gyp1wccxbej31e40qkdib.jpg)

#### 9、压力测试

PASS

### 方法

#### 1、方法定义

Golang方法总是绑定对象实例，并隐式将实例作为第一实参（receiver）。

```go
1.只能为当前包内命名类型定义方法
2.参数receiver可任意命名。如方法中未曾使用，可省略参数名
3.参数receiver类型可以是T或者*T类型
4.不支持方法重载，receiver只是参数签名的组成部分
5.可用实例value或pointer调用全部方法，编译器自动转换
```

一个方法就是一个包含了接收者的函数，接收者可以是命名类型或者结构体类型的一个值或者是一个指针。

所有给定类型的方法属于该类型的方法集。

**方法定义**

```go
func (receiver type) methodName(参数列表)(返回值列表){
  // 参数和返回值可以省略
}
```

```go
type Test struct {}

// 无参数和返回值
func (t Test) metnod01() {}

// 单参数 无返回值
func (t Test) method02(i int) {}

// 多参数 无返回值
func (t Test) metnod03(i, j int) {}

// 无参数 单返回值
func (t Test) method04()(i int) {}

// 多参数 多返回值
func (t Test) method05(i, j int) (x, y int, err error) {}

// 指针类型
// 无参数和返回值
func (t *Test) metnod01() {}

// 单参数 无返回值
func (t *Test) method02(i int) {}

// 多参数 无返回值
func (t *Test) metnod03(i, j int) {}

// 无参数 单返回值
func (t *Test) method04()(i int) {}

// 多参数 多返回值
func (t *Test) method05(i, j int) (x, y int, err error) {}

func main() {
  // ...
}
```

下面定义一个结构体类型和该类型的一个方法

```go
package main

import "fmt"

// Golang 方法

// struct
type User struct {
	Name, Address string
}

// struct metnhod
func (u User) Express(num string) {
	fmt.Printf("u.num=%s, u.name=%s, u.address=%s\n", num, u.Name, u.Address)
}

func main() {
	// 值类型调用方法
	u1 := User{
		"张三",
		"法外狂徒张三的家在哪里？",
	}
	fmt.Printf("u1 type=%T\n", u1)
	u1.Express("88798871")

	// 指针类型调用方法
	u2 := &User{
		"李四",
		"法外狂徒李四的家在哪里？",
	}
	fmt.Printf("u2 type=%T\n", u2)
	u2.Express("15099893012")
}

// 输出
u1 type=main.User
u.num=88798871, u.name=张三, u.address=法外狂徒张三的家在哪里？
u2 type=*main.User
u.num=15099893012, u.name=李四, u.address=法外狂徒李四的家在哪里？
```

首先我们定义了一个叫做User的结构体类型，然后定义了一个该类型的方法叫做Express，该方法的接收者是一个User类型的值。要调用Express方法我们需要一个User类型的值或者指针。

在这个例子中当我们使用指针时，Go调用和解引用指针是的调用可以被执行。注意，当接收者不是一个指针时，该方法操作对应接收者的值的副本（意思就是即使你使用了指针调用函数，但是函数的接收者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作）。

修改Express方法，让它的接收者使用指针类型：

```go
// struct metnhod
func (u *User) Express(num string) {
	fmt.Printf("u.num=%s, u.name=%s, u.address=%s\n", num, u.Name, u.Address)
}
```

注意：当接收者是指针时，即使用值类型调用，那么函数内部也是对指针的操作。

方法不过是一种特殊的函数，只需将其还原，就知道receiver T和*T的差别。

```go
type Data struct {
	x int
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("&d=%p\n", p) // 0xc0000b4008

	d.vauleTest()   // 0xc0000b4018
	d.pointerTest() // 0xc0000b4008

	p.vauleTest()   // 0xc0000b4030
	p.pointerTest() // 0xc0000b4008
}
```

**普通函数与方法的区别**

1. 对于普通函数，接收者为值类型时，不能将指针类型数据直接传递，反之亦然。
2. 对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来也同样可以。

```go
// 普通函数与方法的区别(在接收者分别为值类型和指针类型的时候)

// 1.普通函数
// 接收值类型参数的函数
func valueTest(a int) int {
  return a + 10
}

// 接收指针类型参数的函数
func pointTest(a *int) int {
  return *a + 10
}

func structTestValue() {
  a := 2
  fmt.Println("valueTest:", valueTest(a))
  // 函数的参数作为值类型，则不能直接将指针作为参数传递
  fmt.Println("valueTest:", valueTest(&a)) // 错误写法
  
  b := 5
  fmt.Println("pointerTest:", pointTest(&b))
  // 同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
  fmt.Println("pointTest:", pointTest(b)) // 错误写法
}

// 2.方法
type User struct {
  id int
  name string
}

// 接收者为值类型
func (u User) valueShowName() {
  fmt.Println(u.name)
}

// 接收者为指针类型
func (u *User) pointShowName() {
  fmt.Println(u.name)
}

func structTestFunc() {
  // 值类型调用方法
  userValue := User{1, "张三"}
  userValue.valueShowName()
  userValue.pointShowName()
  
  // 指针类型调用方法
  userPoint := &User{2, "李四"}
  userPoint.valueShowName()
  userPoint.pointShowName()
  //与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}

func main() {
  structTestValue()
	structTestFunc()
}

// 输出
valueTest: 12
pointerTest: 15
张三
张三
李四
李四
```

#### 2、匿名字段

Golang匿名字段：可以像字段成员那样访问匿名字段方法，编译器负责查找。

```go
package main

import (
	"fmt"
)

// 匿名字段

type User struct {
	id   int
	name string
}

type Manager struct {
	User
}

func (u *User) toString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

func main() {
	manager := Manager{User{
		100,
		"法外狂徒张三",
	}}
	// 反射获取manager类型
	// fmt.Println("manager type=", reflect.TypeOf(manager))
	fmt.Printf("manager type=%T, manager=%p\n", manager, &manager)
	fmt.Println(manager.toString())
}

// 输出
manager type=main.Manager, manager=0xc00000c060
User: 0xc00000c060, &{100 法外狂徒张三}
```

通过匿名字段，可获得和继承类似的复用能力。依据编译器查找顺序，只需在外层定义同名方法，就可以实现”override“。

```go
package main

import (
	"fmt"
)

// 匿名字段

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (u *User) toString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

// 通过匿名字段，可获得和继承类似的复用能力。依据编译器查找顺序，只需在外层定义同名方法，就可以实现”override“。
func (m *Manager) toString() string {
	return fmt.Sprintf("Manager: %p, %v", m, m)
}

func main() {
	manager := Manager{User{
		100,
		"法外狂徒张三",
	}, "manager title"}

	fmt.Printf("manager type=%T, manager=%p\n", manager, &manager)
	fmt.Println(manager.toString())
	fmt.Println(manager.User.toString())
}

//  输出
manager type=main.Manager, manager=0xc000064180
Manager: 0xc000064180, &{{100 法外狂徒张三} manager title}
User: 0xc000064180, &{100 法外狂徒张三}
```

#### 3、方法集

Golang方法集：每个类型都有与之关联的方法集，这会影响到接口实现规则。

```go
1.类型 T 方法集包含全部receiver T 方法
2.类型 *T 方法集包含全部receiver T + *T 方法
3.如类型 S 包含匿名字段 T，则S和*S方法集包含 T 和 *T	
4.如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法
5.不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法
```

用实例value和pointer调用方法（含匿名字段）不受方法集约束，编译器总是查找全部方法，并自动转换receiver实参。

Go语言中内部类型方法集提升的规则：

类型 T 方法集包含了全部receiver T 方法。

```GO
package main

import "fmt"

// 方法集

type T struct {
	int
}

func (t T) tFunc() {
	fmt.Println("类型T方法集包含所有receiver T的方法.")
}

func main() {
	t1 := T{
		100,
	}
	fmt.Printf("t1=%v\n", t1)
	t1.tFunc()
}

// 输出
t1={100}
类型T方法集包含所有receiver T的方法.
```

类型 `*T` 方法集包含全部 `receiver T + *T` 方法。

```go
package main

import "fmt"

// 方法集

type T struct {
	int
}

func (t T) tFunc() {
	fmt.Println("类型T方法集包含所有receiver T的方法.")
}

func (t *T) pFunc() {
	fmt.Println("类型*T方法集包含所有receiver *T的方法.")
}

func main() {
	t1 := T{
		100,
	}
	t2 := &t1
	fmt.Printf("t2=%v\n", t2)
	t2.tFunc()
	t2.pFunc()
}

// 输出
t2=&{100}
类型T方法集包含所有receiver T的方法.
类型*T方法集包含所有receiver *T的方法.
```

给定一个结构体类型S和命名为T的类型，方法提升像下面规定的这样被包含在结构体方法集中：

如类型S包含匿名字段T，则S和*S方法集包含T的方法。

这条规则说的是当我们嵌入一个类型，嵌入类型的接收者为值类型的方法将被提升，可以被外部类型的值和指针调用。

```GO
type T struct {
  int
}

type S struct {
  T
}

func (t T) tFunc() {
  fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
}

func main() {
  s1 := S{T{100}}
  s2 := &s1
  fmt.Printf("s1=%v\n", s1)
  s1.tFunc()
  fmt.Printf("s2=%v\n", s2)
  s2.tFunc()
}

// 输出
s1={{100}}
如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
s2=&{{100}}
如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
```

如类型 S 包含匿名字段 `*T`，则 S 和 `*S` 方法集包含 `T + *T` 方法。

这条规则说的是当我们嵌入一个类型的指针，嵌入类型的接收者为值类型或指针类型的方法将被提升，可以被外部类型的值或者指针调用。

```go
type T struct {
  int
}

type S struct {
  T
}

func (t T) tFunc() {
  fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法")
}

func (t *T) pFunc() {
  fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法")
}

func main() {
  s1 := S{T{100}}
  s2 := &s1
  fmt.Printf("s1 is : %v\n", s1)
  s1.tFunc()
  s1.pFunc()
  fmt.Printf("s2 is : %v\n", s2)
  s2.tFunc()
  s2.pFunc()
}

// 输出
s1 is : {{100}}
如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法
s2 is : &{{100}}
如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法
如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法
```

### 4、表达式

Golang表达式：根据调用者不同，方法分为两种表现形式：

```go
instance.method(args...) ---> <type>.func(instance, args...)
```

前者称为method value，后者method expression。

两者都可像普通函数那样赋值和传参，区别在于method value绑定实例，二method expression则须显示传参。

```go
package main

import (
	"fmt"
)

// 表达式

type User struct {
	id   int
	name string
}

func (u *User) Test() {
	fmt.Printf("u.p=%p, u.v=%v\n", u, u)
}

func main() {
	user := User{
		100,
		"user.name",
	}
	user.Test()

	methodValue := user.Test
	methodValue() // 隐式传递 receiver

	methodExpression := (*User).Test
	methodExpression(&user) // 显式传递 receiver
}

// 输出
u.p=0xc00000c060, u.v=&{100 user.name}
u.p=0xc00000c060, u.v=&{100 user.name}
u.p=0xc00000c060, u.v=&{100 user.name}
```

需要注意，method value会复制receiver。

```GO
type User struct {
  id int
  name string
}

func (u User) Test() {
  fmt.Println(u)
}

func main() {
  user := User{
    1,
    "user.name",
  }
  methodValue := user.Test // 立即复制receiver，因为不是指针类型，不受后续修改影响
  
  user.id, user.name = 2, "tony"
  user.Test()
  
  methodValue()
}

// 输出
{2 tony}
{1 user.name}
```

#### 5、自定义error

**抛异常和处理异常**

****

**系统抛异常**

```GO
package main

import "fmt"

// 自定义异常

// system error
func systemError() {
	a := [5]int{1, 2, 3, 4, 5}
	a[1] = 123
	fmt.Println(a)
	index := 10
	a[index] = 10
	fmt.Println(a)
}

// 自己抛
func getCircleArea(radius float32) (area float32) {
	if radius <= 0 {
		panic("半径必须大于0")
	}
	return 3.14 * radius * radius
}

func test01() {
	// 延时执行匿名函数
	// 延时到何时 1.程序正常结束 2.发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Printf("defer fun() err=%s\n", err)
		}
	}()
	getCircleArea(-1)
	fmt.Println("getCircleArea函数报错，此处不执行。")
}

func test02() {
	test01()
	fmt.Println("test02()")
}

func main() {
	test02()
}

// 输出
defer fun() err=半径必须大于0
test02()
```

**返回异常**

```go
package main

import (
	"errors"
	"fmt"
)

// 返回异常
func getCircleArea02(radius float32) (area float32, err error) {
	if radius <= 0 {
		// 构建一个异常对象
		err = errors.New("半径必须大于0")
		return
	}
	area = 3.14 * radius * radius
	return
}

func main() {
	area, err := getCircleArea02(0)
	if err != nil {
		fmt.Printf("err=%s\n", err)
	} else {
		fmt.Printf("area=%f\n", area)
	}
}

// 输出
err=半径必须大于0
```

**自定义error**

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// 自定义异常
type CustomError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("path=%s \n op=%s \n createTime=%s message=%s",
		c.path, c.op, c.createTime, c.message)
}

func Open(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &CustomError{
			path:       fileName,
			op:         "read",
			createTime: fmt.Sprintf("%v\n", time.Now()),
			message:    err.Error(),
		}
	}
	defer file.Close()
	return nil
}

func main() {
	//自定义error
	err = Open("/Users/gilbert/go/src/go/README1.md")
	switch v := err.(type) {
	case *CustomError:
		fmt.Println("get path error,", v)
	default:
	}
}

// 输出
get path error, path=/Users/gilbert/go/src/go/README1.md 
 op=read 
 createTime=2022-01-27 18:26:49.228543 +0800 CST m=+0.000137326
 message=open /Users/gilbert/go/src/go/README1.md: no such file or directory
```

### 面向对象

#### 1、匿名字段

Golang支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

```GO
package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id      int
	address string
}

func main() {
	// 初始化
	student1 := Student{Person{"柱子哥1号", "man", 28}, 1001, "海智中心1号楼"}
	fmt.Println(student1)

	student2 := Student{
		Person: Person{"柱子哥2号", "man", 29}}
	fmt.Println(student2)

	student3 := Student{Person: Person{name: "柱子哥3号"}}
	fmt.Println(student3)
}

// 输出
{{柱子哥1号 man 28} 1001 海智中心1号楼}
{{柱子哥2号 man 29} 0 }
{{柱子哥3号  0} 0 }
```

**同名字段的情况**

```go
type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id      int
	address string
  // 同名字段
  name string
}

func main() {
  var student Student
  // 给自己的字段赋值
  student.name = "Student.name"
  
  // 给父类同名字段fuzhi(Person.name)
  student.Person.name = "Person.name"
  fmt.Println(student)
}
```

**所有的内置类型和自定义类型都是可以作为匿名字段去使用**

```GO
type Person struct {
	name string
	sex  string
	age  int
}

// 自定义类型
type myString string

type Student struct {
	Person
	id      int
	myString
}

func main() {
  student := Student{Person{"person.name", "person.sex", 10}, 100, "student.myString"}
  fmt.Println(student)
}
```

**指针类型匿名字段**

```GO

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	*Person
	id      int
	address string
}

func main() {
  student := Student{&Person{"大柱子", "man", 28}, 100, "钉钉空间3号楼403"}
  fmt.Println(student)
  fmt.Println(student.name)
  fmt.Println(student.Person.name)
}

// 输出
{0xc000090180 100 钉钉空间3号楼403}
大柱子
大柱子
```

#### 2、接口

接口（interface）定义了一个对象的行为规范，只定义规范不实现,由具体的对象来实现规范的细节。

**接口**

****

在Go语言中接口（interface）是一种类型，一种抽象的类型。

​	interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我们就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

​	为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。

**为什么要使用接口**

```go
type Cat struct{}

func (c Cat) Say() string {
  return "喵喵喵"
}

type Dog struct{}

func (d Dog) Say() string {
  return "汪汪汪"
} 

func main() {
  c := Cat{}
  fmt.Pritntln("猫：", c.Say())
  
  d := Dog{}
  fmt.Pritntln("狗：", d.Say())
}
```

上面代码中定义了猫和狗两种类型，都具有Say方法，但是main函数中明显有重复代码，如果我们再加上其他动物的话，我们的代码就会一直重复下去。那么我们能不能把他们当成“能叫的动物”来处理呢？

Go语言未科解决类似上面的问题，就设计了接口这个概念，接口区别于我们之前所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道他是什么，唯一知道的事通过他的方法能做什么。

**接口定义**

Go语言提倡面向接口编程。

```go
接口是一个或多个方法签名的集合。
任何类型的方法集中只要拥有该接口"对应的全部方法"签名，就表示它"实现"了该接口,无须在该类型上显示声明实现了哪个接口。这称为structural typing。所谓对应方法，是指有相同名称、参数列表(不包括参数名)以及返回值。当然，该类型还可以有其他方法.

接口只有方法声明，没有实现和数据字段。
接口可以匿名嵌入其他接口，或嵌入到结构体中。
对象赋值给接口时，就会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
只有当接口存储的类型和对象都为nil时，接口才等于nil。
接口调用不会做receiver的自动转换。
接口同样支持匿名字段方法。
接口也可以实现类似oop中的多态。
空接口可以作为任何类型数据的容器。
一个类型可以实现多个接口。
接口命名习惯以er结尾。
```

每个接口由数个方法组成，接口的定义格式如下：

```GO
type 接口类型名 interface{
  方法名1(参数列表1) 返回值列表1
  方法名2(参数列表2) 返回值列表2
  ……
}
```

其中：

```GO
1.接口名：
	使用type将接口定义为自定义的类型名。Go的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer。接口命名做好要能突出该接口的类型含义。
2.方法名：
当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包(package)之外的代码访问。
3.参数、返回值列表：
	参数列表和返回值列表中的参数变量名可以省略。
```

举例：

```GO
type Writer interface{
  Write([]byte) error
}
```

当你看到这个接口类型的值时，你不知道他是什么，唯一知道的就是可以通过它的Write方法来做一些事情。

**实现接口的条件**

一个对象只要实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。

我们来定义一个Sayer接口：

```GO
// Sayer接口
type Sayer interface {
  say()
}

// 定义cat和dog两个结构体
type Cat struct {}

type Dog struct {}
```

因为Sayer接口中只有一个say方法，所以我们只需要给dog和cat分别实现say方法就可以实现Sayer接口了。

```GO
// dog实现Sayer接口
func (d Dog) say {
  fmt.Println("汪汪汪~")
}

// cat实现Sayer接口
func (c Cat) say {
  fmt.Println("喵喵喵~")
}
```

接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。

**接口类型变量**

实现了接口有什么用？

接口类型变量能够存储所有实现了该接口的实例。例如上面的实例中，Sayer类型的变量能够存储dog和cat类型的变量。

```GO
func main() {
	// 声明Sayer类型变量
	var sayer Sayer
	// sayer.say()  // panic: runtime error: invalid memory address or nil pointer dereference
	// 实例化dog和cat
	cat := Cat{}
	dog := Dog{}

	sayer = cat // 将cat实例赋值给sayer
	sayer.say() // 喵喵喵

	sayer = dog  // 将dog实例赋值给sayer
	sayer.say() // 汪汪汪
}

// 输出
喵喵喵
汪汪汪
```

**值接收者和指针接收者实现接口的区别**

使用值接收者实现接口和使用指针接收者实现接口有什么区别呢？下面我们通过一个例子看一下其中的区别。

我们有一个Move接口和一个Dog结构体。

```GO
type Mover interface {
	move()
}

type Dog struct{}
```

**值接收者实现接口**

```go
// 值接收者实现接口
func (d Dog) move() {
	fmt.Println("dog会动~")
}
```

此时实现接口的是Dog类型：

```GO
func main() {
  var move Mover
	var dogValue = Dog{} // Dog类型
	move = dogValue      // move 可以接收Dog类型
	move.move()

	var dogPointer = &Dog{} // *Dog类型
	move = dogPointer       // move 可以接收*Dog类型
	move.move()
}
```

从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是Dog结构体还是结构体指针`*Dog`类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，Dog指针dogPointer内部会自动求值`*dogPointer`.

**指针接收者实现接口**

同样的代码我们再来测试一下使用指针接收者有什么区别：

```go
func (d *Dog) move() {
  fmt.Println("汪汪汪")
}

func main() {
  var move Mover
  var dogValue Dog{}
  move = dogValue // move不可接收dog类型
  
  var dogPointer &Dog{}  // dogPointer是*Dog类型
  move = dogPointer  // move可以接收*Dog类型
}
```

此时实现Mover接口的是*Dog类型，所以不能给move传Dog类型的dogValue，此时move只能存储 ` *Dog `类型的值。

**下面的代码是一个比较好的面试题**

请问下面的代码是否能通过编译？

```GO
type People interface {
  Speak(string) string
}

type Student struct {}

func (s *Student) Speak(think string) (talk string) {
	if think == "dsb" {
		talk = "大帅币"
	} else {
		talk = "小水币"
	}
	return
}

func main() {
  // 面试题
	var people People = Student{}
	think := "the"
	people.Speak(think)
}

// IDE在var people People = Student{}报错，因为实现People接口的*Studnet类型，使用Student类型无法编译
运行输出：
cannot use Student literal (type Student) as type People in assignment:
        Student does not implement People (Speak method has pointer receiver)

// 修改people类型
var people People = &Student{}

// go run main.go
小水币
```

**类型与接口的关系**

****

**一个类型实现多个接口**

一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。例如：狗可以叫，也可以跑。我们就分别定义Sayer和Runner接口，如下：

```go
package main

import "fmt"

// 类型与接口

// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
// 例如：狗可以叫，也可以跑。我们就分别定义Sayer和Runner接口

type Sayer interface {
	say()
}

type Runner interface {
	run()
}

type Dog struct {
	name string
}

// dog 可以同时实现以上两个接口

// 实现Sayer接口
func (d Dog) say() {
	fmt.Printf("%s会叫~\n", d.name)
}

// 实现Runner接口
func (d Dog) run() {
	fmt.Printf("%s会跑~\n", d.name)
}

func main() {
	var sayer Sayer
	var runner Runner

	dog := Dog{
		"阿黄",
	}
	sayer, runner = dog, dog
	sayer.say()
	runner.run()
}

// 输出
阿黄会叫~
阿黄会跑~
```

**多个类型实现同一接口**

Go语言中不同的类型还可以实现同一接口。如下：

```GO
package main

import "fmt"

type Runner interface {
	run()
}

type Dog struct {
	name string
}

// 实现Runner接口
func (d Dog) run() {
	fmt.Printf("%s会跑~\n", d.name)
}

// 多个类型实现同一接口
// 狗会跑 猫咪也会跑 可以同时实现Runner接口
type Cat struct {
	bread string
}

func (c Cat) run() {
	fmt.Printf("%s也会跑~\n", c.bread)
}

func main() {
	var runner Runner

	dog := Dog{
		"阿黄",
	}
	cat := Cat{
		"加菲猫",
	}
	runner = dog
	runner.run()

	runner = cat
	runner.run()
}

// 输出
阿黄会跑~
加菲猫也会跑~
```

并且一个接口的方法，不一定由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

```GO
// WashingMachine 洗衣机
type WashingMachine interface {
  wash()
  dry()
}

// 甩干器
type dryer struct {}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
  fmt.Println("甩一甩~")
}

// 海尔洗衣机
type haier struct {
  // 嵌入甩干器
  dryer
}

// 实现WashingMachine的wash()方法
func (h haier) wash() {
  fmt.Println("洗洗更健康~")
}
```

**接口嵌套**

接口与接口之间可以通过嵌套创造出新的接口。

```GO
// Sayer
type Sayer interface {
  say()
}

// Mover
type Mover interface {
  move()
}

// 接口嵌套
type Animal interface {
  Sayer
  Mover
}

func main() {
  animal := Monkey{"猴哥"}
	animal.say()
	animal.run()
}

// 输出
猴哥会说话~
猴哥会跑步~
```

**空接口**

****

**空接口的定义**

空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

空接口类型的变量可以存储任意类型的变量。

```GO
func main() {
  // 空接口
	var x interface{}
	str := "2zyyyyy.com"
	x = str
	fmt.Printf("x type = %T, x value = %s\n", x, x)

	num := 100
	x = num
	fmt.Printf("x type = %T, x value = %d\n", x, x)

	bool := true
	x = bool
	fmt.Printf("x type = %T, x value = %v\n", x, x)
}

// 输出
x type = string, x value = 2zyyyyy.com
x type = int, x value = 100
x type = bool, x value = true
```

**空接口的应用**

**空接口作为函数的参数**

使用空接口实现可以接收任意类型的函数参数。

```GO
// 空接口作为函数参数
func test(a interface{}) {
  fmt.Printf("type:%T value:%v\n", a, a)
}
```

空接口作为map的值

使用空接口实现可以保存任意值的字典。

```GO
// 空接口作为map值
studentInfo := make(map[string]interface{})
studentInfo["one"] = "中文"
studentInfo["two"] = 100
studentInfo["three"] = true
fmt.Println(studentInfo)
```

**类型断言**

空接口可以存储任意类型的值，那我们如何获取其存储的而具体数据呢？

**接口值**

一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。

举例：

```GO
var w io.Wtirer

w = os.Stdout
w= new(byets.Buffer)
w = nil
```

图解：

![img](https://tva1.sinaimg.cn/large/008i3skNly1gz1g6ju1gxj310e0tc76d.jpg)

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

```GO
x.(T)
```

其中：

```GO
X：表示类型为interface{}的变量
T：表示断言x可能是的类型
```

该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

举例：

```GO
func typeAssert(s interface{}) (info string) {
	info, ok := s.(string)
	if ok {
		fmt.Printf("断言成功, info:%s\n", info)
	} else {
		fmt.Println("断言失败")
	}
	return
}

func main() {
  typeAssert("123456789")
}

// 输出
断言成功, info:123456789
```

以上事例中如果需要断言多次则需要写多个if判断，这个时候我们可以使用switch语句来实现：

```GO
func typeAssertSwitch(s interface{}) (info string) {
	switch v := s.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("type is unsupport")
	}
	return
}

func main() {
  typeAssertSwitch(false)
}

// 输出
x is a bool is false
```

因为空接口可以存储任意类型的值的特点，所以空接口在Go语言中的使用十分广泛。

关于接口需要注意的是，只有当有两个或两个以上具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时消耗。

### 网络编程

#### 1、互联网协议介绍

互联网的核心是一系列协议，总称为“互联网协议”（Internet  protocol suite），正是这一些协议规定了电脑如何连接和组网。我们理解了这些协议，就理解了互联网的原理。由于这些协议太过庞大和复杂，没有办法在这里一概而全，只能介绍一下我们日常开发中接触较多的几个协议。

**互联网分层模型**

互联网的逻辑实现被分为好几层。每一层都有自己的功能，就像建筑物一样，每一层要靠下一层支持。用户接触的只是最上面的那一层，根本不会感觉到下面的基层。要理解互联网就需要自下而上理解每一层实现的功能。

![img](https://tva1.sinaimg.cn/large/008i3skNly1gz1ig2ukffj317k0u0q68.jpg)

如上图所示，互联网按照不同的模型划分会有不同的分层，但是不论按照什么模型区划分，越往上的层越靠近用户，越往下的层越靠近硬件。在软件开发中我们使用最多的是上图中将互联网划分为5个分层的模型。

**物理层**

我们的电脑要与外界互联网通信，需要先把电脑连接网络，我们可以用双绞线、光纤、无线电波等方式。这就叫做“实物理层”，他就是把电脑连接起来的物理手段。它主要规定了网络的一些电气特性，作用是负责传送0和1的电信号。

**数据链路层**

单纯的0和1没有任何意义，所以我们使用者会为其赋予一些特定的含义，规定解读电信号的方式；例如：多少电信号算一组？每个信号位有何意义？这就是“数据链路层”的功能，它在“物理层”的上方，确定了物理层传输的0和1分组方式以及代表的意义。早期的时候，每家公司都有自己的电信号分组方式。逐渐地，一种叫“以太网”（Ethernet）的协议，占据了主导地位。

以太网规定，一组电信号构成一个数据包。叫做“帧”（frame）。每一帧分成两个部分：标头（head）和数据（data）。其中标头包含数据包的一些说明项，比如发送者、接受者、数据类型等等；数据则是数据包的具体内容。标头的长度，固定为18字节。数据的长度，最短为46字节，最长为1500字节。因此，整个“帧”最短为64字节，最长为1518字节。如果数据很长，就必须分割成多个帧进行发送。

那么，发送者和接收者时如何标识呢？以太网规定，连入网络的所有设备都必须具有“网卡”接口。数据包必须是从一块网卡传送到另外一块网卡。网卡的地址,就是数据包的发送地址和接收地址，这叫做MAC地址。每块网卡出厂的时候，都有一个全世界独一无二的MAC地址，昌都市48个二进制位，通常用12个十六进制数表示。前六个十六进制数是厂商编号，后六个是该厂商的网卡流水号。有了MAC地址，就可以定位网卡和数据包的路径了。

我们会通过ARP协议来获取接受方的MAC地址，有了MAC地址之后，如何把数据准确的发送给接收方呢？其实这里以太网采用了一种很“原始”的方式，他不是把数据包准备送到接收方，而是向本网络内所有计算机都发送，让每台计算机读取到这个包的 标头，找到接收方的MAC地址，然后与自身的MAC地址比较，如果两者相同，就接受这个包，作进一步处理，否则就丢弃这个包。这种发送方式就叫做“广播”（broadcasting）。

**网络层**

按照以太网协议的规则我们可以依靠MAC地址来向外发送数据。理论上依靠MAC地址，你电脑的网卡就可以找到身在时间另一个角落的网卡了，但是这种做法有一个重大缺陷就是以太网采用广播方式发送数据包，所有成员人寿一“包”，不仅效率低，而且发送的数据只能局限在发送者所在的子网络。也就是说如果两台计算机不在一个子网络，广播是传不过去的。这种设计是合理且必要的，因为如果互联网上每一台计算机都会接收到互联网上发送的所有数据包，那是不现实的。

因此，必须找到一种方法区分哪些MAC地址属于同一个子网络，那些不是。如果是同一个子网络，就采用广播方式发送，否则就采用“路由”方式发送。这就导致了“网络层”的诞生。它的作用是引进一套新的地址，使得我们能够区分不同的计算机是否属于同一个自我网络。这套地址就叫做“网址”。

“网络层”出现以后，每台计算机有了两种地址，一种是MAC地址，另一种是网络地址。两种地址之间没有任何联系，MAC地址是绑定在网卡上的，网络地址则是网络管理员分配的。网络地址帮助我们确定计算机所在的子网络，MAC地址则将数据包发送到该子网络中的目标网卡。因此，从逻辑上可以推断，必定是先处理网络地址，然后在处理MAC地址。

规定网络地址的协议，叫做IP协议。他所定义的地址，就被称为IP地址。目前，广泛采用的是IP协议第四版，简称IPv4。IPv4这个版本规定，网络地址由32个二进制位组成，我们通常习惯用分成四段的十进制数表示IP地址，从0.0.0.0一直到255.255.255.255。

根据IP序偶诶哦发送的数据，就叫做IP数据包。IP数据包也分为标头和数据两个部分。标头部分主要包括版本、长度、IP地址等信息。数据部分则是IP数据包的具体内容。IP数据包的标头部分的长度为20到60字节，整个数据包的总长度最大为65535字节。

**传输层**

有了MAC地址和IP协议，我们已经可以在互联网上任意两台主机上建立通信。但问题是同一台主机上会有许多程序都需要用网络收发数据，比如QQ和浏览器这两个程序都需要连接互联网并收发数据，我们如何区分某个数据包到底是归属那个程序的呢？也就是说，我们还需要一个参数，表示这个数据包到底供哪个程序（进程）使用。这个参数就叫做“端口”（port），他其实是每一个使用网卡的程序的编号。每个数据包都发送到主机的特定端口，所以不同的程序就能取到自己所需要的数据。

“端口”是0到65535之间的一个整数，正好16个二进制位。0到1023端口被系统占用，用户只能选择大于1023的端口。有了IP和端口我们就能实现唯一确定互联网上的一个程序，劲儿实现网络的程序通信。

我们必须在数据包中加入端口信息，这就需要新的协议。最简单的实现叫做UDP协议，它的格式几乎就是在数据前面，加上端口号。UDP数据包，也是由标头和数据两部分组成，标头部分主要定义了发出端口和接收端口，数据部分就是具体的内容。UDP数据包非常简单，标头部分一共只有8个字节，总长度不超过65535个字节，正好放进一个IP数据包。

UDP协议的优点是比较简单，容易实现，但是缺点是可靠性差，一旦数据包发出，无法知道对方是否收到。为了解决这个问题，提高网络可靠性，TCP协议就诞生了。TCP协议能够确保数据不会丢失。他的缺点是过程复杂、实现困难、消耗较多的资源。TCP数据包没有长度限制，理论上可以无限长，但是为了保证网络的效率，通常TCP数据包的长度不会超过IP数据包的长度，以确保单个TCP数据包不必再分割。

**应用层**

应用程序收到传输层的数据，接下来就要对数据进行解包。由于互联网是开放架构，数据来源五花八门们必须事先规定号通信的数据格式，否则接收方根本无法获得真正发送的数据内容。应用层的作用就是规定应用程序使用的数据格式，例如我们TCP协议之上常见的email、HTTP、FTP等协议，这些协议就组成了互联网协议的应用层。

如下图所示，发送方的HTTP数据进过互联网的传输过程中会一次添加各层协议的标头信息，接收方收到数据包之后再依次根据协议解包得到数据。

![img](https://tva1.sinaimg.cn/large/008i3skNly1gz2qwzamnhj318u0u0tcq.jpg)



#### 2、socket编程

Socket是BSD UNIX的进程通信机制，通常也称作“套接字”，用于描述IP地址和端口，是一个通信链的句柄。Socket可以理解为TCP/IP网络的API，它定义了许多函数或例程，程序猿可以用它们来开发TCP/IP网络上的应用程序。电脑上运行的应用程序通常通过套接字向网络发出请求或者应答网络请求。

**socket图解**

socket是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，socket其实就是一个门面模式，他把复杂的TCP/IP协议族隐藏在socket后面，对用户来说只需要调用socket规定的相关函数，让socket去组织符合指定的协议数据然后进行通信。

![img](https://tva1.sinaimg.cn/large/008i3skNly1gz2rf1k6otj317n0u0q6s.jpg)

- socket又称套接字，应用程序通常通过套接字向网络发出请求或者 应答网络请求
- 常用的socket类型有两种：流式socket和数据包式socket，流式是一种面向连接的socket，针对与面向连接的TCP服务应用，数据包式socket是一种无连接的socket，针对于无连接的UDP服务应用
- TCP：比较靠谱，面向连接，比较慢
- UDP：不是太靠谱，比较快

举个例子：TCP就是货到付款的快递，送到家还必须见到你人才算一整套流程。UDP就像某快递柜一扔就走不管你收到收不到，一般直播用UDP。

**TCP编程**

**Go语言实现TCP通信**

TCP/IP（Transmission Control Protocol/Internet Protocol）即传输控制协议/网际通信，是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（transport layer）通信协议，因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。

**TCP服务端**

一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次连接就创建一个goroutine去处理。

TCP服务端程序的处理流程：

```GO
1.监听端口
2.接收客户端请求建立链接
3.创建goroutine处理链接
```

我们使用Go语言的net包实现的TCP服务端代码如下：

```GO
package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端

// 处理函数
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Print("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动gotoutine处理连接
	}
}
```

**TCP客户端**

一个TCP客户端进行TCP通信的流程如下：

```GO
1.建立与服务端的连接
2.进行数据收发
3.关闭连接
```

使用Go语言的net包实现的TCP客户端如下：

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// client
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入内容
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			//输入q退出
			return
		}
		_, err := conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			fmt.Println("数据发送失败，err:", err)
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("向服务端发送数据:", string(buf[:n]))
	}
}
```

先启动server的main.go（go run main.go），再启动client的main.go。再client端输入任意内容回车回就可以在server端看到对应的数据，从而实现TCP通信。

![client端](https://tva1.sinaimg.cn/large/008i3skNly1gz3xll0m8tj30q40383yz.jpg)

![server端](https://tva1.sinaimg.cn/large/008i3skNly1gz3yelm9ldj30sa01sq38.jpg)

**UDP编程**

**Go语言实现UDP通信**

UDP协议（User Datagram Protocol）中文名称是用户数据协议,是OSI（Open System Interconnection，开放式系统互联 ）参考木星中一种无连接的传输层协议，不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域。

**UDP服务端**

使用Go语言的net包实现的UDP服务端代码如下：

```GO
package main

import (
	"fmt"
	"net"
)

// UDP server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data=%v addr=%v count=%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
```

**UDP客户端**

```GO
package main

import (
	"fmt"
	"net"
)

// UDP client端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务器异常，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败, err:", err)
		return
	}
	fmt.Printf("recv:%v, addr:%v, count:%v\n", string(data[:n]), remoteAddr, n)
}
```

测试：

```GO
// UDP/client/main.go
go run main.go 
recv:Hello server, addr:127.0.0.1:30000, count:12
```

**TCP黏包**

服务端代码：

```GO
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// socket_stick/server/main.go
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据:", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("listen accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
```

客户端代码：

```GO
package main

import (
	"fmt"
	"net"
)

// socket_stick/client/main.go
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "今天初七开工，杭州中雪。"
		conn.Write([]byte(msg))
	}
}
```

先启动server端后在启动client，可以在终端看到客户端向服务端发送的数据。

![image-20220207112708570](https://tva1.sinaimg.cn/large/008i3skNly1gz4scedayxj31rc04c76o.jpg)

客户端分20次发送的数据，在服务端并没有成功输出20次，而是多条数据”粘“在了一起。

**为什么会出现粘包**

主要原因是TCP数据传输是流式，在保持长连接的时候可以进行多次的收和发。”粘包“可发生在发送端也可发生在接收端：

```GO
1.由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等在一小段时间看看在等待期间是否还有要发送的数据，若有则会一次性把这两端数据发送出去。
2.接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层获取数据。当应用层由于某些原因不能及时的把TCP的数据读取出来，就会造成TCP缓冲区中存放了几段数据。
```

**解决方案**

出现粘包情况的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据宝进行封包和拆包的操作。

*封包：封包就是给一段数据加上包头，这样一来数据包就分为了包头和包体两部分内容了(过滤非法包时封包会加入”包尾“内容)。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。*

我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是放的数据的长度。

```GO
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// socket_stick/proto/proto.go

// Encode将消息编码
func Encode(msg string) ([]byte, error) {
	// 读取消息的长度 转换成int32类型（4个字节）
	length := int32(len(msg))
	pkg := new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓存中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), err
}
```

修改原先的server和client代码(socket_stick/server/main.go)：

```go
package main

import (
	"bufio"
	"fmt"
	"go/github.io/2zyyyyy/chineseDocumentation/socket_stick/proto"
	"io"
	"net"
)

// socket_stick/server/main.go
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		fmt.Println("收到client发来的数据:", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("listen accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
```

socket_stick/client/main.go:

```GO
package main

import (
	"fmt"
	"go/github.io/2zyyyyy/chineseDocumentation/socket_stick/proto"
	"net"
)

// socket_stick/client/main.go
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "今天初七开工，杭州中雪。"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("Encode failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
```

![image-20220207172655989](https://tva1.sinaimg.cn/large/008i3skNly1gz52qq2w6ij30j60fstcj.jpg)

#### 3、http编程

**web工作流**

- web服务器的工作原理可以简单地归纳为以下几点
  - 客户机通过TCP/IP协议建立到服务器的TCP连接
  - 客户端向服务器发送HTTP协议请求包，全球服务器里的资源文档
  - 服务器向客户机发送HTTP协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎负责处理”动态内容“，并将处理得到的数据返回给客户端
  - 客户机与服务器断开。由客户端解释HTML文档，在客户端屏幕上渲染图形结果

**HTTP协议**

- 超文本传输协议（HTTP，Hypertext Transfer Protocol）是互联网上应用最为广泛的一种网络协议，它详细规定了浏览器和万维网服务器之间 相互通信的规则，通过因特网传送万维网文档的数据传送协议
- HTTP协议通常承载于TCP协议之上

**HTTP服务端**

```GO
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 单独写回调函数
	http.HandleFunc("/golang", myHandle)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println("listen server failed, err:", err)
		return
	}
}

// handle 函数
func myHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式 get/post/put/delete/update
	fmt.Println("method", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("testInfo!!!"))
}
```

**HTTP客户端**

```GO
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:8000/golang")
	if err != nil {
		fmt.Println("get failed, err:", err)
	}
	defer res.Body.Close()
	// 200 OK
	fmt.Println(res.Status)
	fmt.Println(res.Header)

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := res.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
```

server：

![image-20220208103523722](https://tva1.sinaimg.cn/large/008i3skNly1gz5wgu9x5sj310s04g3yx.jpg)

client：

![image-20220208103621426](https://tva1.sinaimg.cn/large/008i3skNly1gz5wht15o6j313y03874n.jpg)

#### 4、webSocket编程

**websocket简介**

- websocket是一种在单个TCP联=连接上进行全双工通信的协议
- 文websocket使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据
- 在websocket API中，浏览器和服务器只需要完成一次握手，两者之间就直接可以创建持久性的连接，并进行双向数据输出
- 需要安装第三方的包（go get -u -v github.com/gorilla/websocket）

### 并发编程

#### 1、并发介绍

**进程和线程**

```GO
A.进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
B.线程是进程的一个执行实体，是CPU调度和分派的基本单位，它是比继承更小的能独立运行的基本单位。
C.一个进程可以创建和撤销多个线程；同一个进程中的多个线程之间可以并发执行。
```

**并发和并行**

```GO
A.多线程程序在一个核CPU上运行，就是并发。
B.多线程程序在多个核的CPU上运行，就是并行。
```

**协程和线程**

```GO
协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似用户级线程，这些用户级线程的调度也是自己实现的。
线程：一个线程上可以跑多个协程，协程是轻量级的线程。
```

**goroutine只是由官方实现的超级”线程池**“

每个实例<font color=red>`4~5kb`</font>的栈内存占用和由于实现机制而大幅减少的创建和销毁开销是go高并发的根本原因。

**并发不是并行**

并发主要由切换时间片来实现”同时“运行，并行则是直接利用多核实现多线程的运行，go可以设置使用核数，以发挥多核计算机的能力。

*goroutine奉行通过通信来共享内存，而不是共享内存来通行。*

#### 2、goroutine

在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同事需要自己去调度线程执行任务并维护上下文切换，这一切会耗费开发者大量的精力。那么是否有一种机制，开发人员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？

Go余元中的goroutine就是这样一种机制，goroutine的概念类似于线程，但goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将goroutine中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为他在语言层面已经内置了调度和上下文切换的机制。

在Go语言编程中你不需要自己去写进程、线程、协程，你的技能包里只有一个技能-goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了。

**使用goroutine**

Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。

一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。

**启动单个goroutine**

示例：

```GO
package main

import "fmt"

// goroutine

func hello() {
	fmt.Println("hello goroutine!")
}

func main() {
	go hello()
	fmt.Println("hello goroutine done!")
}
```

以上例子如果没有启动一个goroutine去执行hello函数，那么main中的语句是串行的，也就是会先打印hello goroutine！在打印hello goroutine done！。但是在hello函数前面添加go关键字去启动一个新的goroutine执行hello函数，输出的顺序就会相反或者只打印hello goroutine done!。

![image-20220209103152618](https://tva1.sinaimg.cn/large/008i3skNly1gz71ziifv4j310q04at9x.jpg)

**解析**

在程序启动时，Go程序会给main()函数创建一个默认的goroutine。

当main()函数的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束。所以go hello()的goroutine和main()的goroutine谁先结束就决定输出结果，如果后者先结束那么就不会打印hello函数。在main函数中添加强制等待之后两次输出都会执行。

```go
func main() {
	defer time.Sleep(time.Second)
	go hello() // 启动一个goroutine执行hello函数
	fmt.Println("hello goroutine done!")
}
// 输出
hello goroutine done!
hello goroutine!
```

执行上面的代码你会发现，这一次先打印main goroutine done!，然后紧接着打印Hello Goroutine!。

首先为什么会先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。

**启动多个goroutine**

```go
package main

import (
	"fmt"
	"sync"
)

// goroutine

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
```

![image-20220209141204981](https://tva1.sinaimg.cn/large/008i3skNly1gz78clrqrkj308w07ut95.jpg)

多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。

**注意**

如果主协程退出了，其他任务还执行吗？

```GO
func main() {
	// helloGoroutine()
	// 主协程退出其他任务是否还会执行
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

// 输出
main goroutine 1
new goroutine 1
main goroutine 2
new goroutine 2
```

**goroutine与线程**

**可增长的栈**

OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB），一个goroutine的栈在其生命周期开始只有很小的栈（通常2kb），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少情况会用到这个值。所以在Go语言中一次创建十万左右的goroutine也是可以的。

**goroutine调度**

GMP是Go语言运行时（runtime）层面的实现，是Go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

- G：就是个goroutine，里面除了存放本goroutine信息外，还有与所在P的绑定等信息
- P：管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列也消费完了就会去其他P的队列里抢任务。
- M：（machine）是Go运行时（runtime）对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系，一个goroutine最终是要放到M上执行的；

P与M一般也是一一对应的。他们的关系是：P管理着一组挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉的时候回收旧的M。

P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。在并发量打的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

但从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。其一大特点是goroutine的调度是在用户态下完成的，不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池，不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上，再加本身goroutine的超轻量，以上种种保证了go调度方面的性能。

#### 3、runtime包

**runtime.Gosched()**

让出CPU时间片，重新等待安排任务（大概可以理解为本来计划的好好地周末出去烧烤，但是你妈让你去相亲，两种情况第一就是你的相亲速度非常快，见面就黄不耽误你的烧烤计划，第二种情况就是相亲过程非常久，耽误了烧烤计划，但是还想吃，还得去烧烤）

```GO
package main

import (
	"fmt"
	"runtime"
)

// runtime

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("runtime.Gosched~")

	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下 再次分配任务
		runtime.Gosched()
		fmt.Println("主协程")
	}
}
```

**runtime.Goexit()**

退出当前协程

```GO
func main() {
  // runtime.Goexit()
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}
```

**runtime.GOMAXPROCE**

Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核数。例如在一个8核的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCE是m:n调度中的n）。

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU核心数。

Go1.5版本之前，默认是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

我们可以通过将任务分配到不同的CPU逻辑核心数上实现并行的效果，这里举个例子：

```GO
// 将任务分配到不同的CPU逻辑核心上实现并行的效果
func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
  runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
```

Go语言中的操作系统线程和goroutine的关系:

1. 一个操作系统线程对应用户态多个goroutine
2. go程序可以同事使用多个操作线程
3. goroutine和OS线程是多对多的关系，即m:n

#### 4、channel

单纯地将函数鬓发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量堆内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是CSP（communicating sequential processes），提倡通过同心共享内存而不是通过共享内存而实现通信。

Go中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（first in first out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

**channel类型**

channel是一种类型，一种引用类型。其声明通道类型的格式如下：

`var 变量名 chan 元素类型`

举例说明：

```GO
var ch1 chan int  // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔值的通道
var ch3 chan []int  // 声明一个传递int切片的通道
```

**创建channel**

通道是引用类型，通道类型的空值是nil

```go
var ch chan int
fmt.Println(ch) // <nil>
```

声明的通道需要使用make函数初始化之后才能使用。

创建channel的格式如下：

```go
make(chan 元素类型, [缓冲大小])  // 缓冲大小可选

// 举例说明
ch4 := make(chan int)
ch5 := make(chan bool)
ch6 := make(chan []int)
```

**channel操作**

通道有发送（send）、接收（receiver）和关闭（close）三种操作。

发送和接收都使用<-符号。

现在我们先使用以下语句定义一个通道：

```GO
ch := make(chan int)
```

**发送（send）**

将一个值发送到通道中：

```GO
ch <- 100
```

**接收（receiver）**

从一个通道接收值。

```GO
x := <- ch // 从ch中接收值并赋值给x
<- ch // 从ch中接收值，忽略结果
```

**关闭**

我们通过调用内置的close函数来关闭通道。

```GO
close(ch)
```

关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件时必须操作的，但关闭通道不是必须的。

关闭的通道有以下特点：

```GO
1.对一个关闭的通道再发送值就会导致panic
2.对一个关闭的通道进行接收会一直获取值直到通道为空
3.对一个已关闭的并且没有值的通道执行接收操作会得到对应类型的零值
4.关闭一个已经关闭的会导致panic
```

![使用无缓冲通道在goroutine之间同步](https://tva1.sinaimg.cn/large/e6c9d24ely1gzmdn3pddyj20nx0m6767.jpg)

无缓冲通道又称为阻塞的通道。示例：

```GO
func main() {
  ch := make(chan, int)
  ch <- 100
  fmt.Println("发送成功~")
}

// go run main.go
fatal error: all goroutines are asleep - deadlock!
```

为什么会出现deadlock错误呢？

因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。

上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？

一种方法是启用一个goroutine去接收值，例如：

```GO
func receiver(ch chan int) {
	// 接收
	ret := <-ch
	fmt.Println("接收成功~", ret)
}

func main() {
	ch := make(chan int)
	// receiver(无缓冲通道 发送前先接收 防止死锁)
	go receiver(ch) // 启用goroutine从通道接收值
	// sned
	ch <- 100
	fmt.Println("发送成功！")
}

//  go run main.go
接收成功~ 100
发送成功！
```

无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到两一个goroutine在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。

**有缓冲通道**

为了解决以上问题还有另一种方法就是使用有缓冲的通道。

![使用有缓冲的通道在goroutine之间同步数据](https://tva1.sinaimg.cn/large/e6c9d24ely1gzmdovyeeyj20m70ewgn1.jpg)

我们可以在使用make函数初始化通道的时候为其指定通道的容量，例如：

```GO
func main() {
  ch := make(chan, int, 1) // 创建一个容量=1的有缓冲区的通道
  ch <- 100
  fmt.Println("发送成功!")
}
```

只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。就类似小区的快递柜是有固定格子数量的，格子蛮子就无法存放，导致阻塞，只能等到别人取走一个快递员就能往里面放一个。

我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。

**close()**

可以通过内置的close()函数关闭channel（如果你的管道不往里面存值或者取值的时候一定记得关闭管道）

```GO
func main() {
  ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

// go run main.go
0
1
2
3
4
main结束
```

**如何优雅的从通道循环取值**

当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭了呢？

```GO
func main() {
  // channel 练习
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0~100数据发送到ch1中
	go func() {
		for i := 0; i <= 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值 并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 如果通道关闭再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2接收值并打印
	for i := range ch2 { // 通道关闭后会退出 for range循环
		fmt.Println(i)
	}
}
```

从上面的例子中我们看到有两种方式在接收值的时候判断通道是否关闭，一般情况使用的是for range的方式。

**单向通道**

有的时候 我们会江通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或者只能接收。

Go语言中提供了单向通道来处理这种情况。例如，我们把上面的例子改造如下：

```go
func counter(out chan<- int) { // out是只能发送的channel
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) { // out只能发送 in只能接收
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
```

其中:

```GO
1.chan<- int 是一个只能发送的通道，可以发送但不能接收
2.<-chan int 是一个只能接收的通道，只能接收但不能发送
```

在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的。

**通道总结**

channel常见异常总结如下：

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1gznegf0a37j21bs0ii79d.jpg)

*注意：关闭已经关闭的channel也会引发panic*

#### 5、goroutine池

**worker pool（goroutine池）**

- 本质上是生产者消费者模型
- 可以有效控制goroutine数量，防止暴涨

需求：

- 计算一个数字各位数之和，例如数字123，结果为1+2+3=6
- 随机生成数字进行计算		

#### 6、定时器

- Timer：时间到了，执行只执行一次

- Ticker：时间到了，多次执行

```go
func main() {
    // 1.获取ticker对象
    ticker := time.NewTicker(1 * time.Second)
    i := 0
    // 子协程
    go func() {
        for {
            //<-ticker.C
            i++
            fmt.Println(<-ticker.C)
            if i == 5 {
                //停止
                ticker.Stop()
            }
        }
    }()
    for {
    }
}
```

#### 7、select

**select多路复用**

在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接受将会发生阻塞。你也许会写出如下代码使用遍历的方式来实现：

```go
for {
  // 尝试从ch1接收值
  data, ok := <-ch1
  // 尝试从ch2接收值
  data, ok := <-ch2
  ...
}
```

这种方式虽然可以实现从多个通道接收值的需求，但是运行性鞥会差很多。为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。

select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）。select会一直等待，知道某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下：

```GO
select {
  case <-chan1:
  // 如果chan1成功读取到数据，则进行当前case处理语句
  case chan2 <- 1:
  // 如果成功向chan2写入数据，则进行当前case处理语句
  default:
  // 如果上面都没有成功，则进行default处理流程
}
```

select可以同时监听一个或多个channel，直到其中一个channel ready。

```GO
package main

import (
	"fmt"
	"time"
)

// channel select

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1 func"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2 func"
}

func main() {
	// 两个管道
	out1 := make(chan string)
	out2 := make(chan string)
	// 跑2个子协程 写数据
	go test1(out1)
	go test2(out2)
	// 用select监控
	select {
	case str1 := <-out1:
		fmt.Println("str1:", str1)
	case str2 := <-out2:
		fmt.Println("str2:", str2)
	}
}
```

如果多个channel同时ready，则随机选择一个执行

```GO
func main() {
  int_chan := make(chan int, 1)
	str_chan := make(chan string, 1)
	go func() {
		int_chan <- 1
	}()
	go func() {
		str_chan <- "test"
	}()
	select {
	case value := <-int_chan:
		fmt.Println("int value=", value)
	case value := <-str_chan:
		fmt.Println("string value=", value)
	}
	fmt.Println("main结束~")
}
```

可以用于判断管道是否存满

```GO
func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello~")
		default:
			fmt.Println("channel full!")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
  // 判断通道是否存满
	ch := make(chan string, 10)
	// 子协程写数据
	go write(ch)
	// 取数据
	for s := range ch {
		fmt.Println("res=", s)
		time.Sleep(time.Second)
	}
}
```

#### 8、并发安全和锁

有时候在Go代码中可能会存在多个goroutine同事操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。类比现实生活中的例子有十字路口被各个方向的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

举个例子：

```GO
package main

import (
	"fmt"
	"sync"
)

// 竞态问题

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 1000; i++ {
		x+=1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

// go run main.go
1646
```

以上代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞态，导致最后的结果与预期不符。

**互斥锁**

互斥锁是一种常见的控制共享资源的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。使用互斥锁来修复上面代码的问题：

```GO
func add() {
	for i := 0; i < 1000; i++ {
		// 加锁
		lock.Lock()
		x += 1
		// 解锁
		lock.Unlock()
	}
	wg.Done()
}

// go run main.go
2000
```

使用互斥锁能搞保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同事等待一个锁时，唤醒的策略是随机的。

**读写互斥锁**

互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync.RWMutex类型。

读写锁分两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取的读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

读写锁示例：

```GO
package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁

var (
	n  int64
	wg sync.WaitGroup
	// lock   sync.Mutex
	rwLock sync.RWMutex
)

func write() {
	rwLock.Lock() // 加写锁
	n += 1
	time.Sleep(time.Millisecond * 2) // 假设写耗时2毫秒
	rwLock.Unlock()                  // 解写锁
	wg.Done()
}

func read() {
	rwLock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读耗时1毫秒
	rwLock.RUnlock()             // 解读锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

// go run main.go
2.273914864s
```

需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的又是就发挥不出来。

#### 9、sync

**sync.WaitGroup**

当代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。sync.WaitGroup有以下几个方法：

| 方法名         | 功能                |
| -------------- | ------------------- |
| Add(delta int) | 计数器 + delta      |
| Done()         | 计数器 -  1         |
| Wait()         | 阻塞直到计数器变为0 |

sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N个并发任务时，就将计数器值加N。每个任务完成时通过调用Done()方法将计数器-1.通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

我们利用sync.WaitGroup将上面的代码优化一下：

```go
var wg sync.WaitGroup

func hello() {
  defer wg.Done()
  fmt.Println("Hello Goroutine~~~")
}

func main() {
  wg.Add(1)
  go hello()
  fmt.Println("main goroutine done!")
  wg.Wait()
}
```

*注意：sync.WaitGroup是一个结构体，传递的时候要传递指针。*

**sync.Once**

前言：这是一个进阶知识点

在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

Go语言中的sync包中提供了一个针对只执行一次场景的解决方案：sync.Once

sync.Once只有一个Do方法，其签名如下：

```GO
func (o *Once) Do(f func()) {}
```

*注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。*

**加载配置文件示例**

延迟一个开销很大的初始化操作到真正用到它的时候在执行是一个很好的实践。因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。我们来看一个例子：

```GO
var icons map[string]image.Image

func loadUcons() {
  icons = map[string]image.Image{
    "left": loadIcon("left.png"),
    "up": loadIcon("up.png"),
    "right": loadIcon("right.png"),
    "down": loadIcon("down.png"),
  }
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image {
  if icons == nil {
    loadIcons()
  }
  return icons[name]
}
```

多个goroutine并发调用Icon函数时不是并发安全的，现代的编译器和CPU可能会在保证每个goroutine都满足串行一致的基础上自由地重新排访问内存的顺序。loadIcons函数可能会被重排为以下结果：

```GO
func loadIcons() {
    icons = make(map[string]image.Image)
    icons["left"] = loadIcon("left.png")
    icons["up"] = loadIcon("up.png")
    icons["right"] = loadIcon("right.png")
    icons["down"] = loadIcon("down.png")
} 
```

这种情况下就会出现即使判断了icons不是nil也不意味着变量初始化完成了。考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化icons的时候不会被其他的goroutine操作，但是这样做又会引发性能问题。

使用sync.Once()改造的示例代码如下：

```go
var icons map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons() {
    icons = map[string]image.Image{
        "left":  loadIcon("left.png"),
        "up":    loadIcon("up.png"),
        "right": loadIcon("right.png"),
        "down":  loadIcon("down.png"),
    }
}

// Icon 是并发安全的
func Icon(name string) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
```

sync.Once()其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值的数据和安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化的时候是并发安全的并且初始化操作也不会被执行多次。

**sync.Map**

Go语言中内置的map并不是并发安全的。请看如下示例：

```GO
package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			// strconv.Itoa函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("key:%s, value:%d\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// go run main.go
fatal error: concurrent map writes
key:14, value:14

goroutine 18 [running]:
runtime.throw(0x10d06c2, 0x15)
        /usr/local/go/src/runtime/panic.go:1116 +0x72 fp=0xc0001066c8 sp=0xc000106698 pc=0x1031e32
runtime.mapassign_faststr(0x10b7e20, 0xc000064180, 0x10d56eb, 0x2, 0x0)
        /usr/local/go/src/runtime/map_faststr.go:211 +0x3f1 fp=0xc000106730 sp=0xc0001066c8 pc=0x1011ed1
main.set(...)
……
```

以上代码开启少量几个goroutine的时候可以正常运行，但是并发多了之后执行就会报以上错误。

在这种场景下就需要为map加锁来保证并发安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版map：sync.map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同事sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

```go
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			// strconv.Itoa函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字
			key := strconv.Itoa(n)
			// set(key, n)
			m.Store(key, n)          // 写入
			value, _ := m.Load(key)  // 读取
			fmt.Printf("key:%s, value:%d\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

#### 10、原子操作（atomic包）

**原子操作**

代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法他在用户态就可以完成，因此性能比加锁操作更好。Go语言中原子操作由内置的标准库sync/atomic提供。

**atomic包**

| 方法                                                         | 解释           |
| :----------------------------------------------------------- | :------------- |
| func LoadInt32(addr *int32) (val int32)func LoadInt64(addr `*int64`) (val int64)<br>func LoadUint32(addr`*uint32) (val uint32)<br>func LoadUint64(addr*uint64`) (val uint64)<br>func LoadUintptr(addr`*uintptr) (val uintptr)<br>func LoadPointer(addr*unsafe.Pointer`) (val unsafe.Pointer) | 读取操作       |
| func StoreInt32(addr `*int32`, val int32) func StoreInt64(addr `*int64`, val int64) func StoreUint32(addr `*uint32`, val uint32) func StoreUint64(addr `*uint64`, val uint64) func StoreUintptr(addr `*uintptr`, val uintptr) func StorePointer(addr `*unsafe.Pointer`, val unsafe.Pointer) | 写入操作       |
| func AddInt32(addr `*int32`, delta int32) (new int32) func AddInt64(addr `*int64`, delta int64) (new int64) func AddUint32(addr `*uint32`, delta uint32) (new uint32) func AddUint64(addr `*uint64`, delta uint64) (new uint64) func AddUintptr(addr `*uintptr`, delta uintptr) (new uintptr) | 修改操作       |
| func SwapInt32(addr `*int32`, new int32) (old int32) func SwapInt64(addr `*int64`, new int64) (old int64) func SwapUint32(addr `*uint32`, new uint32) (old uint32) func SwapUint64(addr `*uint64`, new uint64) (old uint64) func SwapUintptr(addr `*uintptr`, new uintptr) (old uintptr) func SwapPointer(addr `*unsafe.Pointer`, new unsafe.Pointer) (old unsafe.Pointer) | 交换操作       |
| func CompareAndSwapInt32(addr `*int32`, old, new int32) (swapped bool) func CompareAndSwapInt64(addr `*int64`, old, new int64) (swapped bool) func CompareAndSwapUint32(addr `*uint32`, old, new uint32) (swapped bool) func CompareAndSwapUint64(addr `*uint64`, old, new uint64) (swapped bool) func CompareAndSwapUintptr(addr `*uintptr`, old, new uintptr) (swapped bool) func CompareAndSwapPointer(addr `*unsafe.Pointer`, old, new unsafe.Pointer) (swapped bool) | 比较并交换操作 |

**示例**

我们填写一个示例来比较下互斥锁和原子操作的性能。

```GO
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 比较互斥锁和原子操作性能

var (
	x    int64
	lock sync.Mutex
	wg   sync.WaitGroup
)

// 普通加锁版本
func add() {
	x++
	wg.Done()
}

// 加强版互斥锁
func mutexAdd() {
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}

// 终极版原子操作
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//  普通 不是并发安全
		go add()
		// 加强 是并发安全 但性能开销大
		// go mutexAdd()
		// 终极 并发安全 性能优于加锁版
		// go atomicAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}
```

|      类型      | 时间（毫秒） |
| :------------: | :----------: |
|   普通函数版   |  3.592815ms  |
|   加强加锁版   |  3.22896ms   |
| 终极原子操作版 |  3.090508ms  |

atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。

#### 11、GMP原理与调度

先略过，后期补充

####12、爬虫案例

**思路**

1. 明确目标
2. 爬取数据
3. 筛选数据
4. 处理数据

### 数据操作

#### 1、Go操作MySQL

新建test数据库，person、place表

```mysql
CREATE TABLE `person` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(260) DEFAULT NULL,
    `sex` varchar(260) DEFAULT NULL,
    `email` varchar(260) DEFAULT NULL,
    PRIMARY KEY (`user_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE place (
    country varchar(200),
    city varchar(200),
    telcode int
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
```

![image-20220301212358265](https://tva1.sinaimg.cn/large/e6c9d24ely1gzup87td3vj20el07wq3o.jpg)

**mysql使用**

使用第三方开元的mysql库：github.com/go-sql-driver/mysql （mysql驱动）
github.com/jmoiron/sqlx （基于mysql驱动的封装）

命令行输入：

```GO
    go get github.com/go-sql-driver/mysql 
    go get github.com/jmoiron/sqlx     
```

连接mysql：

```GO
database, err: = sqlx.Open("mysql","root:XXXX@tcp(127.0.0.1:3306)/test")
//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")  
```

**insert操作**

```go
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:zhangyun..@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu002", "man", "stu02@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	Db.Close()
	fmt.Println("insert succ:", id)
}

// go run main.go
insert succ: 2
```

**select操作**

```GO
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:zhangyun..@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	var person []Person
	defer Db.Close()
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select success:", person)
}

// go run main.go
select success: [{2 stu001 man stu01@qq.com}]
```

**update**

```GO
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:zhangyun..@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	res, err := Db.Exec("update person set username = ? where user_id = ?", "stu_03", 3)
	defer Db.Close()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row failed, ", err)
		return
	}
	fmt.Println("update success:", row)
}

// go run main.go
update success: 1
```

**delete**

```GO
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:zhangyun..@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	res, err := Db.Exec("delete from person where user_id = ?", 2)
	defer Db.Close()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	Db.Close()
	fmt.Println("delete success:", row)
}

// go run main.go
delete success: 1
```

**MySQL事务**

mysql事务特性：

1. 原子性
2. 一致性
3. 隔离性
4. 持久性

Golang MySQL事务应用：

```GO
1.import (""github.com/jmoiron/sqlx)
2.Db.Begin()    开始事务
3.Db.Commit()   提交事务
3.Db.Rollback() 回滚事务
```

```GO
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// MySQL 事务
type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:zhangyun..@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func insert(username, sex, email string) (sql string) {
	sql = fmt.Sprintf("insert into person(username, sex, email)values(%s, %s, %s)", username, sex, email)
	fmt.Println("sql:", sql)
	return sql
}

func main() {
	conn, err := Db.Begin()
	defer Db.Close()
	if err != nil {
		fmt.Println("begin failed, err", err)
		return
	}
	r, err := conn.Exec(insert("'stu_004'", "'man'", "'stu_004@qq.com'"))
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	fmt.Println("insert success", id)

	r, err = conn.Exec(insert("'stu_005'", "'man'", "'tu_005@qq.com'"))
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	fmt.Println("insert success", id)

	// 提交事务
	err = conn.Commit()
	if err != nil {
		fmt.Println("commit failed, err", err)
		return
	}
}

// go run main.go
sql: insert into person(username, sex, email)values('stu_004', 'man', 'stu_004@qq.com')
insert success 4
sql: insert into person(username, sex, email)values('stu_005', 'man', 'tu_005@qq.com')
insert success 5
```

![image-20220302165612786](https://tva1.sinaimg.cn/large/e6c9d24ely1gzvn3u44v9j20ae03m0sx.jpg)

#### 2、Go操作Redis

**Redis介绍**

Redis是完全开源免费的，遵守BSD协议，是一个高性能的key-value数据库。

Redis与其他kv缓存茶农相比有以下三个特点：

1. Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
2. Redis不仅仅支持简单的kv类型的数据，同事还提供string、list（链表）、set（集合）、hash表等数据结构的存储。
3. Redis支持数据的备份，即master-slave模式的数据备份。

**Redis优势**

- 性能极高
  - Redis能读的速度是110000次/s，写的速度是81000/s，单机能够达到15w QPS，通常适合做缓存。
- 丰富的数据类型
  - Redis支持二进制案例的strings、list、hashes、sets及ordered sets数据类型操作。
- 原子
  - Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULII和EXEC指令包起来。
- 丰富的特性
  - Redis还支持publish/subscribe，通知，key过期等等特性。

Redis与其他的k-v存储有何不同？

1. Redis有着更为复杂的数据结构并且提供对他们的原子性操作，这是一个不同于其他数据库的进化路径。Redis的数据类型都是基于基本数据结构的同时对程序员透明，无需进行额外的抽象。
2. Redis运行在内存中但是可以持久化到磁盘，所以在对不同数据集进行高速读写时需要权衡内存，因为数据量不能大于硬件内存。在内存数据库方面的另一个优点是，相比在磁盘上相同的复杂的数据结构，在内存中操作起来非常简单，这样Redis可以做很多内部复杂性很强的事情。同时，在磁盘格式方面他们是紧凑的以追加的方式产生的，因为他们并不需要进行随机访问。

**Redis使用**

使用第三方开源的Redis库：github.com/garyburd/redigo/redis

命令行：

```go
go get github.com/garyburd/redigo/redis
```

**连接Redis**

```GO
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Golang 连接 Redis
func main()  {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn conn_pool failed, err", err)
		return
	}
	fmt.Println("conn_pool conn success!")

	defer client.Close()
}

// go run main.go
redis conn success!
```

**String类型Set、Get操作**

```go
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Golang 连接 Redis
func main()  {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn conn_pool failed, err", err)
		return
	}
	fmt.Println("conn_pool conn success!")

	defer client.Close()

	_, err = client.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println("client do failed, err:", err)
		return
	}

	r, err := redis.Int(client.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed, err:", err)
		return
	}
	fmt.Println(r)
}

//  go run main.go
redis conn succe
100
```

**string批量操作**

```GO
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Golang 连接 Redis
func main() {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn conn_pool failed, err", err)
		return
	}
	fmt.Println("conn_pool conn success!")

	defer client.Close()

	_, err = client.Do("MSet", "abc", 100, "efg", 200)
	if err != nil {
		fmt.Println("client do failed, err:", err)
		return
	}

	r, err := redis.Ints(client.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	// 遍历r 批量取
	for k, v := range r {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}
}

//  go run main.go
redis conn success!
k:0, v:100
k:1, v:200
```

**设置过期时间**

```GO
import (
    "fmt"
    "github.com/garyburd/redigo/conn_pool"
)

func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn conn_pool failed,", err)
        return
    }

    defer c.Close()
    _, err = c.Do("expire", "abc", 10)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

**List队列操作**

```GO
func main() {
  client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn conn_pool failed, err", err)
		return
	}
	fmt.Println("conn_pool conn success!")

	defer client.Close()
  
  // List 队列操作
	_, err = client.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(client.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}
```

**Hash表**

```GO
func main() {
  client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn conn_pool failed, err", err)
		return
	}
	fmt.Println("conn_pool conn success!")

	defer client.Close()
  
  // hash表
	_, err = client.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(client.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}
```

**连接池**

```GO
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Redis 连接池

var pool *redis.Pool  // 创建Redis连接池

func init() {
	pool = &redis.Pool{
		// 实例化一个连接池
		MaxIdle: 16,  // 初始连接数量
		MaxActive: 0, // redis的最大连接数量（0：不确定）
		IdleTimeout: 300,  // 连接关闭时间300秒（300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			// 要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	client := pool.Get()  // 从连接池，取一个连接
	defer client.Close()  // 函数运行结束，将连接放回连接池

	_, err := client.Do("Set", "wanli", 200)
	if err != nil {
		fmt.Println("redis do failed, err:", err)
		return
	}

	r, err := redis.Int(client.Do("Get", "wanli"))
	if err != nil {
		fmt.Println("get wanli failed, err:", err)
		return
	}
	fmt.Println(r)
	_ = pool.Close() // 关闭连接池
}

//  go run main.go
200
```

**3、Go操作ETCD**

**ETCD介绍**

ETCD是使用Go语言开发的一个开源的、高可用的分布式k-v存储系统，可以用于配置共享和服务的注册和发现。

类似项目有zookeeper和consul。

ETCD具有以下特点：

1. 完全复制：集群中的每个节点都可以使用完整的存档
2. 高可用性：ETCD可用于避免硬件的单点故障或网络问题
3. 一致性：每次读取都会返回跨多主机的最新写入
4. 简单：包括一个定义良好、面向用户的API（gRPC）
5. 安全：实现了带有可选的客户端正数身份验证的自动化TLS
6. 快速：每秒10000次写入的基准速度
7. 可靠：使用Raft算法实现了强一致、高可用的服务存储目录

**ETCD应用场景**

**服务发现**

服务发现要解决的也是分布式系统中最常见的问题之一，即在同一个分布式集群中的进程或服务，要如何才能找到对方并建立连接。本质上来说，服务发现就是想要了解集群中是否有进程在监听udp或tcp端口，并且通过名字就可以查找和连接。

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1gzwqkp3virj20w80isgmh.jpg)

**配置中心**

将一些配置信息放到ETCD上进行集中管理。

这类场景的使用方式通常是这样：应用在启动的时候主动从etcd获取一次配置信息，同时，在etcd节点上注册一个watcher并等待，以后每次配置有更新的时候，etcd都会实时通知订阅者，以此达到获取最新配置信息的目的。

**分布式锁**

因为etcd使用Raft算法保持了数据的强一致性，某此操作存储到集群中的值必然是全局一致的，所以很容易实现分布式锁。锁服务有两种使用方式，一是保持独占，而是控制时序。

- 保持独占

  保持独占即所有获取锁的用户最终只有一个可以得到。etcd为此提供了一套实现分布式锁原子操作CAS（CompareAndSwap）的API。通过设置prevExist值，可以保证在多个节点同时去创建某个目录时，只有一个成功。而创建成功的用户就可以认为是获得了锁

- 控制时序

  即所有想要获得锁的用户都会被安排执行，但是获得锁的顺序也是全局唯一的，同时决定了执行顺序。etcd为此也提供了一套API（自动创建有序建），对一个目录建值时指定为POST动作，这样etcd会自动在目录下生成一个当前最大的值作为键，存储这个新的值（客户端编号）。同时还可以使用API按顺序列出所有当前目录下的键值。此时这些键的值就是客户端的时序，而这些键中存储的值是代表客户端的编号

  ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1gzwtzp4lomj216w0iyjt2.jpg)

**为什么用ETCD而不是ZooKeeper**

为什么不选择zookeeper？

- 部署维护复杂，其使用的Paxos强一致性算法复杂难懂。官方只提供了Java和C两种语言的接口
- 使用Java编写引入大量的依赖，维护麻烦
- 最近几年发展缓慢，不如etcd和consul等后起之秀

etcd优点

- 简单。使用Go编写部署简单；支持HTTP/JSON API，使用简单；使用Raft算法保证强一致性让用户易于理解
- etcd默认数据一更新就进行持久化
- etcd支持SSL客户端安全认证

最后，etcd作为一个年轻的项目，正在高速迭代和开发中，这既是一个优点，也是一个缺点。优点是它的未来具有无限的可能性，缺点是无法得到大项目长时间使用的检验。然而，目前 CoreOS、Kubernetes和CloudFoundry等知名项目均在生产环境中使用了etcd，所以总的来说，etcd值得你去尝试。

**ETCD集群**

etcd作为一个高可用键值存储系统，天生就是为集群化而设计的。由于Raft算法在做决策时需要多数节点的投票，所以etcd一般部署集群推荐奇数个节点，推荐的数量为3、5或者7个节点构成一个集群。

**操作ETCD**

这里使用官方的`etcd/clientv3`包来连接etcd并进行相关操作。

**安装**

```go
go get go.etcd.io/etcd/clientv3
```

**Put和Get操作**

put命令用来设置键值对数据，get命令用来根据key获取值。

……

#### 4、ZooKeeper

**基本操作**

安装

```GO
go get github.com/samuel/go-zookeeper/zk
```

**简单的分布式server**

目前分布式系统已经很流行了，一些开源框架也被广泛应用，如dubbo、Motan等。对于一个分布式服务，最基本的一项功能就是服务的注册和发现，而利用zk的EPHEMERAL节点则可以很方便的实现该功能。EPHEMERAL节点正如其名，是临时性的，其生命周期是和客户端会话绑定的，当会话连接断开时，节点也会被删除。下边我们就来实现一个简单的分布式server。

**server**

服务启动时，创建zk连接，并在go_servers节点下创建一个新节点，节点名为”ip:port”，完成服务注册
服务结束时，由于连接断开，创建的节点会被删除，这样client就不会连到该节点

**client**

先从zk获取go_servers节点下所有子节点，这样就拿到了所有注册的server
从server列表中选中一个节点（这里只是随机选取，实际服务一般会提供多种策略），创建连接进行通信
这里为了演示，我们每次client连接server，获取server发送的时间后就断开。

#### 5、Go操作Kafka

**kafka介绍**

- kafka是什么

  - kafka使用Scala开发，支持多语言客户端（C艹、Java、Python、Golang等）
  - kafka最先由LinkedIn公司开发，之后成为Apache的顶级项目
  - kafka是一个分布式的、分区化、可复制提交的日志服务
  - LinkedIn使用kafka实现了公司不同应用程序之间的松耦合，作为一个可扩展、高可靠的消息系统
  - 支持高throughput的应用
  - scale out：无需停机即可扩展机器
  - 持久化：通过将数据持久化到硬盘以及replication防止数据丢失
  - 支持online和offline的场景

- kafka特点

  - kafka是分布式的，其所有的构件borker（服务端集群）、producer（消息生产）、consumer（消息消费者）都可以是分布式的
  - 在消息的生产时可以使用一个表示topic来区分，且可以进行分区；每一个分区都是一个顺序的、不可变的消息队列，并且可以持续的增加
  - 同时为发布和订阅提供高吞吐量，据了解mkafka每秒可以生产约25万消息（50mb），每秒处理55万消息（110mb）
  - 消息被处理的状态是在consumer端维护，而不是由server端维护。当失败时能自动平衡

- 常用的场景

  - 监控：主机通过kafka发送与系统和应用程序健康相关的指标，然后这些信息会被收集和处理从而创建监控仪表盘并发送警告
  - 消息队列：应用程序使用kafka作为传统的消息系统实现标准的队列和消息的发布->订阅。例如搜索和内容提要（content feed）。比起大多数的消息系统来说kafka有更好的吞吐量，内置的分区，冗余及容错性，这让kafka成为了一个很好的大规模消息处理应用的解决方案。消息系统一般吞吐量相对较低，但是需要更小的端到端延时，并常常依赖于kafka提供的强大的持久性保障。在这个领域，kafka足以媲美传统消息领域，如ActiveMR或RabbitMQ
  - 站点的用户活动追踪：为了更好地理解用户行为，改善用户体验，将用户查看了哪个页面、点击了哪些内容等信息发送到每个数据中心的kafka集群上，并通过Hadoop进行分析，生成日常报告。
  - 流处理：保存收集流数据，以提供之后对接的storm或其他流式计算框架进行处理。很多用户会将从原始topic来的数据进行阶段性处理，汇总，扩充或者其他的方式转换到新的topic下再继续后面的处理。例如一个文章推荐的处理流程，可能实现从RSS数据源中抓取文章的内容，然后将其丢入一个叫做“文章”的topic中了后续操作可能是需要对这个内容进行清理，比如恢复正常数据或者删除重复数据，最后再将内容匹配的结果返回给用户。这就在一个独立的topic之外，产生了一系列的实时数据处理的流程。
  - 日志聚合：使用kafka代替日志聚合（log aggregation）。日志聚合一般来说就是从服务器上收集日志文件，然后放到一个集中地位置（文件服务器或HDFS）进行处理。然后kafka忽略掉文件的细节，将更清晰地抽象成一个个日志或事件的消息流。这就让kafka处理过程延迟更低，更容易支持多数据源和分布式数据处理。比起以日志为中心的系统比如scribe或者flume来说，kafka提供同样高效的性能和因为复制导致的更的耐用性保证，以及更低的端到端延迟。
  - 持久性日志：kafka可以为一种外部的持久性日志的分布式系统提供服务。这种日志可以在节点接备份数据，并为故障节点数据恢复提供一种重新同步的机制。kafka中日志压缩功能为这种用法提供了条件。在这种用法中，kafka类似于Apache bookkeeper项目

  **kafka包含以下基础概念**

  1. Topic（话题）：kafka中用于区分不同类别信息的类别名称。由producer指定
  2. Producer（生产者）：将消息发布到kafka特定的topic的对象（过程）
  3. Consumer（消费者）：订阅并处理特定的topic中的消息的对象（过程）
  4. Broker（kafka服务集群）：已发布的消息保存在一组服务器中，称之为kafka集群。集群中的每一个服务器都是一个代理（Broker）。消费者可以订阅一个或多个话题，并从Broker拉数据，从而消费这些已发布的消息
  5. Partition（分区）：topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）
  6. Message（消息）：消息是通信的基本单位，每个producer可以向一个topic（主题）发布一些消息。

**消息**

消息由一个固定大小的报头和可变长度但不透明的字节阵列负载。报头包含格式版本和CRC32效验和以检测损坏或截断

**kafka深层介绍**

**架构介绍**

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h06170itdhj21hm0srdlj.jpg)

- producer：生产者，消息的产生者，是消息的入口。
- kafka cluster：kafka集群，一台或多台服务器组成
  - broker：broker 试制部署了 kafka 实例的服务器节点。每个服务器上有一个或多个 kafka 的实例，我们暂且认为每个 broker 对应一台服务器。每个 kafka 集群内的 broker 都有一个不重复的编号，如图中的 broker-0、broker-1等
  - topic：消息的主题，可以理解为消息的分类，kafka 的数据就保存在 topic 中。在每个 broker 上可以创建多个 topic。实际应用中通常是一个业务线建一个 topic
  - partition：topic 的分区，每个 topic 可以有多个分区，分区的作用是做负载，提高 kafka 的屯度量。同一个 topic在不同的分区的数据是不重复的，partition 的表现形式就是一个又一个的文件夹
  - replication：每一个分区都有多个副本，副本的作用就是做备胎。当主分区（leader）故障的时候会选择一个备胎（follower）上位，成为 leader。在 kafka 中默认副本的最大数量是 10 个，且副本的数量不能大于 broker 的数量，follower 和 leader 绝对是在不同的机器，同一机器对同一个分区也只可能存放一个副本（包括自己）
  - consumer：消费者，即消息的消费方，是消息的出口
    - consumer group：我们可以将多个消费组组成一个消费者组，在 kafka 的设计中同一个分区的数据只能被消费者组中的某一个消费者消费。同一个消费者组的消费者可以消费同一个 topic 的不同分区的数据，这也是为了提高 kafka 的吞吐量

**工作流程**

我们看上面的架构图中，producer 就是生产者，是数据的入口。producer 在写入数据的时候把数据写入到leader 中，不会直接将数据写入到 follower，那如何寻找 leader 呢？写入的流程又是怎样的？我们来看下面这张图：

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0733yf0i6j211e0dqaao.jpg)

1. 生产者从 kafka 集群获取分区 leader 信息
2. 生产者将消息发送给 leader
3. leader 将消息写入本地磁盘
4. follower 从 leader 拉取消息数据
5. follower 将消息写入本地磁盘后向 leader 发送 ack
6. leader 收到所有的 follower 的 ack 之后向生产者发送 ack

**选择 partition 的原则**

那么在kafka中，如果某个 topic 有多个 partition，producer 又怎么知道该将数据发送给那个 partition呢？kafka 中有几个原则：

1. partition 在写入的时候可以指定需要写入的 partition，如果有指定，则写入相应的 partition
2. 如果没有指定 partition，但是设置了数据的 key，则会根据 key 的值hash 出一个 partition
3. 如果既没有指定 partition，又没有设置 key，则会采用轮询方式，即每次取一小段时间的数据写入某个 partition，下一小段的时间写入下一个 partition

**ACK 应答机制**

producer 再向 kafka 写入消息的时候，可以设置参数来确定是否确认 kafka 接收到数据，这个参数可以设置的值为 0，1，all

- 0：表示 producer 往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低，但是效率最高
- 1：表示 producer 往集群发送数据只要leader 应答就可以发送下一条，之确保 leader 发送成功
- all：表示 producer 往集群发送数据需要所有的 follower 都完成从 leader 的同步才会发送下一条，确保 leader 发送成功和所有的副本都完成备份。安全性最高，但是效率最低

最后需要注意的是，如果往不存在的 topic 写数据，kafka 会自动创建 topic，partition 和 replication的数量默认配置都是 1.

**Topic 和数据日志**

topic 是同一类别的消息记录（record）的集合。在 kafka 中，一个主题通常有多个订阅者。对于每个主题，kafka 集群维护了一个分区数据日志文件结构如下：

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h09fm1isk0j20bg07daac.jpg)

每个 partition 都是一个有序并且不可变的消息记录集合。当心的数据写入时，就被追加到 partition 的末尾。在每个 partition 中，每条消息都会被分配一个顺序的唯一标识，这个表示被称为 offset，即偏移量。注意，kafka 只保证在同一个partition 内部消息是有序的，在不同的 partition 之间，并不能保证消息有序。

kafka 可以配置一个保留期限，用来标识日志会在 kafka 集群内保留多长时间。kafka 集群会保留在保留期限内所有被发布的消息，不管这些消息是否被消费过。比如保留期限设置为 2 天，那么数据被发布到 kafka 集群的两天内，所有的这些数据都可以被消费。但是超过两天，这些数据将会被清空，以便为后续的数据腾出空间。由于 kafka 会将数据进行持久化存储（即写入到硬盘上），所有保留的数据大小可以设置一个比较大的值。

**Partition结构**

partition 在服务器上的表现形式就是一个一个的文件夹，每个 partition 的文件夹下面会有多组 segment文件，每组 segment 文件又包含.index 文件、.log文件、.timeindex 文件三个文件，其中.log 文件就是实际存储 message 的地方，而.index 和.timeindex 文件为索引文件，用于检索消息。

**消费数据**

多个消费者实例可以组成一个消费者组，并用一个标签来标识这个消费者组。一个消费者组中的不同消费者实例可以运行在不同的进程甚至不同的服务器上。

如果所有的消费者实例都在同一个消费者组中，那么消息记录会被很好地均衡的发送到每个消费者实例。

如果所有的消费者实例都在不同的消费者组，那么每一条消息记录会被广播到每一个消费者实例。

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h09i5vox4tj20eu080wew.jpg)

举个例子：上图所示一个两个节点的 kafka 集群上拥有一个四个 partition（P0~P3）的 topic有两个消费者组都在消费这个 topic 的数据，，消费者组 A 有2个消费者实例，消费者组B 有 4 个消费者实例.

从图中我们可以看到，在同一个消费者组中，每个消费者实例可以消费多个分区，但是每个分区最多只能被消费者中的一个实力消费。也就是说，如果有一个 4 个分组的主题，那么消费者组中最多只能有 4 个消费者实例去消费，多出来的都不会被分配到分区。其实这也很多理解，如果允许两个消费者实例同时消费同一个分区，那么就无法记录这个分区被这个消费者组消费的 offset 了。如果在消费者组中动态的上线或下线消费者，那么 kafka 集群会自动调整分区与消费者实例间的对应关系。

**kafka 使用**

- 安装
  - brew install zookeeper
  - brew install kafka
- 启动
  - 启动 zookeeper命令：`/usr/local/Cellar/kafka/3.1.0/bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties &`
  - ![image-20220316144451066](https://tva1.sinaimg.cn/large/e6c9d24ely1h0bpzj4mm6j20tq0esn2t.jpg)
  - 启动 kafka 命令：`/usr/local/Cellar/kafka/3.1.0/bin/kafka-server-start /usr/local/etc/kafka/server.properties &`
  - ![image-20220316144556431](https://tva1.sinaimg.cn/large/e6c9d24ely1h0bq0mharuj20qd0evmzc.jpg)

- 查看是否启动成功
  - 使用命令：`ps aux | grep kafka`和`ps aux | grep zookeeper`
  - 可以正常看到进行就说明已经启动成功了
  - ![image-20220316144805399](https://tva1.sinaimg.cn/large/e6c9d24ely1h0bq2x7ug1j20uu04t0ub.jpg)

- 创建 topic（test）
  - 创建一个名为 test 的 topic：`kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test`
  - ![image-20220316145358686](https://tva1.sinaimg.cn/large/e6c9d24ely1h0bq8zeykxj20mr0140sp.jpg)

- 查看 topic 列表
  - `kafka-topics --list --bootstrap-server  localhost:9092`
  - ![image-20220316145505307](https://tva1.sinaimg.cn/large/e6c9d24ely1h0bqa55mdsj20d4014wed.jpg)

- 向 topic 发送消息

  - 向 topic（test）中发送消息：`kafka-console-producer --broker-list localhost:9092 --topic test`

  ![image-20220316145633387](https://tva1.sinaimg.cn/large/e6c9d24ely1h0bqbnwpevj20fl01wdfu.jpg)

- 接收消息

  - 从 topic（test）中接收消息：` kafka-console-consumer --bootstrap-server localhost:9092 --topic test --from-beginning`

  ![image-20220316145741038](https://tva1.sinaimg.cn/large/e6c9d24ely1h0bqctzd2wj20ir029q2y.jpg)

**Go 操作kafka**

**第三方库**

Go语言中连接kafka使用第三方库: github.com/Shopify/sarama。

**Go操作RabbitMQ**

****

**RabbitMQ 介绍**

- 简单释义
  - 消息总线（Message Queue），是一种跨进程、异步的通信机制，用于上下游传递消息。由消息系统来确保消息的可靠传递。
- 背景描述
  - 当前市面上的 mq 产品很多，比如 RabbitMQ、Kafka、ActiveMQ、ZeaoMQ 和阿里巴巴捐献给 Apache 的 RocketMQ都支持 MQ 的功能。
- 适用场景
  - 上下游逻辑解耦&&物理解耦
  - 保证数据最终一致性
  - 广播
  - 错峰流控等

**RabbitMQ 特点**

RabbitMQ是由 erlang 语言开发的 AWQP 的开源实现。

AWQP:Advanced Message Queue，高级消息队列协议。它是应用层协议的一个开放标准，为面向消息的中间件设计，基于此协议的客户端与消息中间件可传递消息，并不受产品、开发语言等的限制。

- 可靠性 （reliablity）：使用了一些机制来保证可靠性，比如持久化、传输确认、发布确认
- 灵活的路由（flexible routing）：在消息进入队列之前，通过 exchange 来路有消息。对于典型的路由功能，rabbit 已经提供了一些内置的 exchange 来实现。针对更复杂的路由功能，可以将多个 exchange 绑定在一起，也通过插件机制实现自己的 exchange
- 消息集群 （clustering）：多个 RabbitMQ服务器可以组成一个集群，形成一个逻辑 broker
- 高可用（highty avaliable queues）：队列可以再急群众的机器上进行镜像，使得在部分节点出问题的情况下队列仍然可用。
- 多种协议（multi-protocol）：支持多重消息队列协议，如 STOMP、MQTT 等
- 多种语言客户端（many clients）：几乎支持所有常用语言，如 Java、Ruby 等
- 管理界面（management UI）：提供了易用的用户界面，使得用户可以监控和管理消息 broker 的许多方面
- 跟踪机制（Tracing）：如果消息异常，RabbitMQ 提供了消息的跟踪机制，使用者可以找出发生了什么
- 插件机制（plugin system）：提供了许多插件，来从多方面进行拓展，也可以编辑自己的插件

**RabbitMQ 简单使用**

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0dvcnfgv5j20b5034mx5.jpg)

所有 MQ 产品从模型抽象来说，都是一样的过程：

- 消费者（consumer）订阅某个队列
- 生产者（product）创建消息，然后发布到队列中（queue），最终将消息发送到监听的消费者

> 这只是最简单抽象的描述，具体到RabbitMQ则由更详细的概念需要解释。

![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0dvg11ixlj20fe04eaab.jpg)

- Broker：表示消息队列服务器实体
- Virtual Host：虚拟主机。标识一批交换机、消息队列和相关对象。虚拟主机是共享相同的身份认证和加密环境的独立服务器域。每个vhost本质上就是一个mini版的RabbitMQ服务器，拥有自己的队列、交换器、绑定和权限机制。vhost是AMQP概念的基础，必须在链接时指定，RabbitMQ默认的vhost是 /
- Exchange：交换机，用来接收生产者发送的消息并将这些消息路由给服务器中的队列
- Queue：消息队列，用来保存消息知道发送给消费者。他是消息的容器，也是消息的重点。一个消息可投入一个或多个队列。消息一直在队列里面，等待消费者连接到这个队列将其取走
- Banding：绑定，用于消息队列和交换机之间的关联。一个绑定就是基于路由键将交换机和消息队列连接起来的路由规则，所以可以将交换器理解成一个由绑定构成的路由表
- Channel：通道，多路复用连接中的一条独立的双向数据流通道。新到是建立在真实的TCP连接内地虚拟链接，AMQP命令都是通过新到发出去的，不管是发布消息、订阅队列还是接收消息，这些动作都是通过信道完成。因为对于操作系统来说，建立和销毁TCP都是非常昂贵的开销，所以引入了信道的概念，以复用一条TCP连接
- Connection：网络连接，比如一个 TCP 连接
- Publisher：消息的生产者，也是一个向交换器发布消息的客户端应用程序
- Consumer：消息的消费者，表示一个从一个消息队列中取得消息的客户端应用程序
- Message：消息，消息是不具名的，它是由消息头和消息体组成。消息体是不透明的，而消息头则是由一系列的可选属性组成，这些属性包括routing-key(路由键)、priority(优先级)、delivery-mode(消息可能需要持久性存储[消息的路由模式])等

**RabbitMQ 的 6 种工作方式**

1. simple 简单模式

   ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0e51e41opj20b5034mx5.jpg)

   - 消息产生者将消息放入队列
   - 消息的消费者监听消协队列，如果队列中有消息，就消费掉。消息被那走后，自动从队列中删除（隐患：消息可能没有被消费者正确处理，已经从队列中消失了，造成消息的丢失）
   - 应用场景：聊天（中间有一个过度的服务器；P 端，C 端）

2. work 工作模式（资源的竞争）

   ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0e5d6rww2j209t03p0sr.jpg)

   - 消息产生着将消息放入队列消费者可以有多个，消费之 1、消费者 2，同时监听同一个队列。C1、C2共同争抢当前消息队列内容，谁先拿到谁来消费消息（隐患：高并发情况下，默认会产生某一个消息被多个消费者共同使用，可以设置可一个开关 syncronize，与同步锁的性能不一样，保证一条消息只能被一个消费者使用）
   - 应用场景：红包；大项目中的资源调度（任务分配系统无需知道哪一个任务执行系统在空闲，直到将任务扔到消息队列中，空闲的系统自动争抢）

3. publish/subscribe发布订阅（共享资源）

   ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0f17byzsbj20b604kwel.jpg)

   - x 代表交换机 rabbitMQ 内部组件，erlang 消息产生者是代码完成，代码的执行效率不高，消息产生者将消息放入交换机，交换机发布订阅把消息发送到所有消息队列中，对应消息队列的消费者拿到消息进行消费
   - 相关场景：邮件群发、群聊天、广播（广告）

4. routing 路由模式

   ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0f19jm0upj20bv04d74f.jpg)

   - 消息生产者将消息发送给交换机按照路由判断，路由是字符串（info）当前产生的消息携带路由字符（对象的方法），交换机根据路由的 key，只能匹配上路由 key 对应的消息队列，对应的消费者才能消费消息
   - 根据业务功能定义路由字符串
   - 从系统的代码逻辑中获取对应的功能字符串，将消息任务扔到对应的队列中业务场景

5. topic 主题模式（路由模式的一种）

   ![img](https://tva1.sinaimg.cn/large/e6c9d24ely1h0f60mf8vrj20dj04d74g.jpg)

   - 星号井号代表通配符
   - 星号代表多个单词,井号代表一个单词
   - 路由功能添加模糊匹配
   - 消息产生者产生消息,把消息交给交换机
   - 交换机根据key的规则模糊匹配到对应的队列,由队列的监听消费者接收消息消费

6. RPC（后续补充）

**RabbitMQ 安装**

*macOS 演示*

使用命令`brew install rabbitmq`，如果没有安装erlang 会顺便一起安装。

![image-20220319142700911](https://tva1.sinaimg.cn/large/e6c9d24ely1h0f6bu9fehj20lw0kddkt.jpg)

**配置环境变量**

1. 编辑文件

   ```bash
   vim ~/.bash_profile
   ```

2. 添加 rabbitmq 地址

   ```base
   # rabbitmq
   PATH_RABBITMQ='/usr/local/Cellar/rabbitmq/3.9.13/sbin'
   export PATH=$PATH:$PATH_RABBITMQ
   ```

3. 使配置生效

   ```BASE
   source ~/.bash_profile
   ```

**常见用法**

- 启动 RabbitMQ

  ```base
  rabbitmq-server 
  ```

  ![image-20220319143333330](https://tva1.sinaimg.cn/large/e6c9d24ely1h0f6inkxpgj20mv0fxn02.jpg)

  启动成功后，可以访问`http://localhost:15672/`，初始账号密码：`gusts/guest`

  ![image-20220319143742416](https://tva1.sinaimg.cn/large/e6c9d24ely1h0f6mzowypj20zl0kn40m.jpg)

**simple模式**

- 消息产生者将消息放入队列
- 消息的消费者监听(while) 消息队列,如果队列中有消息,就消费掉,消息被拿走后,自动从队列中删除(隐患 消息可能没有被消费者正确处理,已经从队列中消失了,造成消息的丢失)应用场景:聊天(中间有一个过度的服务器;p端,c端)

做 simple 简单模式之前我们要先建一个`virture host`，并且给它分配一个用户名，用来隔离数据，根据自己的需要自行创建。（需要在后台创建用户分配权限，具体自行百度）

**代码逻辑**

- 目录结构

![image-20220322193916193](https://tva1.sinaimg.cn/large/e6c9d24ely1h0iw7pgd4hj20fe08ymxk.jpg)

- rabbitmq.go

  ```go
  package rabbitMQ
  
  import (
  	"fmt"
  	"log"
  
  	"github.com/streadway/amqp"
  )
  
  const MQURL = "amqp://simpleU:123456@127.0.0.1:5672/simple"
  
  // rabbitMQ 结构体
  type RabbitMQ struct {
  	conn      *amqp.Connection
  	channel   *amqp.Channel
  	QueueName string // 队列名称
  	Exchange  string // 交换机名称
  	Key       string // bind key 名称
  	Mqurl     string // 连接信息
  }
  
  // 结构体实例
  func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
  	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
  }
  
  // 关闭 channel 和 connection
  func (r *RabbitMQ) Destroy() {
  	_ = r.channel.Close()
  	_ = r.conn.Close()
  }
  
  // 错误处理函数
  func (*RabbitMQ) failOnError(err error, msg string) {
  	if err != nil {
  		log.Fatalf("%s:%s\n", msg, err)
  		panic("error:" + msg)
  	}
  }
  
  // 创建简单模式下的 RabbitMQ
  func NewRabbitMQSimple(queueName string) *RabbitMQ {
  	// 创建 RabbitMQ 实例
  	rabbitMQ := NewRabbitMQ(queueName, "", "")
  	var err error
  	// 获取 connection
  	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
  	rabbitMQ.failOnError(err, "failed to connect rabbitmq!")
  	// 获取 channel
  	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
  	rabbitMQ.failOnError(err, "failed to open channel!")
  	return rabbitMQ
  }
  
  // 直接模式队列生产
  func (r *RabbitMQ) PublishSimple(msg string) {
  	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
  	_, err := r.channel.QueueDeclare(
  		r.QueueName,
  		false, //是否持久化
  		false, //是否自动删除
  		false, //是否具有排他性
  		false, //是否阻塞处理
  		nil,   //额外的属性
  	)
  	if err != nil {
  		fmt.Println(err)
  	}
  	// 调用 channel，发送消息到队列中
  	r.channel.Publish(
  		r.Exchange,
  		r.QueueName,
  		false, //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
  		false, //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
  		amqp.Publishing{
  			ContentType: "text/plain",
  			Body:        []byte(msg),
  		})
  }
  
  // simple 模式下的消费者
  func (r *RabbitMQ) Consumersimple() {
  	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
  	q, err := r.channel.QueueDeclare(
  		r.QueueName,
  		false, //是否持久化
  		false, //是否自动删除
  		false, //是否具有排他性
  		false, //是否阻塞处理
  		nil,   //额外的属性
  	)
  	if err != nil {
  		fmt.Println(err)
  	}
  	// 接收消息
  	msgs, err := r.channel.Consume(
  		q.Name, // queue
  		//用来区分多个消费者
  		"", // consumer
  		//是否自动应答
  		true, // auto-ack
  		//是否独有
  		false, // exclusive
  		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
  		false, // no-local
  		//列是否阻塞
  		false, // no-wait
  		nil,   // args
  	)
  	if err != nil {
  		fmt.Println(err)
  	}
  	forever := make(chan bool)
  	// 启用协程处理消息
  	go func() {
  		for m := range msgs {
  			//消息逻辑处理，可以自行设计逻辑
  			log.Printf("Received a message: %s", m.Body)
  		}
  	}()
  	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  	<-forever
  }
  ```

- mainSimlpePublish.go

  ```go
  package main
  
  import (
  	"fmt"
  	"go/github.io/2zyyyyy/chineseDocumentation/MQ/rabbitMQ"
  )
  
  func main() {
  	rabbitMQ := rabbitMQ.NewRabbitMQSimple("" + "simple")
  	rabbitMQ.PublishSimple("Hello MQ!")
  	fmt.Println("发送成功~")
  }
  ```

- mainSimpleRecieve.go代码

  ```go
  package main
  
  import "go/github.io/2zyyyyy/chineseDocumentation/MQ/rabbitMQ"
  
  func main() {
  	rabbitmq := rabbitMQ.NewRabbitMQSimple("" + "simpleU")
  	rabbitmq.ConsumerSimple()
  }
  ```

**work 模式**

**publish 模式**

**routing 模式**

**topic 模式**

以上暂时略过

****

#### 6、Go 操作ElasticSearch

**ElasticSearch介绍**

- 介绍

  Elasticsearch（ES）是一个基于Lucene构建的开源、分布式、RESTful接口的全文搜索引擎。Elasticsearch还是一个分布式文档数据库，其中每个字段均可被索引，而且每个字段的数据均可被搜索，ES能够横向扩展至数以百计的服务器存储以及处理PB级的数据。可以在极短的时间内存储、搜索和分析大量的数据。通常作为具有复杂搜索场景情况下的核心发动机

- ElasticSearch能做什么

  - 当你经营一家网上商店，你可以让你的客户搜索你卖的商品。在这种情况下，你可以使用ElasticSearch来存储你的整个产品目录和库存信息，为客户提供精准搜索，可以为客户推荐相关商品。
  - 当你想收集日志或者交易数据的时候，需要分析和挖掘这些数据，寻找趋势，进行统计，总结，或发现异常。在这种情况下，你可以使用Logstash或者其他工具来进行收集数据，当这引起数据存储到ElasticsSearch中。你可以搜索和汇总这些数据，找到任何你感兴趣的信息
  - 对于程序员来说，比较有名的案例是GitHub，GitHub的搜索是基于ElasticSearch构建的，在github.com`/search`页面，你可以搜索项目、用户、issue、pull request，还有代码。共有`40~50`个索引库，分别用于索引网站需要跟踪的各种数据。虽然只索引项目的主分支（master），但这个数据量依然巨大，包括20亿个索引文档，30TB的索引文件

**ElasticSearch基本概念**

- Near Realtime(NRT) 几乎实时

  Elasticsearch是一个几乎实时的搜索平台。意思是，从索引一个文档到这个文档可被搜索只需要一点点的延迟，这个时间一般为毫秒级

- Cluster 集群

  群集是一个或多个节点（服务器）的集合， 这些节点共同保存整个数据，并在所有节点上提供联合索引和搜索功能。一个集群由一个唯一集群ID确定，并指定一个集群名（默认为“elasticsearch”）。该集群名非常重要，因为节点可以通过这个集群名加入群集，一个节点只能是群集的一部分

  确保在不同的环境中不要使用相同的群集名称，否则可能会导致连接错误的群集节点。例如，你可以使用logging-dev、logging-stage、logging-prod分别为开发、阶段产品、生产集群做记录

- Node 节点

  节点是单个服务器实例，它是群集的一部分，可以存储数据，并参与群集的索引和搜索功能。就像一个集群，节点的名称默认为一个随机的通用唯一标识符（UUID），确定在启动时分配给该节点。如果不希望默认，可以定义任何节点名。这个名字对管理很重要，目的是要确定你的网络服务器对应于你的ElasticSearch群集节点

  我们可以通过群集名配置节点以连接特定的群集。默认情况下，每个节点设置加入名为“elasticSearch”的集群。这意味着如果你启动多个节点在网络上，假设他们能发现彼此都会自动形成和加入一个名为“elasticsearch”的集群

  在单个群集中，你可以拥有尽可能多的节点。此外，如果“elasticsearch”在同一个网络中，没有其他节点正在运行，从单个节点的默认情况下会形成一个新的单节点名为”elasticsearch”的集群

- Index 索引

  索引是具有相似特性的文档集合。例如，可以为客户数据提供索引，为产品目录建立另一个索引，以及为订单数据建立另一个索引。索引由名称（必须全部为小写）标识，该名称用于在对其中的文档执行索引、搜索、更新和删除操作时引用索引。在单个群集中，你可以定义尽可能多的索引

- Type 类型

  在索引中，可以定义一个或多个类型。类型是索引的逻辑类别/分区，其语义完全取决于你。一般来说，类型定义为具有公共字段集的文档。例如，假设你运行一个博客平台，并将所有数据存储在一个索引中。在这个索引中，你可以为用户数据定义一种类型，为博客数据定义另一种类型，以及为注释数据定义另一类型

- Document文档

  文档是可以被索引的信息的基本单位。例如，你可以为单个客户提供一个文档，单个产品提供另一个文档，以及单个订单提供另一个文档。本文件的表示形式为JSON（JavaScript Object Notation）格式，这是一种非常普遍的互联网数据交换格式

  在索引/类型中，你可以存储尽可能多的文档。请注意，尽管文档物理驻留在索引中，文档实际上必须索引或分配到索引中的类型

- Shards & Replicas分片与副本

  索引可以存储大量的数据，这些数据可能超过单个节点的硬件限制。例如，十亿个文件占用磁盘空间1TB的单指标可能不适合对单个节点的磁盘或可能太慢服务仅从单个节点的搜索请求

  为了解决这一问题，Elasticsearch提供细分你的指标分成多个块称为分片的能力。当你创建一个索引，你可以简单地定义你想要的分片数量。每个分片本身是一个全功能的、独立的“指数”，可以托管在集群中的任何节点

  - Shards分片的重要性主要体现在以下两个特征：
    - 副本为分片或节点失败提供了高可用性。为此，需要注意的是，一个副本的分片不会分配在同一个节点作为原始的或主分片，副本是从主分片那里复制过来的
    - 副本允许用户扩展你的搜索量或吞吐量，因为搜索可以在所有副本上并行执行

- ES基本概念与关系型数据库的比较

  | ES 概念                                        | 关系型数据库       |
  | ---------------------------------------------- | ------------------ |
  | Index（索引）支持全文检索                      | Database（数据库） |
  | Type（类型）                                   | Table（表）        |
  | Document（文档），不同文档可以有不同的字段集合 | Row（数据行）      |
  | Field（字段）                                  | Column（数据列）   |
  | Mapping（映射）                                | Schema（模式）     |

**ElasticSearch安装**

- 安装

  - macOS 使用命令：`brew install elasticsearch`

- 启动

  - 使用命令：`elasticsearch`

- 校验

  - 浏览器访问地址：`http://127.0.0.1:9200/`

    ![image-20220324201002194](https://tva1.sinaimg.cn/large/e6c9d24ely1h0l8cdmqf3j20do09f74z.jpg)

**Kinaba安装**

- 介绍

  - Kibana是一个开源的分析和可视化平台，设计用于和Elasticsearch一起工作。

    你可以使用Kibana来搜索、查看、并和存储在Elasticsearch索引中的数据进行交互。

    你可以轻松地执行高级数据分析，并且以各种图标、表格和地图的形式可视化数据。

    Kibana使得理解大量数据变得很容易。它简单的、基于浏览器的界面使你能够快速创建和共享动态仪表板，实时显示Elasticsearch查询的变化。

- 安装

  - macOS 使用命令：`brew install kibana`

- 启动

  - 使用命令：`kibana`

    ![image-20220324201428774](https://tva1.sinaimg.cn/large/e6c9d24ely1h0l8h0vselj20ne09wad7.jpg)

- 校验

  - 浏览器访问：`http://localhost:5601`

    ![image-20220324201526746](https://tva1.sinaimg.cn/large/e6c9d24ely1h0l8hx308mj210g0nbmyt.jpg)







































