package piscine

func NRune(s string, n int) rune {
	rRune := []rune(s)
	if n <= 0 && n < len(rRune) {
		return rRune[n]
	}
	return 0
}
