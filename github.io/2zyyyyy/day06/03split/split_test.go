package split

import (
	"reflect"
	"testing"
)

func TestSplit01(t *testing.T) { // 测试函数必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")      // 程序输出的结果
	want := []string{"a", "b", "c"} // 期望结果
	// 因为slice不能直接比较，故借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		// 测试未通过输出
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func TestSplit02(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}
