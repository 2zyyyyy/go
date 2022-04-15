package main

import (
	"fmt"
	"reflect"
)

// 结构体与反射

type User struct {
	Name    string `json:"jsonName" db:"dbName"`
	Id, Age int
}

// Boy 匿名字段
type Boy struct {
	User
	Addr string
}

// Hello 结构体方法
func (u User) Hello() {
	fmt.Println("Hello")
}

// Run 结构体方法
func (u User) Run(name string) {
	fmt.Println(u.Name, "想要润了")
}

// Poni 传入interface{}
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)           // 类型： main.User
	fmt.Println("字符串类型：", t.Name()) // 字符串类型： User
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v) // {1001 40 月满轩尼诗}
	// 可以获取所有属性
	// 获取结构体字段个数：t.NameField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v\n", f.Name, f.Type) // Id : int
		// 获取字段值信息
		// Interface()：获取字段对应的值
		value := v.Field(i).Interface()
		fmt.Println("value:", value) // value: 1001
	}
	fmt.Println("======================方法=====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name) // Hello
		fmt.Println(m.Type) // func(main.User)
	}
}

func anonymousField() {
	m := Boy{
		User: User{"月满轩尼诗", 1001, 40},
		Addr: "浙江省杭州市西湖区石马新村",
	}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	for i := 0; i < t.NumField(); i++ {
		// anonymous：匿名
		fmt.Printf("%v\n", t.Field(i))
		// 值信息
		fmt.Printf("%v\n", reflect.ValueOf(m).Field(i))
	}
}

// 修改结构体的值
func setValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("wanli")
	}
}

// 调用方法
func callMethods() {
	user := User{
		Id:   1002,
		Age:  24,
		Name: "仙道",
	}
	v := reflect.ValueOf(user)
	// 获取方法
	m := v.MethodByName("Run")
	// 构建参数
	args := []reflect.Value{reflect.ValueOf("柱子哥")}
	// 没参数的情况下：var arg2 []reflect.Value
	// 调用方法,需传入方法的参数
	m.Call(args)
}

// 获取字段tag
func getFieldTag() {
	var s User
	v := reflect.ValueOf(&s)
	// 类型
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}

func main() {
	//user := User{1001, 40, "月满轩尼诗"}
	//Poni(user)
	//anonymousField()
	//setValue(&user)
	//fmt.Printf("%#v\n", user)
	//callMethods()
	getFieldTag()
}
