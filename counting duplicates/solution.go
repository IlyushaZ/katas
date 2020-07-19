package kata

import "strings"

func duplicateCount(s1 string) int {
	var res int
	chars := make(map[string]int)

	for _, val := range strings.Split(strings.ToLower(s1), "") {
		chars[val] += 1
		if chars[val] == 2 {
			res++
		}
	}

	return res
}
