package piscine

func NRune(s string, n int) rune {
	rRune := []rune(s)
	if n >= 0 && n < len(rRune) {
		for i := 0; i <= len(rRune); i++ {
			if i == n-1 {
				return rRune[i]
			}
		}
	} else if n < 0 && len(rRune)+n >= 0 {
		return rRune[len(rRune)+n]
	} else if n == len(rRune)-1 {
		return rRune[len(rRune)-1]
	}
	return 0
}
