package split

import "strings"

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(input, sep string) (res []string) {
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
