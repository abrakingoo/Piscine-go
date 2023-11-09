package piscine

func JumpOver(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}
	newS := ""

	for i := 2; i < len(str); i += 3 {
		newS += string(str[i])
	}
	return newS + "\n"
}
