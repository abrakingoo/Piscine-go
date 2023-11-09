package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	shoppingSummary := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else {
			shoppingSummary[word]++
			word = ""
		}
	}
	shoppingSummary[word]++ // Count the last word
	return shoppingSummary
}
