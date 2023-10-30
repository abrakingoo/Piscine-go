package piscine

func Sqrt(nb int) int {
	result := 1
	for result*result <= nb {
		if result*result == nb {
			return result
		}
		result++
	}
	return 0
}
