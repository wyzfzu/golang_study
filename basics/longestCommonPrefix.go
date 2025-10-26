package basics

import "fmt"

func longestCommonPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}

	var prefix = s[0]
	for i := 1; i < len(s); i++ {
		prefix = commonPrefix(prefix, s[i])
		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}

func commonPrefix(s1 string, s2 string) string {
	minLen := min(len(s1), len(s2))
	idx := 0
	for idx < minLen && s1[idx] == s2[idx] {
		idx++
	}

	return s1[:idx]
}

func TestLongestCommonPrefix() {
	strs := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"cir", "car"},
	}

	for _, v := range strs {
		fmt.Println("longestCommonPrefix, strs=", v, ", prefix=", longestCommonPrefix(v))
	}
}
