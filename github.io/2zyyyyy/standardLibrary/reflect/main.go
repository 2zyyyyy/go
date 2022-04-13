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

func main() {
	var x float64 = 3.141
	//reflectType(x)
	reflectValue(x)
}
