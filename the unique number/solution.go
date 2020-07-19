package kata

func FindUniq(arr []float32) float32 {
	counts := make(map[float32]int)
	var unique float32

	for _, value := range arr {
		counts[value] += 1
	}

	for key, value := range counts {
		if value == 1 { unique = key }
	}

	return unique
}
