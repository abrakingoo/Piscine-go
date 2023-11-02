package piscine

func SplitWhiteSpaces(s string) []string {
	str := []string{}
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			str = append(str, word)
			word = ""
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		str = append(str, word)
	}
	return str
}
