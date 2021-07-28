### 基础语法

#### 1、关键字

- 25个关键字和保留字

  | break    | default     | func   | interface | select |
  | -------- | ----------- | ------ | --------- | ------ |
  | case     | defer       | go     | map       | struct |
  | chan     | else        | goto   | package   | switch |
  | const    | fallthrough | if     | range     | type   |
  | continue | for         | import | return    | var    |

- 36个预定义标识符

  | append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
  | ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
  | copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
  | int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
  | print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

### 2、变量

- 省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：

  ```go
  v_name := value
  // 例如
  var intVal int 
  intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
  ```

  

- 变量声明没有初始化默认零值

  - 数值类型（包括complex64/128）为 **0**

  - 布尔类型为 **false**

  - 字符串为 **""**（空字符串）

  - 以下几种类型为 **nil**：

    ```go
    var a *int
    var a []int
    var a map[string] int
    var a chan int
    var a func(string) int
    var a error // error 是接口
    ```

- 局部变量却没有在相同的代码块中使用它，会得到编译错误。但是全局变量是允许声明但不使用的。

### 3、常量

- 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

  ```go
  // 常量的定义格式
  const identifier [type] = value
  ```

- 你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
  - 显式类型定义： `const b string = "abc"`
  - 隐式类型定义： `const b = "abc"`
- iota
  - iota，特殊常量，可以认为是一个可以被编译器修改的常量。
  - iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

### 4、运算符

- 使用指针变量与不使用的区别：

- ```go
  func main(){
      var a int = 4
      var ptr int
      ptr = a 
      fmt.Println(ptr)//4
      a = 15
      fmt.Println(ptr)//4
      
      var b int = 5 
      var ptr1 *int
      ptr1 = &b 
      fmt.Println(*ptr1)//5
      b=15 
      fmt.Println(*ptr1)//15
  }
  ```

### 5、条件语句

​	条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句

​	Go 语言提供了以下几种条件判断语句：

| 语句           | 描述                                                         |
| :------------- | :----------------------------------------------------------- |
| if 语句        | **if 语句** 由一个布尔表达式后紧跟一个或多个语句组成。       |
| if...else 语句 | **if 语句** 后可以使用可选的 **else 语句**, else 语句中的表达式在布尔表达式为 false 时执行。 |
| if 嵌套语句    | 你可以在 **if** 或 **else if** 语句中嵌入一个或多个 **if** 或 **else if** 语句。 |
| switch 语句    | **switch** 语句用于基于不同条件执行不同动作。                |
| select 语句    | **select** 语句类似于 **switch** 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。 |

### 6、循环语句

​	Go 语言提供了以下几种类型循环处理语句：

| 循环类型 | 描述                                 |
| :------- | :----------------------------------- |
| for 循环 | 重复执行语句块                       |
| 循环嵌套 | 在 for 循环中嵌套一个或多个 for 循环 |

- 循环控制语句

  - 循环控制语句可以控制循环体内语句的执行过程。GO 语言支持以下几种循环控制语句：

  - | 控制语句      | 描述                                             |
    | :------------ | :----------------------------------------------- |
    | break 语句    | 经常用于中断当前 for 循环或跳出 switch 语句      |
    | continue 语句 | 跳过当前循环的剩余语句，然后继续进行下一轮循环。 |
    | goto 语句     | 将控制转移到被标记的语句。                       |

- 无限循环

  - 如果循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：

  ```go
  package main
  
  import "fmt"
  
  func main() {
      for true  {
          fmt.Printf("这是无限循环。\n");
      }
  }
  ```



### 7、函数

​	函数是基本的代码块，用于执行一个任务。Go 语言最少有个 main() 函数。你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。函数声明告诉了编译器函数的名称，返回类型，和参数。

- 函数定义

  ```go
  func function_name( [parameter list] ) [return_types] {
     函数体
  }
  ```
  - 函数定义解析：
    - func：函数由 func 开始声明
    - function_name：函数名称，参数列表和返回值类型构成了函数签名。
    - parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
    - return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
    - 函数体：函数定义的代码集合。

- 函数调用

  - 当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。

- 函数参数

  - 函数如果使用参数，该变量可称为函数的形参。形参就像定义在函数体内的局部变量。调用函数，可以通过两种方式来传递参数：

    | 传递类型 | 描述                                                         |
    | :------- | :----------------------------------------------------------- |
    | 值传递   | 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。 |
    | 引用传递 | 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。 |

    默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

- 函数用法

  | 函数用法                   | 描述                                     |
  | :------------------------- | :--------------------------------------- |
  | 函数作为另外一个函数的实参 | 函数定义后可作为另外一个函数的实参数传入 |
  | 闭包                       | 闭包是匿名函数，可在动态编程中使用       |
  | 方法                       | 方法就是一个包含了接受者的函数           |

### 8、变量作用域

​	作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。Go 语言中变量可以在三个地方声明：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

​    Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：

```go
package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
   /* 声明局部变量 */
   var g int = 10

   fmt.Printf ("结果： g = %d\n",  g)
}

// 实际输出：g = 10
```

- 初始化局部和全局变量

| 数据类型 | 初始化默认值 |
| :------- | :----------- |
| int      | 0            |
| float32  | 0            |
| pointer  | nil          |

### 9、数组

- 声明数组

  Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

  ```go
  var variable_name [SIZE] variable_type
  ```

  以上为一维数组的定义方式。例如以下定义了数组 balance 长度为 10 类型为 float32：

  ```go
  var balance [10] float32
  ```

- 初始化数组

  ```go
  var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
  
  balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
  ```

  如果数组长度不确定，可以使用 **...** 代替数组的长度，编译器会根据元素个数自行推断数组的长度：

  ```go
  var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
  
  balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
  ```

  如果设置了数组的长度，我们还可以通过指定下标来初始化元素：

  ```go
  //  将索引为 1 和 3 的元素初始化
  balance := [5]float32{1:2.0,3:7.0}
  ```

- 数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

  ```go
  // 定义变量取salary balance[]中第十个元素
  var salary float32 = balance[9]
  ```

### 10、多维数组

​		Go 语言支持多维数组，以下为常用的多维数组声明方式：

```go
var variable_name [size 1][size 2]...[size n] variable_type
```

​		以下实例声明了三维的整型数组：

```go
var thredim [5][10][4]int
```

- 二维数组

  二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：

  ```go
  // variable_type 为 Go 语言的数据类型，arrayName 为数组名，二维数组可认为是一个表格，x 为行，y 为列
  var arrayName [x][y] variable_type
  ```

  二维数组中的元素可通过`a[i][j]`来访问。实例：

  ```go
  // 多维数组
  func multidimensional_array() {
  	// Step 1: 创建数组
  	array := [][]int{}
  
  	// Step 2: 使用 appped() 函数向空的二维数组添加两行一维数组
  	row1 := []int{1, 2, 3}
  	row2 := []int{4, 5, 6}
  	array = append(array, row1)
  	array = append(array, row2)
  
  	// Step 3: 打印两行数据
  	fmt.Println("row1:", array[0])
  	fmt.Println("row2:", array[1])
  
  	// Step 4: 访问第最后一个元素
  	fmt.Println(array[1][2])
  }
  ```

- 初始化二维数组

  多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组：

  ```go
  a := [3][4]int{  
   {0, 1, 2, 3} ,   /*  第一行索引为 0 */
   {4, 5, 6, 7} ,   /*  第二行索引为 1 */
   {8, 9, 10, 11},   /* 第三行索引为 2 */
  }
  ```

- 访问二维数组

  二维数组通过指定坐标来访问。如数组中的行索引与列索引，例如：

  ```go
  val := a[2][3]
  或
  var value int = a[2][3]
  ```

  以上实例访问了二维数组 val 第三行的第四个元素。二维数组可以使用循环嵌套来输出元素：

  ```go
  func forPrintArray() {
  	/* 数组 - 5 行 2 列*/
  	array := [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
  
  	/* 输出数组元素 */
  	for i := 0; i < 5; i++ {
  		for j := 0; j < 2; j++ {
  			fmt.Printf("a[%d][%d] = %d\n", i, j, array[i][j])
  		}
  	}
  }
  ```

### 11、向函数传递数组

​	如果你想向函数传递数组参数，你需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明：

- 方式一：形参设置数组大小

  - ```go
    void myFunction(param [10]int) {
      ...
    }
    ```

- 方式二：形参未设置数组大小

  - ```go
    void myFunction(param []int) {
      ...
    }
    ```

- 实例

  ​	让我们看下以下实例，实例中函数接收整型数组参数，另一个参数指定了数组元素的个数，并返回平均值：

  ```go
  // 向函数传递数组
  func getAverage(array []int, size int) float32 {
  	var i, sum int
  	var avg float32
  	for i = 0; i < size; i++ {
  		sum += array[i]
  	}
  	avg = float32(sum) / float32(size)
  	return avg
  }
  
  func main() {
    /* 数组长度为 5 */
  	array := [5]int{1000, 2, 3, 17, 50}
  
  	/* 数组作为参数传递给函数 */
  	avg := getAverage(array[:], 5)
  	/* 输出返回的平均值 */
  	fmt.Printf("平均值为：%f\n", avg)
  }
  ```

  - Go 语言的数组是值，其长度是其类型的一部分，作为函数参数时，是 **值传递**，函数中的修改对调用者不可见

  - Go 语言中对数组的处理，一般采用 **切片** 的方式，切片包含对底层数组内容的引用，作为函数参数时，类似于 **指针传递**，函数中的修改对调用者可见

### 12、指针

​	我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。以下实例演示了变量在内存中地址：

```go
package main

import "fmt"

func main() {
   var a int = 10  

   fmt.Printf("变量的地址: %x\n", &a  )
}
```

- 什么是指针

  一个指针变量指向了一个值的内存地址。类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

```go
var var_name *var-type
```

​	var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

```go
var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```

- 如何使用指针

  指针使用流程：

  - 定义指针变量。
  - 为指针变量赋值。
  - 访问指针变量中指向地址的值。

  在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

  ```go
  package main
  
  import "fmt"
  
  func main() {
     var a int= 20   /* 声明实际变量 */
     var ip *int        /* 声明指针变量 */
  
     ip = &a  /* 指针变量的存储地址 */
  
     fmt.Printf("a 变量的地址是: %x\n", &a  )
  
     /* 指针变量的存储地址 */
     fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
  
     /* 使用指针访问值 */
     fmt.Printf("*ip 变量的值: %d\n", *ip )
  }
  ```

  ![image-20210713212857866](https://tva1.sinaimg.cn/large/008i3skNly1gsfn84jcnnj30tg04wt9o.jpg)

- 空指针

  当一个指针被定义后没有分配到任何变量时，它的值为 nil。

  nil 指针也称为空指针。

  nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

  一个指针变量通常缩写为 ptr。

  查看以下实例：

  ```go
  package main
  
  import "fmt"
  
  func main() {
     var  ptr *int
  
     fmt.Printf("ptr 的值为 : %x\n", ptr  )
  }
  ```

### 13、指针数组

​	在我们了解指针数组前，先看个实例，定义了长度为 3 的整型数组：

```go
// 指针数组
const MAX int = 3
func ptrArray() {
	array := []int{10, 100, 200}
	for i := 0; i < MAX; i++ {
		fmt.Printf("array[%d]=%d\n", i, array[i])
	}
}
```

​	有一种情况，我们可能需要保存数组，这样我们就需要使用到指针。以下声明了整型指针数组：

```go
var ptr [MAX]*int
```

​	ptr 为整型指针数组。因此每个元素都指向了一个值。以下实例的三个整数将存储在指针数组中：

```go
// 指针数组
func ptrArray() {
	array := []int{10, 100, 200}
	var ptr [MAX]*int

	// 循环赋值(将array的地址赋值给ptr)
	for i := 0; i < MAX; i++ {
		ptr[i] = &array[i]
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("ptr[%d]=%d\n", i, *ptr[i])
	}
}
```



### 14、指向指针的指针

​	如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：

![img](https://tva1.sinaimg.cn/large/008i3skNly1gswpwf21suj30bi01tweb.jpg)

​	指向指针的指针变量声明格式如下：

```go
var ptr **int
```

​	以上指向指针的指针变量为整型。访问指向指针的指针变量需要用两个`*`，如下所示：

```go
// 指向指针的指针
func ptrPtrValue() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	// 指针ptr的地址
	ptr = &a

	// 指向指针pptr的地址
	pptr = &ptr

	// 获取pptr(指向指针的指针的值)
	fmt.Printf("变量a = %d\n", a)
	fmt.Printf("变量*ptr = %d\n", *ptr)
	fmt.Printf("变量**ptr = %d\n", *pptr)
}
```



### 15 向函数传递指针参数



### 16、结构体

​	Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

- 结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
  - Title ：标题	
  - Author ： 作者
  - Subject：学科
  - ID：书籍ID

- 定义结构体

  结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

  ```go
  type struct_variable_type struct {
    member definition
     member definition
     ...
     member definition
  }
  ```

  一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

  ```go
  variable_name := structure_variable_type {
    value1,
    value2,
    ...
    valuen
  }
  或
  variable_name := structure_variable_type { 
    key1: value1,
    key2: value2,
    ...,
    keyn: valuen
  }
  ```

  实例如下：

  ```go
  package main
  
  import "fmt"
  
  type Books struct {
     title string
     author string
     subject string
     book_id int
  }
  
  
  func main() {
  
      // 创建一个新的结构体
      fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
  
      // 也可以使用 key => value 格式
      fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
  
      // 忽略的字段为 0 或 空
     fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
  }
  ```

- 访问结构体成员

  ​	如果要访问结构体成员，需要使用点号`.`操作符，格式为：`结构体.成员名`

  结构体类型变量使用struct关键字定义，实例如下：

  ```go
  package main
  
  import "fmt"
  
  type Books struct {
     title string
     author string
     subject string
     book_id int
  }
  
  func main() {
     var Book1 Books        /* 声明 Book1 为 Books 类型 */
     var Book2 Books        /* 声明 Book2 为 Books 类型 */
  
     /* book 1 描述 */
     Book1.title = "Go 语言"
     Book1.author = "www.runoob.com"
     Book1.subject = "Go 语言教程"
     Book1.book_id = 6495407
  
     /* book 2 描述 */
     Book2.title = "Python 教程"
     Book2.author = "www.runoob.com"
     Book2.subject = "Python 语言教程"
     Book2.book_id = 6495700
  
     /* 打印 Book1 信息 */
     fmt.Printf( "Book 1 title : %s\n", Book1.title)
     fmt.Printf( "Book 1 author : %s\n", Book1.author)
     fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
     fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
  
     /* 打印 Book2 信息 */
     fmt.Printf( "Book 2 title : %s\n", Book2.title)
     fmt.Printf( "Book 2 author : %s\n", Book2.author)
     fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
     fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
  }
  ```

- 结构体作为函数参数

  ​	你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量：

  ```go
  package main
  
  import "fmt"
  
  type Books struct {
     title string
     author string
     subject string
     book_id int
  }
  
  func main() {
     var Book1 Books        /* 声明 Book1 为 Books 类型 */
     var Book2 Books        /* 声明 Book2 为 Books 类型 */
  
     /* book 1 描述 */
     Book1.title = "Go 语言"
     Book1.author = "www.runoob.com"
     Book1.subject = "Go 语言教程"
     Book1.book_id = 6495407
  
     /* book 2 描述 */
     Book2.title = "Python 教程"
     Book2.author = "www.runoob.com"
     Book2.subject = "Python 语言教程"
     Book2.book_id = 6495700
  
     /* 打印 Book1 信息 */
     printBook(Book1)
  
     /* 打印 Book2 信息 */
     printBook(Book2)
  }
  
  func printBook( book Books ) {
     fmt.Printf( "Book title : %s\n", book.title)
     fmt.Printf( "Book author : %s\n", book.author)
     fmt.Printf( "Book subject : %s\n", book.subject)
     fmt.Printf( "Book book_id : %d\n", book.book_id)
  }
  ```

  - 你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量：

- ```go
  func printBook( book Books ) {
     fmt.Printf( "Book title : %s\n", book.title)
     fmt.Printf( "Book author : %s\n", book.author)
     fmt.Printf( "Book subject : %s\n", book.subject)
     fmt.Printf( "Book book_id : %d\n", book.book_id)
  }
  ```

- 结构体指针

  你可以定义指向结构体的指针类似于其他指针变量，格式如下：

  ```go
  var struct_pointer *Books
  ```

  以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：

  ```go
  struct_printer = &Book1
  ```

  使用结构体指针访问结构体成员，使用`.`操作符：`struct_pointer.title`

- 结构体内属性的首字母大小写问题

  - 首字母大写相当于 Public
  - 首字母小写相当于 Private

  **注意:** 这个 public 和 private 是相对于包（go 文件首行的 package 后面跟的包名）来说的。

  **敲黑板，划重点**

  当要将结构体对象转换为 JSON 时，对象中的属性首字母必须是大写，才能正常转换为 JSON。



### 14、切片（Slice）

​	Go 语言切片是对数组的抽象。

​	Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

- 定义切片

  声明一个未指定大小的数组来定义切片。`var identifier []type`,切片不需要说明长度。或使用 **make()** 函数来创建切片:

  ```go
  var slice1 []type = make([]type, len)
  // 或者简写为：
  slice1 := make([]type, len)
  // 也可以指定容量，其中 capacity 为可选参数。
  make([]T, length, capacity)
  ```

- 切片初始化

  ```go
  // 直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。
  s := []int {1, 2, 3}
  ```

- len()和cap()函数

  ​	切片是可索引的，并且可以由 len() 方法获取长度。切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。以下为具体实例：

  ```go
  package main
  
  import "fmt"
  
  func main() {
     var numbers = make([]int,3,5)
  
     printSlice(numbers)
  }
  
  func printSlice(x []int){
     fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
  }
  ```

  ```go
  // 以上输出结果
  len=3 cap=5 slice=[0 0 0]
  ```

- 空（nil）切片

  切片定义后未初始化之前默认为nil，长度为0

  ```go
  package main
  
  import "fmt"
  
  func main() {
     var numbers []int
  
     printSlice(numbers)
  
     if(numbers == nil){
        fmt.Printf("切片是空的")
     }
  }
  
  func printSlice(x []int){
     fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
  }
  ```

- 切片截取

  可以通过设置下限及上限来设置截取切片 *[lower-bound:upper-bound]*，实例如下：

  ```go
  func slice_substring() {
  	/* 创建切片 */
  	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums), cap(nums), nums)
  	// 打印原始切片
  	fmt.Println("nums:", nums)
  
  	// [1, 4]所以1到4(不包含)
  	fmt.Println("nums[1:4]:", nums[1:4])
  
  	// 默认下限为0
  	fmt.Println("nums[:3]:", nums[:3])
  
  	// 默认上限weilen(nums)
  	fmt.Println("nums[4:]:", nums[4:])
  
  	nums_one := make([]int, 0, 5)
  	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums_one), cap(nums_one), nums_one)
  
  	// 打印子切片从索引[0, 2)
  	nums_two := nums[:2]
  	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums_two), cap(nums_two), nums_two)
  
  	// 打印索引[0, 2)
  	nums_three := nums[2:5]
  	fmt.Printf("len=%d cap=%d slice=%v\n", len(nums_three), cap(nums_three), nums_three)
  }
  
  // 输出
  len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
  numbers == [0 1 2 3 4 5 6 7 8]
  numbers[1:4] == [1 2 3]
  numbers[:3] == [0 1 2]
  numbers[4:] == [4 5 6 7 8]
  len=0 cap=5 slice=[]
  len=2 cap=9 slice=[0 1]
  len=3 cap=7 slice=[2 3 4]
  ```

- append()和copy()

  如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。

  ```go
  func slice_append_cppy() {
  	var nums []int
  	printSlice(nums)
  
  	/* 允许追加空切片 */
  	nums = append(nums, 0)
  	printSlice(nums)
  
  	/* 向切片添加一个元素 */
  	nums = append(nums, 1)
  	printSlice(nums)
  
  	/* 同时添加多个元素 */
  	nums = append(nums, 2, 3)
  	printSlice(nums)
  
  	/* 创建切片 numbers1 是之前切片的两倍容量*/
  	nums_cap_double := make([]int, len(nums), (cap(nums))*2)
  
  	/* 拷贝 numbers 的内容到 numbers1 */
  	copy(nums_cap_double, nums)
  	printSlice(nums_cap_double)
  }
  ```

  