package kata

func LongestConsec(strarr []string, k int) string {
	if k == 1 { return longestString(strarr) }

	if k <= 0 || len(strarr) == 0 { return "" }

	var longest string

	for index, value := range strarr {
		seq := value

		for i := 1; i < k; i++ {
			if len(strarr) > index + k - 1 {
				seq += strarr[index + i]
			}
		}

		if len(seq) > len(longest) { longest = seq }
	}

	if len(longest) < k { return "" }

	return longest
}

func longestString(strArr []string) string {
	longest := strArr[0]

	for _, value := range strArr {
		if len(value) > len(longest) {
			longest = value
		}
	}

	return longest
}

