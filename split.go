package piscine

func Split(s, sep string) []string {
	new := []string{}
	sRune := []rune(s)
	sepLen := len(sep)
	start := 0

	for i := 0; i < len(sRune); i++ {
		if i+sepLen <= len(sRune) && string(sRune[i:i+sepLen]) == sep {
			new = append(new, string(sRune[start:i]))
			start = i + sepLen
			i += sepLen - 1
		}
	}
	if start < len(sRune) {
		new = append(new, string(sRune[start:]))
	}
	return new
}
