package main

import (
	"fmt"
	"reflect"
)

// 反射获取interface类型信息
func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	default:
		fmt.Printf("a is other type\n")
	}
}

// 反射获取interface值信息
func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is", v.Float())
	}
}

// 反射修改值信息
func reflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(5.8)
		fmt.Println("a is", v.Float())
	case reflect.Ptr:
		// elem()获取地址指向的值
		v.Elem().SetFloat(6.3)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}

func main() {
	var x float64 = 3.141
	//reflectType(x)
	//reflectValue(x)
	// 反射认为下面是指针类型 不是float64类型
	reflectSetValue(&x)
	fmt.Println("main:", x)
}
