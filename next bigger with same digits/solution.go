package kata

import (
	"sort"
	"strconv"
	"strings"
)

func NextBigger(n int) int {
	numbers := intToSlice(n)

	for i := 1; i < len(numbers); i++ {
		nums := numbers[len(numbers)-i:]
		prev := numbers[len(numbers)-i-1]

		minBiggerIndex := getMinBigger(nums, prev)
		if minBiggerIndex == -1 {
			continue
		}

		numbers[len(numbers)-1-i] = nums[minBiggerIndex]
		nums[minBiggerIndex] = prev

		res := sliceToInt(numbers)
		if res > n {
			return res
		}
	}

	return -1
}

func getMinBigger(seq []int, n int) int {
	sort.Ints(seq)
	for i := 0; i < len(seq); i++ {
		if seq[i] > n {
			return i
		}
	}

	return -1
}

func intToSlice(n int) []int {
	s := strconv.Itoa(n)
	res := make([]int, len(s))

	for i, v := range strings.Split(s, "") {
		converted, _ := strconv.Atoi(v)
		res[i] = converted
	}

	return res
}

func sliceToInt(seq []int) int {
	strs := make([]string, len(seq))
	for i, v := range seq {
		strs[i] = strconv.Itoa(v)
	}

	res, _ := strconv.Atoi(strings.Join(strs, ""))

	return res
}