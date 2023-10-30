package piscine

import "math"

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		if result > math.MaxInt32/i {
			return 0
		}
		result *= i
	}
	return result
}
