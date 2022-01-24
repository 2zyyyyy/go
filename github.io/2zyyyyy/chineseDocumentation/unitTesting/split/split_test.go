package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	// 定义一个测试类型
	type Test struct {
		input string
		sep   string
		want  []string
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

// func TestMoreSplit(t *testing.T) {
//     got := Split("abcd", "bc")
//     want := []string{"a", "d"}
//     if !reflect.DeepEqual(want, got) {
//         t.Errorf("excepted:%v, got:%v", want, got)
//     }
// }

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}
