package piscine

func IterativeFactorial(nb int) int {
	result := 1
	maxInt := 1<<31 - 1

	if nb < 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			if result > maxInt/i {
				return 0
			}
			result *= i
		}
	}
	return result
}
