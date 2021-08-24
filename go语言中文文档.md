## Go 语言中文文档记录

### 一、Go基础

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



#### 8、基本类型

#### 9、数据Array

#### 10、切片Slice

#### 11、Slice底层实现

#### 12、指针

#### 13、Map

#### 14、Map实现原理

#### 15、结构体























