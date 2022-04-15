package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MarshalFile 序列化和反序列化
// 序列化数据到指定的文件
func MarshalFile(filename string, data interface{}) (err error) {
	// 1.数据的序列化
	result, err := Marshal(data)
	if err != nil {
		return
	}
	// 2.将序列化好的数据，写出到filename
	return ioutil.WriteFile(filename, result, 0666)
}

// Marshal 序列化的方法
// 传入的结构体-> []byte
// 基本思路：反射解析出传入的数据 转字节切片
func Marshal(data interface{}) (result []byte, err error) {
	// 获取类型
	typeInfo := reflect.TypeOf(data)
	valueInfo := reflect.ValueOf(data)
	// 判断类型
	if typeInfo.Kind() != reflect.Struct {
		return
	}

	var conf []string

	// 获取所有字段去处理
	for i := 0; i < typeInfo.NumField(); i++ {
		// 取字段
		labelField := typeInfo.Field(i)
		// 取值
		labelValue := valueInfo.Field(i)
		// 获取字段的类型
		fieldType := labelField.Type
		//判断字段的类型
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		// 拼的是[server]和[mysql]
		// 获取tag
		tagValue := labelField.Tag.Get("yaml")
		if len(tagValue) == 0 {
			tagValue = labelField.Name
		}
		label := fmt.Sprintf("\n[%s]\n", tagValue)
		conf = append(conf, label)
		// 拼K-V
		for j := 0; j < fieldType.NumField(); j++ {
			// 这里渠道的是大写
			keyField := fieldType.Field(j)
			// 取tag
			fieldTagValue := keyField.Tag.Get("yaml")
			if len(fieldTagValue) == 0 {
				fieldTagValue = keyField.Name
			}
			//取值
			valueField := labelValue.Field(j)
			// interface()取真正对应的值
			item := fmt.Sprintf("%s=%v\n", fieldTagValue, valueField.Interface())
			conf = append(conf, item)
		}
	}
	// 遍历切片转类型
	for _, value := range conf {
		byteValue := []byte(value)
		result = append(result, byteValue...)
	}
	return
}

// UnMarshalFile 文件读取数据，做反序列化
func UnMarshalFile(filename string, result interface{}) (err error) {
	// 1.文件读取
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// 2.进行反序列化
	return UnMarshal(data, result)
}

// 反序列化
// []byte -> struct

func UnMarshal(data []byte, result interface{}) (err error) {
	// 判断是否是指针
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		return
	}
	// 判断是否是结构体
	if typeInfo.Elem().Kind() != reflect.Struct {
		return
	}
	// 转类型，按行切割
	lineArray := strings.Split(string(data), "\n")
	// 定义全局标签名 也就是server和mysql
	var myFileName string

	for _, line := range lineArray {
		// 各种判断
		line = strings.TrimSpace(line)
		// 处理文档中有空行或注释的情况
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}
		// 按照括号去判断
		if line[0] == '[' {
			// 当做大标签处理
			myFileName, err = myLabel(line, typeInfo.Elem())
			if err != nil {
				return
			}
			continue
		}
		// 按正常数据处理
		err = myField(myFileName, line, result)
		if err != nil {
			return
		}
	}
	return
}

// 解析属性
// 参数：大标签名，行数据，对象
func myField(fieldName string, line string, result interface{}) (err error) {
	fmt.Println(line)
	key := strings.TrimSpace(line[0:strings.Index(line, "=")])
	val := strings.TrimSpace(line[strings.Index(line, "=")+1:])
	// 解析到结构体
	//resultType := reflect.TypeOf(result)
	resultValue := reflect.ValueOf(result)
	// 拿到字段值，这里直接设置不知道类型
	labelValue := resultValue.Elem().FieldByName(fieldName)
	// 拿到该字段类型
	fmt.Println(labelValue)
	labelType := labelValue.Type()
	// 第一次进来应该是server
	// 存放取到的字段名
	var keyName string
	// 遍历server结构体的所有字段
	for i := 0; i < labelType.NumField(); i++ {
		// 获取结构体字段
		field := labelType.Field(i)
		tagVal := field.Tag.Get("yaml")
		if tagVal == key {
			keyName = field.Name
			break
		}
	}

	// 给字段赋值
	// 取字段值
	filedValue := labelValue.FieldByName(keyName)
	// 修改值
	switch filedValue.Type().Kind() {
	case reflect.String:
		filedValue.SetString(val)
	case reflect.Int:
		i, err2 := strconv.ParseInt(val, 10, 64)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		filedValue.SetInt(i)
	case reflect.Uint:
		i, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			fmt.Println(err)
			return err
		}
		filedValue.SetUint(i)
	case reflect.Float32:
		f, _ := strconv.ParseFloat(val, 64)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		filedValue.SetFloat(f)
	}
	return
}

// 处理大标签
func myLabel(line string, typeInfo reflect.Type) (fileName string, err error) {
	// 去字符串头尾
	labelName := line[1 : len(line)-1]
	// 循环去结构体找tag，对应上才能解析
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		tagValue := field.Tag.Get("yaml")
		// 判断tag
		if labelName == tagValue {
			fileName = field.Name
			break
		}
	}
	return
}
