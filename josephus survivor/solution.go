package kata

func JosephusSurvivor(n, k int) int {
	if n == 1 || k == 1 {
		return n
	}

	seq := make([]int, n)
	for i := 0; i < len(seq); i++ {
		seq[i] = i + 1
	}

	for i := k - 1; len(seq) > 1; i += k - 1 {
		for i >= len(seq) {
			i = i - len(seq)
		}

		seq = append(seq[:i], seq[i+1:]...)
	}

	return seq[0]
}
