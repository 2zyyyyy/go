package split

import (
	"reflect"
	"testing"
)

// 测试组

func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义一个存储测试用例数据的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "天亦有情天易老", sep: "天", want: []string{"亦有情", "易老"}},
	}

	// 遍历切片，循环执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}
}
