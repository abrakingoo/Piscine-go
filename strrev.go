package piscine

func StrRev(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i, j := 0, length-1; i < length/2; i++ {
		runes[i], runes[j-i] = runes[j-i], runes[i]
	}
	return string(runes)
}
