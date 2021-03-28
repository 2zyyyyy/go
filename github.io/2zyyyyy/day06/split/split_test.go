package split

import (
	"reflect"
	"testing"
)

// 测试组
//go test -v -run=Split/simple只会运行simple对应的子测试用例

func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义一个存储测试用例数据的切片
	//tests := []test{
	//	{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
	//	{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
	//	{input: "abcd", sep: "bc", want: []string{"a", "d"}},
	//	{input: "天亦有情天易老", sep: "天", want: []string{"亦有情", "易老"}},
	//}

	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	// 遍历map，循环执行测试用例
	for name, tc := range tests {
		//got := Split(tc.input, tc.sep)
		//if !reflect.DeepEqual(got, tc.want) {
		//	t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
		//}
		// 使用t.run()执行子测试
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e:f:g:h:i:j:k:l:m:n", ":")
	}
}
