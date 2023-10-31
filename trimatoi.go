package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	sign := 1
	prevChar := ' '
	for _, char := range s {
		if char == '-' && !(prevChar >= '0' && prevChar <= '9') {
			sign = -1
		}
		if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		}
		prevChar = char
	}
	return res * sign
}
