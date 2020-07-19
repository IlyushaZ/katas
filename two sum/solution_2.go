package kata

func TwoSum2(numbers []int, target int) [2]int {
	var tmp int
	var result [2]int
	hash := make(map[int]int)

	for index, val := range numbers {
		hash[val] = index
	}

	for i := 0; i < len(numbers); i++ {
		tmp = target - numbers[i]

		if val, ok := hash[tmp]; ok {
			result[0], result[1] = i, val
			break
		}
	}

	return result
}
