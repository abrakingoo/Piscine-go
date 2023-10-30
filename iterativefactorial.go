package piscine

func IterativeFactorial(nb int) int {
	result := 1
	maxInt := 1<<31 - 1

	if nb < 0 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		if result > maxInt/i {
			return 1
		}
		result *= i
	}
	return result
}
