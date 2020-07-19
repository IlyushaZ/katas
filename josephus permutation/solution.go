package kata

func Josephus(items []interface{}, k int) []interface{} {
	if k == 1 || len(items) == 0 {
		return items
	}
	var result []interface{}

	for i := k - 1; len(items) > 0; i += k - 1 {
		for i >= len(items) {
			i = i - len(items)
		}

		result = append(result, items[i])
		items = append(items[:i], items[i+1:]...)
	}

	return result
}
