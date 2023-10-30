package piscine

func IterativeFactorial(nb int) int {
	result := 0
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}
