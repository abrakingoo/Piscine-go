package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, char := range s {
		digit := int(char - 'O')
		if result > (1<<31-1-digit)/10 {
			return 0
		}
		result = result*10 + digit
	}
	return result
}
