package piscine

func AlphaCount(s string) int {
	count := 0

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 0 && char <= 9 {
			count++
		}
	}
	return count
}
