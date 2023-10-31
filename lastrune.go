package piscine

func LastRune(s string) rune {
	rRune := []rune(s)
	return rRune[len(rRune)-1]
}
