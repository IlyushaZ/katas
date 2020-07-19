package kata

import (
	"fmt"
)

func SumFracts(arr [][]int) string {
	commDen := 1
	for _, n := range arr {
		commDen *= n[1]
	}

	var num int
	for _, n := range arr {
		n[0], n[1] = n[0] * (commDen / n[1]), commDen
		num += n[0]
	}

	if num % commDen == 0 {
		return fmt.Sprintf("%d", num / commDen)
	}

	gcd := gcd(num, commDen)
	return fmt.Sprintf("%d/%d", num / gcd, commDen / gcd)
}

func gcd(a, b int) int {
	var rem int
	for {
		if b == 0 {
			break
		}

		rem = a % b
		a, b = b, rem
	}

	return a
}
