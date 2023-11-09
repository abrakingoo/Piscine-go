package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	shoppingSummary := make(map[string]int)
	word := ""
	for _, char := range str {
		if char != ' ' {
			word += string(char)
		} else {
			if word != "" {
				shoppingSummary[word]++
				word = ""
			}
		}
	}
	if word != "" {
		shoppingSummary[word]++
	}
	return shoppingSummary
}
