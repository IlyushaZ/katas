package kata

func MaximumSubarraySum(numbers []int) int {
	var maxVal int
	var buf int

	for _, val := range numbers {
		buf += val
		maxVal = max(maxVal, buf)

		if buf < 0 { buf = 0 }
	}

	return maxVal
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}