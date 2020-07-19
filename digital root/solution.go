package kata

func DigitalRoot(n int) int {
	sum := sumDigits(n)
	if countDigits(sum) > 1 {
		sum = DigitalRoot(sum)
	}

	return sum
}

func countDigits(n int) int {
	var count int

	for n != 0 {
		n /= 10
		count++
	}

	return count
}

func sumDigits(n int) int {
	var sum int

	for n != 0 {
		rem := n % 10
		sum += rem
		n = n / 10
	}

	return sum
}