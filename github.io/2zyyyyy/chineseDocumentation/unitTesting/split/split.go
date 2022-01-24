package split

import (
	"strings"
)

// golang unit testing demo

func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	// fmt.Printf("strings.Count(s, sep)+1=%d\n", strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])

		// fmt.Printf("s=%s, i=%d, result=%s\n", s, i, result)
		s = s[i+(len(sep)):] // i+1会导致sep长度大于1 产生bug 需要将1改成sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
