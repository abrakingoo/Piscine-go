package piscine

func NRune(s string, n int) rune {
	rRune := []rune(s)
	for index, char := range rRune {
		if index+1 == n {
			return char
		}
	}
	return 0
}
