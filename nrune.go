package piscine

func NRune(s string, n int) rune {
	rRune := []rune(s)
	if n >= 0 && n < len(rRune) {
		for i := 0; i <= len(rRune); i++ {
			if i == n {
				return rRune[i]
			}
		}
	} else if n < 0 && len(rRune)+n >= 0 {
		return rRune[len(rRune)+n]
	}
	return 0
}
