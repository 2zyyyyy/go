package split

import (
	"fmt"
	"strings"
)

// golang unit testing demo

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		fmt.Printf("s=%s, i=%d, result=%s\n", s, i, result)
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}