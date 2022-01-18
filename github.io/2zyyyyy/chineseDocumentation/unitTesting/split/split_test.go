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