package longest_common_prefix

import (
	"fmt"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	m := make(map[string]struct{})
	for i := 0; i < len(strs[0]); i++ {
		subStr := strs[0][:i+1]
		m[subStr] = struct{}{}
	}
	fmt.Println(m)

	for k := range m {
		for _, str := range strs {
			if !strings.HasPrefix(str, k) {
				delete(m, k)
			}
		}
	}

	if len(m) == 0 {
		return ""
	}

	maxLength := 0
	var maxLengthKey string
	for k := range m {
		if len(k) > maxLength {
			maxLength = len(k)
			maxLengthKey = k
		}
	}

	return maxLengthKey
}
