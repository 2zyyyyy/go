package split

import "strings"

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) (res []string) {
	// Index 返回字符串sep在s中的索引，-1表示字符串s不包含sep
	i := strings.Index(s, sep)
	for i > -1 {
		res = append(res, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	res = append(res, s)
	return res
}
