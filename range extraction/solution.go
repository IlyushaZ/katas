package kata

import (
	"strconv"
)

// THAT'S THE WORST SOLUTION YOU HAVE EVER SEEN
func Solution(list []int) string {
	var result string

	for i := 0; i < len(list); i++ {
		result += strconv.Itoa(list[i])

		v := list[i] + 1
		var j int
		for j = i + 1; j < len(list) && list[j] == v; j++ {
			v++
		}

		if j - i > 2 {
			result += "-" + strconv.Itoa(v-1)
			i = j - 1
		}

		result += ","
	}

	return result[:len(result)-len(",")]
}
