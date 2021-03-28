package split

import "strings"

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(input, sep string) []string {
	// 提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。
	//这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配。
	var res = make([]string, 0, strings.Count(input, sep)+1)
	// Index 返回字符串sep在s中的索引，-1表示字符串s不包含sep
	index := strings.Index(input, sep)
	for index > -1 {
		res = append(res, input[:index])
		input = input[index+len(sep):] // 这里使用len(sep)获取sep的长度
		index = strings.Index(input, sep)
	}
	res = append(res, input)
	return res
}
