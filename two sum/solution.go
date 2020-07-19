package kata

func TwoSum(numbers []int, target int) [2]int {
	result := [2]int{}

	for index, val := range numbers {
		for i := index + 1; i < len(numbers); i++ {
			if numbers[i] == target - val {
				result[0], result[1] = index, i
				return result
			}
		}
	}

	return result
}
